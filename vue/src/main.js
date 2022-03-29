import Vue from 'vue';
import axios from 'axios';
import App from './App.vue';
import vuetify from './plugins/vuetify';
import VueApexCharts from 'vue-apexcharts'
import '@mdi/font/css/materialdesignicons.css';

import '@/assets/styling/main.scss';

Vue.use(VueApexCharts);

Vue.component('apexchart', VueApexCharts)

Vue.config.productionTip = process.env.NODE_ENV !== 'development';
Vue.prototype.$axios = axios;

new Vue({
  vuetify,
  render: h => h(App)
}).$mount('#app');
