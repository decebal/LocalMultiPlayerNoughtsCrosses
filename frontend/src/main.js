import Vue from 'vue';
import VueNativeSock from 'vue-native-websocket';

import App from './App.vue';
import store from './store';
import { HOST, PORT } from './config';

Vue.use(VueNativeSock, `ws://${HOST}:${PORT}/ws`, {
  store,
  format: 'json',
});

Vue.config.productionTip = false;

new Vue({
  store,
  render: h => h(App),
}).$mount('#app');
