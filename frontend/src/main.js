import Vue from 'vue'
import App from './App.vue'

import './main.scss'
import './markdown.css'
import vuetify from './plugins/vuetify';
import router from './router'

Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
