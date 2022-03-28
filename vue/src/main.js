import Vue from 'vue';
import axios from 'axios';
import App from './App.vue';
import vuetify from './plugins/vuetify';
import '@carbon/charts/styles.css';
import chartsVue from '@carbon/charts-vue';
import '@/assets/styling/main.scss';
import VueApexCharts from 'vue-apexcharts'

Vue.use(VueApexCharts);
Vue.use(chartsVue);

Vue.component('apexchart', VueApexCharts)

Vue.config.productionTip = false;
Vue.prototype.$axios = axios;

new Vue({
  vuetify,
  render: h => h(App)
}).$mount('#app');