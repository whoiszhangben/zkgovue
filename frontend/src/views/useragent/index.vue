<template>
  <div class="home">
    <mt-field
      :label="item.name"
      v-for="item in navigatorList"
      :key="item.code"
      type="textarea"
      :rows="item.rows"
      :value="item.value"
    ></mt-field>
  </div>
</template>

<script>
// @ is an alias to /src
import { MessageBox } from "mint-ui";
export default {
  name: "userAgent",
  data() {
    return {
      navigatorList: [
        {
          code: "appVersion",
          name: "版本",
          value: navigator.appVersion,
          rows: 6,
        },
        {
          code: "appName",
          name: "名称",
          value: navigator.appName,
          row: 1,
        },
        {
          code: "language",
          name: "语言",
          value: navigator.language,
          rows: 1,
        },
        {
          code: "platform",
          name: "平台",
          value: navigator.platform,
          rows: 1,
        },
        {
          code: "userAgent",
          name: "信息",
          value: navigator.userAgent,
          rows: 6,
        },
      ],
    };
  },
  components: {
    MessageBox,
  },
  mounted() {
    MessageBox.confirm("即将将您的浏览器内核信息上报到服务器, 是否继续")
      .then((res) => {
        let postData = {
          appVersion: navigator.appVersion,
          appName: navigator.appName,
          language: navigator.language,
          platform: navigator.platform,
          userAgent: navigator.userAgent,
        };
        console.log(postData);
        this.$http
          .post("/api/info", postData)
          .then((res) => {
            console.log(res);
          })
          .catch((err) => {
            console.log(err);
          });
      })
      .catch((err) => {
        console.log("您取消了操作");
      });
  },
};
</script>
