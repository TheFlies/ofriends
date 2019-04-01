import Vue from 'vue'
import FlagIcon from 'vue-flag-icon'

import countries from './countries.json'

Vue.use(FlagIcon)
Vue.prototype.$countries = countries
