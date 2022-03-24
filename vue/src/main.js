import Vue from 'vue';
import App from './App.vue';
import vuetify from './plugins/vuetify';
import '@carbon/charts/styles.css';
import chartsVue from '@carbon/charts-vue';
import '@/assets/styling/main.scss';

Vue.use(chartsVue);

Vue.config.productionTip = false;

new Vue({
  vuetify,
  render: h => h(App)
}).$mount('#app');
