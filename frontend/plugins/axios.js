import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080'; // Замените на ваш сервер
axios.defaults.withCredentials = true; // Для отправки токенов через cookies
