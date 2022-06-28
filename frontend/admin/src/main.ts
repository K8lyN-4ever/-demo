import { VuesticPlugin, useGlobalConfig } from 'vuestic-ui';
import element from 'element-plus'
import 'element-plus/theme-chalk/index.css';
import { createApp } from 'vue'
import { createI18n } from 'vue-i18n'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'
import vuesticGlobalConfig from './services/vuestic-ui/global-config'

const i18nConfig = {
  locale: 'en',
  fallbackLocale: 'en',
  messages: {
    en: require('@/i18n/en.json'),
    ch: require('@/i18n/cn.json'),
    es: require('@/i18n/es.json'),
    ir: require('@/i18n/ir.json'),
    br: require('@/i18n/br.json')
  }
}

// initAMapApiLoader({
//   key: "8b81ae24a6396cb04290148fcc45f317"
// })

const app = createApp(App)
app.use(store)
app.use(router)
app.use(element)
// if (process.env.VUE_APP_GTM_ENABLED === 'true') {
  // const gtmConfig = {
  //   id: process.env.VUE_APP_GTM_KEY,
  //   debug: false,
  //   vueRouter: router,
  // }
  // app.use(createGtm(gtmConfig))
// }
app.use(createI18n(i18nConfig))
app.use(VuesticPlugin, vuesticGlobalConfig)
// app.use(VueAMap)
app.mount('#app')

