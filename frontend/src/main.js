import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

import Default from "./layouts/Default";
import Post from "./layouts/Post";

import { library } from "@fortawesome/fontawesome-svg-core";
import { faTelegram } from "@fortawesome/free-brands-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { faBars } from "@fortawesome/free-solid-svg-icons";
import { faChevronUp } from "@fortawesome/free-solid-svg-icons";
import { Plugin } from "vue-fragment";

library.add(faTelegram);
library.add(faBars);
library.add(faChevronUp);
Vue.component("font-awesome-icon", FontAwesomeIcon);

Vue.component("default-layout", Default);
Vue.component("post-layout", Post);
Vue.use(require("vue-moment"));
Vue.use(Plugin);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
