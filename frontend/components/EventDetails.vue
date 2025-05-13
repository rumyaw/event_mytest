<!-- components/EventDetails.vue -->
<template>
  <div class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-[1000] p-4" @click.self="$emit('close')">
    <div :class="['rounded-2xl shadow-lg max-w-lg w-full overflow-hidden', isDarkMode ? 'dark-theme-bg' : 'light-theme-bg']">
      <!-- Header with close button -->
      <div class="relative">
        <div class="absolute top-0 left-0 w-full h-32 bg-gradient-to-b from-purple-600 to-transparent opacity-30"></div>
        <div class="flex justify-between items-start p-6 relative z-10">
          <div>
            <span :class="['text-xs px-3 py-1 rounded-full', getCategoryClass(event?.category)]">
              {{ event?.category || 'Без категории' }}
            </span>
            <h2 class="text-2xl font-semibold mt-2">{{ event?.title }}</h2>
          </div>
          <button @click="$emit('close')" class="p-2 rounded-full hover:bg-opacity-10 hover:bg-black transition-all">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
      
      <!-- Content -->
      <div class="px-6 pb-6" v-if="event">
        <!-- Date and location -->
        <div class="flex items-center mb-4 text-sm opacity-80">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          <span>{{ formatDate(event.date) || 'Дата неизвестна' }}</span>
        </div>
        
        <div v-if="event.location" class="flex items-center mb-4 text-sm opacity-80">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
          <span>{{ event.location.city }}, {{ event.location.venue }}</span>
        </div>
        
        <div v-if="event.lat && event.lng" class="flex items-center mb-4 text-sm opacity-80">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20l-5.447-2.724A1 1 0 013 16.382V5.618a1 1 0 011.447-.894L9 7m0 13l6-3m-6 3V7m6 10l4.553 2.276A1 1 0 0021 18.382V7.618a1 1 0 00-.553-.894L15 4m0 13V4m0 0L9 7" />
          </svg>
          <span>Координаты: {{ event.lat.toFixed(4) }}, {{ event.lng.toFixed(4) }}</span>
        </div>
        
        <!-- Description -->
        <div class="my-6">
          <h3 class="text-lg font-semibold mb-2">Описание</h3>
          <p class="opacity-80">{{ event.description || 'Описание отсутствует' }}</p>
        </div>
        
        <!-- Tags -->
        <div v-if="event.tags && event.tags.length" class="mb-6">
          <h3 class="text-lg font-semibold mb-2">Теги</h3>
          <div class="flex flex-wrap gap-2">
            <span v-for="tag in event.tags" :key="tag" 
                  :class="['px-3 py-1 rounded-full text-xs', 
                          isDarkMode ? 'bg-purple bg-opacity-50' : 'bg-purple bg-opacity-30']">
              {{ tag }}
            </span>
          </div>
        </div>
        
        <!-- Organizer and price -->
        <div class="grid grid-cols-2 gap-4 mb-6">
          <div v-if="event.organizer" class="text-sm">
            <div class="font-semibold mb-1">Организатор</div>
            <div class="opacity-80">{{ event.organizer }}</div>
          </div>
          <div v-if="event.price !== undefined" class="text-sm">
            <div class="font-semibold mb-1">Стоимость</div>
            <div class="opacity-80">{{ formatPrice(event.price) }}</div>
          </div>
        </div>
        
        <!-- Actions -->
        <div class="flex gap-4">
          <button
            :class="['p-3 rounded-xl flex-1 flex items-center justify-center gap-2', isFavorite ? 'cancel-btn' : 'action-btn']"
            @click="toggleFavorite"
          >
            <svg v-if="isFavorite" class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24">
              <path d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z" />
            </svg>
            <svg v-else class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
            {{ isFavorite ? 'Удалить из избранного' : 'В избранное' }}
          </button>
          
          <button class="p-3 rounded-xl flex-1 register-btn flex items-center justify-center gap-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1" />
            </svg>
            Зарегистрироваться
          </button>
        </div>
      </div>
      
      <!-- Error state -->
      <div class="p-6" v-else>
        <p class="text-red-500">Событие не выбрано</p>
        <button class="mt-4 p-3 rounded-xl w-full cancel-btn" @click="$emit('close')">Закрыть</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, computed } from 'vue'

const props = defineProps({
  event: Object,
  isDarkMode: Boolean
})

const emit = defineEmits(['close', 'toggleFavorite'])

const favorites = ref([])
const isFavorite = ref(false)

// Load favorites from localStorage on client-side
if (process.client) {
  try {
    const saved = localStorage.getItem('favorites')
    if (saved) {
      favorites.value = JSON.parse(saved)
    }
  } catch (e) {
    console.error('Error loading favorites:', e)
  }
}

watch(() => props.event, (newEvent) => {
  if (newEvent && favorites.value) {
    isFavorite.value = favorites.value.some(fav => fav?.id === newEvent.id)
  }
}, { immediate: true })

const toggleFavorite = () => {
  if (!props.event) return
  
  const wasFavorite = isFavorite.value
  if (wasFavorite) {
    favorites.value = favorites.value.filter(fav => fav?.id !== props.event.id)
  } else {
    favorites.value.push(props.event)
  }
  
  isFavorite.value = !isFavorite.value
  
  if (process.client) {
    localStorage.setItem('favorites', JSON.stringify(favorites.value))
  }
  
  emit('toggleFavorite', { favorites: favorites.value, isAdded: !wasFavorite })
}

function getCategoryClass(category) {
  const categories = {
    'tech': 'bg-blue-500 bg-opacity-70 text-white',
    'music': 'bg-purple-500 bg-opacity-70 text-white',
    'art': 'bg-pink-500 bg-opacity-70 text-white',
    'default': props.isDarkMode ? 'bg-gray-700 text-gray-200' : 'bg-gray-300 text-gray-700'
  }
  
  return categories[category] || categories.default
}

function formatDate(dateString) {
  if (!dateString) return null
  
  try {
    const date = new Date(dateString)
    return date.toLocaleDateString('ru-RU', {
      day: 'numeric',
      month: 'long',
      year: 'numeric',
      hour: '2-digit', 
      minute: '2-digit'
    })
  } catch {
    return dateString
  }
}

function formatPrice(price) {
  if (price === 0) return 'Бесплатно'
  return `${price.toLocaleString('ru-RU')} ₽`
}
</script>

<style scoped>
.dark-theme-bg {
  background-color: #2A2A40;
  color: #E4E4F0;
}

.light-theme-bg {
  background-color: #FFFFFF;
  color: #1A1A2E;
}

.action-btn {
  background-color: #6C63FF;
  color: #FFFFFF;
  transition: all 0.3s ease;
}

.register-btn {
  background-color: #4CAF50;
  color: #FFFFFF;
  transition: all 0.3s ease;
}

.cancel-btn {
  background-color: #3A3A50;
  color: #E4E4F0;
  transition: all 0.3s ease;
}

.light-theme-bg .cancel-btn {
  background-color: #F0F0F5;
  color: #1A1A2E;
}

.action-btn:hover, .cancel-btn:hover, .register-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.text-accent {
  color: #6C63FF;
}
</style>