<template>
  <div :class="['min-h-screen transition-colors duration-500', isDarkMode ? 'dark-theme' : 'light-theme']" @mousemove="handleParallax">
    <Navbar />

    <!-- Звёздный фон -->
    <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">

      <!-- Мерцающие звёзды -->
      <div v-for="i in 100" :key="'star-' + i" class="twinkle-star" :style="randomTwinkleStarStyle()" />
      <!-- Падающие звёзды -->
      <div v-for="i in 30" :key="'falling-' + i" class="falling-star" :style="randomFallingStarStyle()" />
    </div>

    <div class="container mx-auto p-6 relative z-10 pt-20">
      <div class="event-card p-6 rounded-2xl shadow-lg w-full max-w-4xl mx-auto fade-in relative overflow-hidden">
        <div class="event-header p-6 rounded-t-xl">
          <h1 class="text-3xl font-orbitron font-bold">IT Конференция 2025</h1>
          <div class="mt-4 flex gap-2 flex-wrap">
            <span class="px-3 py-1 bg-accent bg-opacity-20 rounded-full text-sm">#конференция</span>
            <span class="px-3 py-1 bg-accent bg-opacity-20 rounded-full text-sm">#офлайн</span>
          </div>
        </div>

        <div class="p-6 space-y-4">
          <div class="flex items-center">
            <span class="mr-2">📅</span>
            <p>Дата: 10 апреля 2025 | Продолжительность: 2 дня</p>
          </div>
          <div class="flex items-center">
            <span class="mr-2">🏢</span>
            <p>Организатор: ООО "ТехноСфера"</p>
          </div>
          <div class="flex items-center">
            <span class="mr-2">📍</span>
            <p>Формат: Офлайн | Место: Москва, ул. Ленина, 10</p>
          </div>
        </div>

        <div class="map-container mt-6 h-64 rounded-xl flex items-center justify-center border-2 border-dashed">
          <p>Здесь будет карта</p>
        </div>

        <div class="p-6 flex justify-end gap-4">
          <button class="px-6 py-2 bg-accent text-white rounded-lg hover:bg-opacity-90 transition">
            Зарегистрироваться
          </button>
          <button class="exit-button px-6 py-2 text-white rounded-lg hover:bg-opacity-90 transition">
            Выход
          </button>
        </div>
      </div>
    </div>

    <!-- Переключатель темы -->
    <button @click="toggleTheme" class="fixed bottom-4 right-4 p-3 rounded-full shadow-lg z-50 transition-all bg-accent text-white hover:scale-105">
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
import { ref, onMounted } from 'vue'
import Navbar from '~/components/Navbar.vue'

const isDarkMode = ref(true)
const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value
}

const randomFallingStarStyle = () => {
    const size = Math.floor(Math.random() * 4) + 1
    const left = Math.random() * 100
    const duration = 5 + Math.random() * 5
    return {
        width: `${size}px`,
        height: `${size}px`,
        left: `${left}%`,
        top: `-${size * 2}px`, // чтобы не улетали вверх
        animationDuration: `${duration}s`,
        animationDelay: `${Math.random() * 3}s`,
        position: 'absolute',
        background: 'white',
        borderRadius: '50%',
        animationName: 'falling',
        animationTimingFunction: 'linear',
        animationIterationCount: 'infinite'
    }
}


const randomTwinkleStarStyle = () => {
  const size = Math.random() * 2 + 0.5
  return {
    width: `${size}px`,
    height: `${size}px`,
    top: `${Math.random() * 100}%`,
    left: `${Math.random() * 100}%`,
    position: 'absolute',
    background: 'white',
    borderRadius: '50%',
    animation: `twinkle ${2 + Math.random() * 3}s ease-in-out infinite`,
    opacity: Math.random()
  }
}

// Параллакс
const parallaxContainer = ref(null)
const handleParallax = (e) => {
  const { innerWidth, innerHeight } = window
  const x = (e.clientX - innerWidth / 2) / innerWidth * 10
  const y = (e.clientY - innerHeight / 2) / innerHeight * 10
  if (parallaxContainer.value) {
    parallaxContainer.value.style.transform = `translate(${x}px, ${y}px)`
  }
}
</script>

<style scoped>
/* Падающие звёзды */
.falling-star {
  position: absolute;
  background-color: white;
  border-radius: 9999px;
  opacity: 0.8;
  animation: falling linear infinite;
}

@keyframes falling {
  0% {
    transform: translateY(0);
    opacity: 1;
  }
  100% {
    transform: translateY(calc(100vh + 20px));
    opacity: 0;
  }
}


/* Мерцающие звёзды */
.twinkle-star {
  position: absolute;
  background-color: white;
  border-radius: 50%;
}

@keyframes twinkle {
  0%, 100% { opacity: 0.2; }
  50% { opacity: 1; }
}

/* Плавное появление */
.fade-in {
  animation: fadeIn 1.2s ease forwards;
  opacity: 0;
}

@keyframes fadeIn {
  to { opacity: 1; }
}

/* Темы */
.dark-theme {
  background-color: #1A1A2E;
  color: #E4E4F0;
}

.dark-theme .event-card {
  background-color: #2A2A40;
}

.dark-theme .event-header {
  background-color: #3A3A50;
}

.dark-theme .map-container {
  background-color: #3A3A50;
  border-color: #6C63FF;
}

.light-theme {
  background-color: #D5D5E0;
  color: #1A1A2E;
}

.light-theme .event-card {
  background-color: #E4E4F0;
}

.light-theme .event-header {
  background-color: #F0F0F5;
}

.light-theme .map-container {
  background-color: #F0F0F5;
  border-color: #6C63FF;
}

/* Общие стили */
.text-accent {
  color: #6C63FF;
}

.bg-accent {
  background-color: #6C63FF;
}

.exit-button {
  background-color: #FF6584;
}

.exit-button:hover {
  background-color: #e55a78;
}

/* Неоновая полоска сбоку карточки */
.event-card::before {
  content: '';
  position: absolute;
  top: 0;
  bottom: 0;
  left: -6px;
  width: 4px;
  background: linear-gradient(#6C63FF, #FF6584);
  border-radius: 9999px;
  filter: blur(6px);
}

</style>
