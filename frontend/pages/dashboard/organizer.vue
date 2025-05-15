<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']">
    <Navbar />

    <!-- Фон с падающими звездами как на главной странице -->
    <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
      <div v-for="i in 30" :key="i" class="falling-star" :style="randomFallingStarStyle()" />
    </div>

    <div class="container mx-auto p-6 space-y-8 relative z-10 pt-20">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-accent">Панель организатора</h1>
        <p class="text-opacity-80">Создавайте и управляйте вашими мероприятиями</p>
      </div>

      <!-- Profile Card -->
      <div class="max-w-md mx-auto bg-white dark:bg-gray-800 rounded-xl shadow-md overflow-hidden md:max-w-2xl mb-6 events-section">
        <div class="md:flex p-6">
          <div class="md:flex-shrink-0 flex items-center justify-center">
            <img
              v-if="user && user.avatar"
              :src="user.avatar"
              alt="User avatar"
              class="h-24 w-24 rounded-full object-cover"
            />
            <div v-else class="h-24 w-24 rounded-full bg-gray-300 dark:bg-gray-600 flex items-center justify-center text-gray-500 dark:text-gray-400 text-3xl font-semibold">
              {{ user ? user.name.charAt(0).toUpperCase() : '' }}
            </div>
          </div>
          <div class="mt-4 md:mt-0 md:ml-6 flex flex-col justify-center">
            <h2 class="text-xl font-semibold">{{ user ? user.name : '' }}</h2>
            <p class="text-opacity-80">{{ user ? user.email : '' }}</p>
            <p class="text-opacity-80 capitalize">{{ user ? user.role : '' }}</p>
            <div class="flex gap-3 mt-4">
              <button
                @click="goToStatistics"
                class="px-4 py-2 bg-purple-600 text-white rounded-lg transition-colors duration-300 hover:bg-purple-700 flex items-center gap-2"
              >
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
                </svg>
                Статистика
              </button>
              <button
                @click="logout"
                class="px-4 py-2 action-btn rounded-lg transition-colors duration-300"
              >
                Выйти
              </button>
            </div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Events Section -->
        <div class="lg:col-span-2">
          <div class="p-6 rounded-2xl events-section shadow-lg">
            <div class="flex justify-between items-center mb-6">
              <h2 class="text-xl font-bold text-accent">Мои мероприятия</h2>
              <button class="text-sm font-medium text-blue-500 hover:text-blue-400 transition-colors">
                Смотреть все
              </button>
            </div>

            <!-- События -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div
                v-for="event in filteredEvents"
                :key="event.id"
                class="p-4 rounded-xl event-card shadow-md cursor-pointer hover:shadow-lg transition-all duration-300 group"
                @click="goToEvent(event.id)"
              >
                <div class="flex justify-between items-start mb-2">
                  <h3 class="text-lg font-bold">{{ event.title }}</h3>
                  <span
                    class="px-2 py-1 rounded-full text-xs"
                    :class="{
                      'bg-blue-100 text-blue-800': event.status === 'upcoming',
                      'bg-green-100 text-green-800': event.status === 'current',
                      'bg-gray-100 text-gray-800': event.status === 'past'
                    }"
                  >
                    {{ event.status === 'upcoming' ? 'Предстоящее' :
                       event.status === 'current'  ? 'Текущее'     :
                       'Завершённое' }}
                  </span>
                </div>
                <p class="text-sm text-opacity-80 mb-2">{{ formatDate(event.date) }}</p>
                <p class="text-sm text-opacity-60">
                  {{ event.description ? event.description.substring(0, 100) + '...' : '' }}
                </p>
                <div class="mt-3 flex flex-wrap gap-2">
                  <span
                    v-for="(tag, i) in event.tags.split(',')"
                    :key="i"
                    class="px-2 py-1 rounded-full text-xs"
                    :class="tagColors[i % tagColors.length]"
                  >
                    {{ tag.trim() }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Right Column: Add New Event -->
        <div class="col-span-1 p-6 rounded-2xl events-section shadow-lg">
          <h2 class="text-xl font-bold text-accent mb-4">Добавить мероприятие</h2>
          <form @submit.prevent="addEvent">
            <!-- Название -->
            <div class="mb-4">
              <label class="block mb-1 text-opacity-80">Название</label>
              <input
                v-model="newEvent.title"
                type="text"
                class="w-full p-3 rounded-xl input-field"
                placeholder="Название"
              />
            </div>

            <!-- Дата -->
            <div class="mb-4">
              <label class="block mb-1 text-opacity-80">Дата (ДД.ММ.ГГГГ)</label>
              <input
                v-model="newEvent.date"
                type="text"
                class="w-full p-3 rounded-xl input-field"
                placeholder="ДД.ММ.ГГГГ"
              />
            </div>

            <!-- Описание -->
            <div class="mb-4">
              <label class="block mb-1 text-opacity-80">Описание</label>
              <textarea
                v-model="newEvent.description"
                class="w-full p-3 rounded-xl input-field"
                placeholder="Описание"
                rows="3"
              ></textarea>
            </div>

            <!-- Теги -->
            <div class="mb-4">
              <label class="block mb-1 text-opacity-80">Теги</label>
              <input
                v-model="newEvent.tags"
                type="text"
                class="w-full p-3 rounded-xl input-field"
                placeholder="Теги через запятую"
              />
            </div>

            <!-- Файл изображения -->
            <div class="mb-4">
              <label class="block mb-1 text-opacity-80">Изображение</label>
              <input 
                type="file" 
                @change="onFileChange" 
                class="w-full p-2 rounded-lg"
              />
              <p v-if="newEvent.image" class="text-sm text-opacity-60 mt-1">{{ newEvent.image.name }}</p>
            </div>

            <div class="flex gap-3 mt-6">
              <button type="submit" class="flex-1 p-3 rounded-xl action-btn">
                <div class="flex items-center justify-center gap-2">
                  <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                  </svg>
                  Добавить
                </div>
              </button>
              <button type="button" class="flex-1 p-3 rounded-xl cancel-btn" @click="resetForm">
                Отменить
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Theme Toggle Button -->
    <button @click="toggleTheme" class="theme-toggle fixed bottom-6 right-6 p-3 rounded-full shadow-lg hover:scale-110 transition-transform">
      <svg v-if="isDarkMode" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
      </svg>
      <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import axios from 'axios'
import Navbar from '@/components/Navbar.vue'

// State
const user = ref(null);
const error = ref(null);
const events = ref([]);
const theme = useState('theme');
const isDarkMode = computed(() => theme.value === 'dark');
const activeTab = ref('participating');

// Форма для нового мероприятия
const newEvent = ref({
  title: '',
  date: '',
  description: '',
  tags: '',
  image: null
});

const tabs = [
  { id: 'participating', label: 'Текущие' },
  { id: 'upcoming', label: 'Предстоящие' },
  { id: 'participated', label: 'Прошедшие' }
];

const tagColors = [
  'bg-blue-500/20 text-blue-500',
  'bg-purple-500/20 text-purple-500',
  'bg-green-500/20 text-green-500',
  'bg-yellow-500/20 text-yellow-500'
];

onMounted(async () => {
  axios.defaults.baseURL = 'http://localhost:8080'
  axios.defaults.withCredentials = true

  // Загружаем и профиль, и список мероприятий
  await Promise.all([fetchUserData(), fetchEvents()])
});

async function fetchUserData() {
  try {
    const { data } = await axios.get('/profile')
    user.value = data
  } catch (e) {
    console.error('Error fetching user data:', e)
    error.value = 'Не удалось загрузить данные о пользователе.'
  }
}

async function fetchEvents() {
  try {
    const { data } = await axios.get('/events')
    events.value = data
  } catch (e) {
    console.warn('Не удалось загрузить мероприятия, используем локальные данные.')
  }
}

function logout() {
  axios.post('/logout')
  navigateTo('/login')
}

function toggleTheme() {
  theme.value = isDarkMode.value ? 'light' : 'dark'
}

function goToEvent(id) {
  navigateTo(`/event/${id}`)
}

function onFileChange(e) {
  newEvent.value.image = e.target.files[0]
}

function addEvent() {
  console.log('Добавляем событие', newEvent.value)
  resetForm()
}

function resetForm() {
  newEvent.value = { title: '', date: '', description: '', tags: '', image: null }
}

function formatDate(d) {
  return d
}

function randomFallingStarStyle() {
  const size = Math.floor(Math.random() * 4) + 1
  const left = Math.random() * 100
  const top = Math.random() * -100
  const duration = 5 + Math.random() * 5
  return {
    width: `${size}px`,
    height: `${size}px`,
    left: `${left}%`,
    top: `${top}%`,
    animationDuration: `${duration}s`,
    animationDelay: `${Math.random() * 3}s`
  }
}

function goToStatistics() {
  navigateTo('/statistics')
}

// Computed
const filteredEvents = computed(() => {
  return (events.value || []).filter(evt => evt.status === activeTab.value)
})
</script>

<style scoped>
/* Light Theme */
.light-theme {
  background-color: #F5F5F7;
  color: #1A1A2E;
}

.light-theme .events-section {
  background-color: #FFFFFF;
}

.light-theme .event-card {
  background-color: #FFFFFF;
  border: 1px solid #E4E4F0;
}

.light-theme .input-field {
  background-color: #F0F0F5;
  border: 1px solid #E4E4F0;
  color: #1A1A2E;
}

.light-theme .theme-toggle,
.light-theme .action-btn {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.light-theme .cancel-btn {
  background-color: #F0F0F5;
  color: #1A1A2E;
}

/* Dark Theme */
.dark-theme {
  background-color: #1A1A2E;
  color: #E4E4F0;
}

.dark-theme .events-section {
  background-color: #2A2A40;
}

.dark-theme .event-card {
  background-color: #3A3A50;
  border: 1px solid #4A4A60;
}

.dark-theme .input-field {
  background-color: #3A3A50;
  border: 1px solid #E4E4F0;
  color: #E4E4F0;
}

.dark-theme .theme-toggle,
.dark-theme .action-btn {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .cancel-btn {
  background-color: #3A3A50;
  color: #E4E4F0;
}

/* Падающие звезды */
.falling-star {
  position: absolute;
  background-color: #6C63FF;
  border-radius: 50%;
  animation: fall linear infinite;
  filter: blur(1px);
}

@keyframes fall {
  0% {
    top: -10%;
    opacity: 1;
  }
  100% {
    top: 110%;
    opacity: 0;
  }
}

/* Общие стили */
.event-card {
  transition: all 0.3s ease;
}

.event-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.action-btn, .cancel-btn {
  transition: all 0.3s ease;
}

.action-btn:hover, .cancel-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.action-btn {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.text-opacity-80 {
  opacity: 0.8;
}

.text-opacity-60 {
  opacity: 0.6;
}

.text-accent {
  color: #6C63FF;
}

.animate-fade-slide {
  animation: fadeSlide 0.5s ease-out;
}

@keyframes fadeSlide {
  0% { opacity: 0; transform: translateY(-20px); }
  100% { opacity: 1; transform: translateY(0); }
}
</style>
