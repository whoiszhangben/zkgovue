<template>
  <div class="menu-wrapper">
    <mt-cell v-for="item in list" :key="item.id">
      {{ item.token }}：{{ item.message }}
    </mt-cell>
  </div>
</template>

<script>
import { Cell } from "mint-ui";
import EventBus from '@/events/event-bus.js'

export default {
  data() {
    return {
      list: [
      ]
    }
  },
  components: {
    Cell
  },
  created() {
    this.registerEvent()
  },

  destroyed() {
    this.destroyEvent()
  },
  methods: {
    registerEvent() {
      EventBus.$on('add-message', (data) => {
        data.id = Math.floor(Math.random() * 10000)
        this.list.push(data)
      })
    },

    destroyEvent() {
      EventBus.$off('add-message')
    }
  }
}
</script>
<style>
  .user {
    display: inline-block;
    max-width: 10em;
    /*overflow: hidden;*/
    /*white-space: nowrap;*/
    /*word-wrap: normal;*/
    text-overflow: ellipsis;
    color: #3c9cfe;
    cursor: pointer;
  }
</style>
