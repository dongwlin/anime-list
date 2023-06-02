import { createApp } from 'vue';
import './style.css';
import App from './App.vue';
import router from "@/router/index.js";
import "element-plus/theme-chalk/index.css";
import * as ElementPlusIconsVue from '@element-plus/icons-vue';
import { createPinia } from 'pinia';

const app = createApp(App);
app.use(router);
app.use(createPinia());
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}


app.mount('#app');
