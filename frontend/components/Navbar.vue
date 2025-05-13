<template>
  <nav class="nav-bar p-4 flex justify-between items-center shadow-md" :class="isDarkMode ? 'dark-theme' : 'light-theme'">
    <div class="flex items-center">
      <!-- Burger menu button - visible only on mobile -->
      <button @click="toggleMenu" class="md:hidden mr-4">
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
      </button>
      <NuxtLink to="/" class="logo text-2xl font-bold"><img src="public\img\logo.png" alt="Logo" class="logo-img"></NuxtLink>
    </div>
    <div class="hidden md:flex gap-3 items-center">
      <button class="nav-btn">
        <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
        </svg>
        <NuxtLink to="/events">Мероприятия</NuxtLink>
      </button>
        <button class="nav-btn">
          <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192l-3.536 3.536M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-5 0a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <NuxtLink to="/faq">Поддержка</NuxtLink> 
        </button>

      <div class="relative">
        <button @click="toggleNotifications" class="nav-btn relative">
        <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
        </svg>
        <span class="notification-dot"></span>
          Уведомления
        </button>
        <NotificationPanel 
          :show="showNotifications" 
          @close="showNotifications = false"
        />
      </div>
      <div class="relative">
        <button @click="toggleFriends" class="nav-btn">
        <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
        </svg>
          Друзья
        </button>
        <FriendsPanel 
          :show="showFriends" 
          @close="showFriends = false"
        />
      </div>
      <button class="nav-btn logout-btn">
        <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
        </svg>
        Выход
      </button>
      <button class="nav-btn login-btn">
      <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16l4-4m0 0l-4-4m4 4H3m13 4v1a3 3 0 003 3h4a3 3 0 003-3V7a3 3 0 00-3-3h-4a3 3 0 00-3 3v1" />
      </svg>
      <NuxtLink to="/login">Войти</NuxtLink> 
    </button>
    </div>

    <!-- Mobile menu - shown when burger is clicked -->
    <div v-if="isMenuOpen" class="md:hidden fixed inset-0 bg-black bg-opacity-50 z-40" @click="toggleMenu"></div>
    <div v-if="isMenuOpen" class="md:hidden fixed top-16 left-0 right-0 bg-white dark:bg-gray-800 shadow-lg z-50 p-4" :class="isDarkMode ? 'dark-theme' : 'light-theme'">
      <div class="flex flex-col gap-3">
        <!-- Mobile versions of your navigation items -->
        <button @click="toggleTheme" class="mobile-nav-btn">
          <svg v-if="isDarkMode" class="h-5 w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <svg v-else class="h-5 w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
          </svg>
          {{ isDarkMode ? 'Светлая тема' : 'Тёмная тема' }}
        </button>
        <button class="mobile-nav-btn">
          <svg class="h-5 w-5 mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
          <NuxtLink to="/events">Мероприятия</NuxtLink>
        </button>
        <button class="nav-btn">
          <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-3.536 3.536m0 5.656l3.536 3.536M9.172 9.172L5.636 5.636m3.536 9.192l-3.536 3.536M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-5 0a4 4 0 11-8 0 4 4 0 018 0z" />
          </svg>
          <NuxtLink to="/faq">Поддержка</NuxtLink> 
        </button>

      <div class="relative">
        <button @click="toggleNotifications" class="nav-btn relative">
        <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
        </svg>
        <span class="notification-dot"></span>
          Уведомления
        </button>
        <NotificationPanel 
          :show="showNotifications" 
          @close="showNotifications = false"
        />
      </div>
      <div class="relative">
        <button @click="toggleFriends" class="nav-btn">
        <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
        </svg>
          Друзья
        </button>
        <FriendsPanel 
          :show="showFriends" 
          @close="showFriends = false"
        />
      </div>
      <button class="nav-btn logout-btn">
        <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
        </svg>
        Выход
      </button>
      <button class="nav-btn login-btn">
      <svg class="h-5 w-5 mr-1" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16l4-4m0 0l-4-4m4 4H3m13 4v1a3 3 0 003 3h4a3 3 0 003-3V7a3 3 0 00-3-3h-4a3 3 0 00-3 3v1" />
      </svg>
      <NuxtLink to="/login">Войти</NuxtLink> 
    </button>
      </div>
    </div>
  </nav>
</template>

<script setup>
import NotificationPanel from './NotificationPanel.vue'
import FriendsPanel from './FriendsPanel.vue'
import { ref, computed } from 'vue'
const showNotifications = ref(false)
const showFriends = ref(false)

const toggleNotifications = () => {
  showNotifications.value = !showNotifications.value
  if (showNotifications.value) showFriends.value = false
}

const toggleFriends = () => {
  showFriends.value = !showFriends.value
  if (showFriends.value) showNotifications.value = false
}

const theme = useState('theme')
const isDarkMode = computed(() => theme.value === 'dark')
const isMenuOpen = ref(false)
const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
}

const toggleTheme = () => {
  theme.value = isDarkMode.value ? 'light' : 'dark'
}
</script>

<style scoped>
.logo-img {
  max-width: 100px;
  height: auto;
}
.mobile-nav-btn {
  @apply w-full p-3 rounded-lg flex items-center;
}

.light-theme .mobile-nav-btn {
  @apply bg-gray-100 text-gray-800;
}

.dark-theme .mobile-nav-btn {
  @apply bg-gray-700 text-gray-200;
}

/* Prevent scrolling when menu is open */
body.menu-open {
  overflow: hidden;
}
.nav-bar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  transition: all 0.3s ease;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  height: 64px;
}
.logo {
  color: #6C63FF;
  font-family: -apple-system, BlinkMacSystemFont, 'San Francisco', sans-serif;
  letter-spacing: 0.5px;
}

.light-theme.nav-bar {
  background: #E4E4F0;
}

.dark-theme.nav-bar {
  background: #2A2A40;
}

.light-theme .nav-btn {
  background: rgba(240, 240, 245, 0.8);
  color: #1A1A2E;
}

.light-theme .nav-btn:hover {
  background: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .nav-btn {
  background: rgba(58, 58, 80, 0.8);
  color: #E4E4F0;
}

.nav-btn {
  padding: 8px 14px;
  border-radius: 16px;
  font-weight: 500;
  font-size: 14px;
  display: flex;
  align-items: center;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav-btn:hover {
  background: #6C63FF;
  color: #FFFFFF;
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(108, 99, 255, 0.3);
}

.nav-btn:active {
  transform: scale(0.98);
}

.logout-btn {
  background: rgba(255, 101, 132, 0.8);
}

.logout-btn:hover {
  background: #FF6584;
  color: #FFFFFF;
}
.login-btn {
  background: rgba(108, 99, 255, 0.8);
}
.login-btn:hover {
  background: #6C63FF;
  color: #FFFFFF;
}

.notification-dot {
  position: absolute;
  top: 6px;
  right: 12px;
  width: 8px;
  height: 8px;
  background: #FF6584;
  border-radius: 50%;
  box-shadow: 0 0 4px rgba(255, 101, 132, 0.6);
}
</style>