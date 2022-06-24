import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import MintUI from 'mint-ui'
import 'mint-ui/lib/style.css'

import { utils } from './utils'

import axios from 'axios';
Vue.prototype.$http = axios;

Vue.use(MintUI)

Vue.config.productionTip = false

const startApp = () => {
  new Vue({
    router,
    store,
    render: h => h(App)
  }).$mount('#app')
}

export function init() {
  let query = utils.parseQuery();
  // 判断是否获取到临时授权码
  const { code } = query;
  


}


init();


