<template>
  <div>
    <mt-header title="我的聊天室">
      <router-link to="/" slot="left">
        <mt-button icon="back">返回</mt-button>
      </router-link>
    </mt-header>
    <mt-field label="测试发言" placeholder="请文明发言" type="textarea" rows="4" v-model="message"></mt-field>
    <mt-button type="primary" size="large" style="text-align:center;" @click="onSubmit">发送</mt-button>
    <message-log></message-log>
  </div>
</template>

<script>
import { Field, Button, Header,MessageBox } from "mint-ui";
import config from "@/config/index";
import EventBus from "@/events/event-bus.js";
import MessageLog from "./messageLog";
import store from "@/store";
export default {
  name: "websocket",
  data() {
    return {
      token: "",
      nickName: "",
      wenzi_message_websocket: undefined,
      message: "",
    };
  },
  components: {
    Field,
    Button,
    Header,
    MessageLog,
  },

  created() {
    console.log("hahaha");
    this.connect();
    this.registerEvent();
    this.heartBeat();
  },

  destroyed() {
    this.destroyEvent();
  },
  mounted() {
    if (!store.getters.nickName) {
      MessageBox.prompt("欢迎来畅聊，请输入昵称").then(({ value, action }) => {
        this.nickName = value;
      })
    }
  },
  methods: {
    heartBeat() {
      //  心跳
      setInterval(() => {
        var data = {
          action: "heart_beat",
        };
        this.sendWebSocketData(data);
      }, 30000);
    },
    sendWebSocketData(data) {
      if (
        !this.wenzi_message_websocket ||
        this.wenzi_message_websocket.readyState === 3
      ) {
        return;
      }
      // websocket 请求加上token 作为唯一标识
      data.token = this.token;
      data = JSON.stringify(data);
      this.wenzi_message_websocket.send(data);
    },
    onSubmit() {
      EventBus.$emit("send-message", {
        action: "sendMessage",
        token: this.$store.getters.token,
        nickName: this.nickName,
        message: this.message,
      });
      this.message = "";
    },
    connect() {
      let self = this;
      // 参考链接：swoole 快速起步 swoole 部分： https://wiki.swoole.com/wiki/page/479.html
      if (window.wenzi_message_websocket) {
        this.wenzi_message_websocket = window.wenzi_message_websocket;
        return;
      }

      console.log("开始连接");
      var wenzi_message_websocket = new WebSocket(
        config.wenzi_message_websocket_uri
      );

      wenzi_message_websocket.onopen = (evt) => {
        var data = {
          action: "open",
        };

        this.sendWebSocketData(data);
      };

      wenzi_message_websocket.onmessage = (evt) => {
        console.log(evt.data);
        var res = JSON.parse(evt.data);
        if (res.action === "open") {
          // 没有token ，设置token
          if (this.token === "") {
            this.token = res.data.token;
            this.$store.dispatch("modifyToken", res.data.token);
          }
        }
        if (res.action === "close") {
          alert(res.data.message);
          this.closeWebPage();
        }

        if (res.action === "replyMessage") {
          EventBus.$emit("add-message", {
            token: res.data.token,
            nickName: res.data.nickName,
            message: res.data.message,
          });
        }
      };

      wenzi_message_websocket.onclose = (evt) => {
        console.log(evt);
        console.log("断开连接");
        window.wenzi_message_websocket = null;
        this.wenzi_message_websocket = null;
        console.log("reconnection");
        this.connect();
      };

      wenzi_message_websocket.onerror = function (evt, e) {
        console.log("Error occured: " + evt.data);
      };

      window.wenzi_message_websocket = wenzi_message_websocket;
      this.wenzi_message_websocket = window.wenzi_message_websocket;
    },

    registerEvent() {
      EventBus.$on("send-message", (data) => {
        this.sendWebSocketData(data);
      });
    },

    destroyEvent() {
      EventBus.$off("send-message");
    },
  },
};
</script>
