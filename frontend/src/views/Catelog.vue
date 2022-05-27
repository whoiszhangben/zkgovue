<template>
  <div>
    <mt-header title="demo列表"></mt-header>
    <mt-cell v-for="item in indexList" :key="item.code" :value="item.value" is-link :to="item.url"></mt-cell>
  </div>
</template>

<script>
// @ is an alias to /src
import { Cell, Header } from "mint-ui";
export default {
  name: "Catelog",
  data() {
    return {
      indexList: [
        {
          code: "a",
          value: "userAgent收集",
          url: "/userAgent",
        },
        {
          code: "b",
          value: "websocket测试",
          url: "/websocket",
        },
      ],
    };
  },
  components: {
    Cell,
    Header,
  },
};
</script>
