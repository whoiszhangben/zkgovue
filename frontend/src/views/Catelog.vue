<template>
  <div>
    <mt-header title="功能测试"></mt-header>
    <mt-cell v-for="item in indexList" :key="item.code" :value="item.value" is-link :to="item.url"></mt-cell>
  </div>
</template>

<script>
// @ is an alias to /src
import { Cell, Header } from 'mint-ui'
export default {
  name: 'Catelog',
  data() {
    return {
      indexList: [
        {
          code: 'a',
          value: 'userAgent收集',
          url: '/userAgent',
        },
        {
          code: 'b',
          value: 'websocket聊天室',
          url: '/websocket',
        },
      ],
    }
  },
  components: {
    Cell,
    Header,
  },
}
</script>
