import { createApp } from 'vue';
import { createPinia } from 'pinia';
import App from './App.vue';
import router from './router'; // если у вас используется роутер

const app = createApp(App);
app.use(createPinia());  // Здесь подключаем Pinia
app.use(router);  // если есть роутер
app.mount('#app');
