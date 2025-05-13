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
            <div class="flex items-center">
              <span class="opacity-80">{{ event.organizer }}</span>
              <button 
                @click="toggleSubscribe" 
                class="ml-2 text-xs px-2 py-1 rounded-full transition-colors"
                :class="isSubscribed ? 'bg-[#6C63FF]/20 text-[#6C63FF]' : 'bg-gray-200/30 hover:bg-[#6C63FF]/20'"
              >
                {{ isSubscribed ? 'Подписан' : 'Подписаться' }}
              </button>
            </div>
          </div>
          <div v-if="event.price !== undefined" class="text-sm">
            <div class="font-semibold mb-1">Стоимость</div>
            <div class="opacity-80">{{ formatPrice(event.price) }}</div>
          </div>
        </div>
        
        <!-- Rating -->
        <div class="mb-6">
          <h3 class="text-lg font-semibold mb-2">Оценка</h3>
          <div class="flex items-center">
            <div class="flex">
              <button 
                v-for="star in 5" 
                :key="star" 
                @click="setRating(star)"
                class="focus:outline-none"
              >
                <svg 
                  class="w-6 h-6" 
                  :class="star <= rating ? 'text-yellow-400' : 'text-gray-400'"
                  xmlns="http://www.w3.org/2000/svg" 
                  fill="currentColor" 
                  viewBox="0 0 24 24"
                >
                  <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z" />
                </svg>
              </button>
            </div>
            <span class="ml-2 text-sm">{{ displayRating }}</span>
          </div>
        </div>
        
        <!-- Comments -->
        <div class="mb-6">
          <h3 class="text-lg font-semibold mb-2">Комментарии ({{ comments.length }})</h3>
          
          <!-- Comment form -->
          <div class="mb-4">
            <form @submit.prevent="addComment">
              <textarea
                v-model="newComment"
                placeholder="Оставьте комментарий..."
                :class="['w-full p-2 rounded-lg resize-none transition-all', 
                       isDarkMode ? 'bg-[#3A3A50] text-white' : 'bg-gray-100 text-gray-800']"
                rows="2"
              ></textarea>
              <div class="flex justify-end mt-2">
                <button 
                  type="submit"
                  :disabled="!newComment.trim()"
                  :class="['px-4 py-2 rounded-lg text-sm font-medium transition-all',
                          newComment.trim() ? 'bg-[#6C63FF] text-white' : 'bg-gray-300 text-gray-500']"
                >
                  Отправить
                </button>
              </div>
            </form>
          </div>
          
          <!-- Comments list -->
          <div class="space-y-3 max-h-36 overflow-y-auto pr-2">
            <div 
              v-for="(comment, index) in comments" 
              :key="index"
              :class="['p-3 rounded-lg text-sm', 
                       isDarkMode ? 'bg-[#3A3A50]/50' : 'bg-gray-100']"
            >
              <div class="flex justify-between items-start">
                <span class="font-medium">{{ comment.author }}</span>
                <span class="text-xs opacity-60">{{ formatCommentDate(comment.date) }}</span>
              </div>
              <p class="mt-1">{{ comment.text }}</p>
            </div>
            <div v-if="comments.length === 0" class="text-center text-sm opacity-60 py-2">
              Пока нет комментариев
            </div>
          </div>
        </div>
        
        <!-- Share buttons -->
        <div class="mb-6">
          <h3 class="text-lg font-semibold mb-2">Поделиться</h3>
          <div class="flex space-x-3">
            <button 
              @click="shareEvent('vk')"
              class="p-2 rounded-lg transition-colors bg-[#4C75A3]/20 hover:bg-[#4C75A3]/30"
              title="Поделиться ВКонтакте"
            >
              <svg class="w-6 h-6 text-[#4C75A3]" viewBox="0 0 24 24" fill="currentColor">
                <path d="M21.547 7h-3.29a.743.743 0 0 0-.655.392s-1.312 2.416-1.734 3.23c-1.43 2.85-2.053 3.377-2.35 3.233-.565-.237-.494-1.9-.494-2.826V7.183c0-.575-.213-.783-.778-.783h-3.643c-.421 0-.678.267-.678.584 0 .541.82.717.893 2.26v3.253c.038.967-.166 1.179-.557 1.179-1.012 0-3.572-3.468-5.08-7.46C3.113 6.494 2.95 6 2.357 6H.803C.136 6 0 6.416 0 6.761c0 .6.096 3.627 4.537 7.85C7.237 17.056 10.775 18 13.927 18c1.82 0 2.112-.167 2.112-.908v-2.058c0-.67.142-.804.614-.804.37 0 1.002.176 2.476 1.607C20.857 17.446 21.046 18 21.895 18h3.291c.667 0 .998-.271.808-.817-.207-.632-2.436-2.855-2.539-2.988-.347-.407-.247-.631 0-1.024.173-.291 1.16-1.72 1.798-2.733.515-.814.902-1.493 1.097-1.926.184-.402.164-.662-.45-.902"></path>
              </svg>
            </button>
            
            <button 
              @click="shareEvent('telegram')"
              class="p-2 rounded-lg transition-colors bg-[#0088cc]/20 hover:bg-[#0088cc]/30"
              title="Поделиться в Telegram"
            >
              <svg class="w-6 h-6 text-[#0088cc]" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm4.64 6.8c-.15 1.58-.8 5.42-1.13 7.19-.14.75-.42 1-.68 1.03-.58.05-1.02-.38-1.58-.75-.88-.58-1.38-.94-2.23-1.5-.99-.65-.35-1.01.22-1.59.15-.15 2.71-2.48 2.76-2.69.01-.03.01-.14-.07-.2-.08-.06-.19-.04-.27-.02-.12.03-1.89 1.2-5.31 3.53-.5.34-.96.51-1.37.5-.45-.02-1.32-.26-1.97-.47-.8-.26-1.43-.4-1.37-.85.03-.22.27-.45.74-.68 2.87-1.25 4.79-2.07 5.76-2.46 2.74-1.12 3.31-1.31 3.69-1.32.08 0 .27.02.39.12.1.08.13.19.14.27-.01.06.01.24 0 .38z"/>
              </svg>
            </button>
            
            <button 
              @click="shareEvent('whatsapp')"
              class="p-2 rounded-lg transition-colors bg-[#25D366]/20 hover:bg-[#25D366]/30"
              title="Поделиться в WhatsApp"
            >
              <svg class="w-6 h-6 text-[#25D366]" viewBox="0 0 24 24" fill="currentColor">
                <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
              </svg>
            </button>
            
            <button 
              @click="copyEventLink"
              class="p-2 rounded-lg transition-colors bg-gray-400/20 hover:bg-gray-400/30"
              title="Копировать ссылку"
            >
              <svg class="w-6 h-6 text-gray-500" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
              </svg>
            </button>
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
    
    <!-- Notification -->
    <div 
      v-if="showNotification" 
      class="fixed bottom-6 left-1/2 transform -translate-x-1/2 bg-[#6C63FF] text-white px-6 py-4 rounded-xl shadow-lg z-50 animate-fade-up"
    >
      {{ notificationMessage }}
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

// Favorites logic
const favorites = ref([])
const isFavorite = ref(false)

// Subscription logic
const subscriptions = ref([])
const isSubscribed = ref(false)

// Rating logic
const rating = ref(0)
const displayRating = computed(() => {
  return rating.value ? `${rating.value}/5` : 'Нет оценки'
})

// Comments logic
const comments = ref([])
const newComment = ref('')

// Notification logic
const showNotification = ref(false)
const notificationMessage = ref('')

// Load data from localStorage on client-side
if (process.client) {
  try {
    // Load favorites
    const savedFavorites = localStorage.getItem('favorites')
    if (savedFavorites) {
      favorites.value = JSON.parse(savedFavorites)
    }
    
    // Load subscriptions
    const savedSubscriptions = localStorage.getItem('subscriptions')
    if (savedSubscriptions) {
      subscriptions.value = JSON.parse(savedSubscriptions)
    }
    
    // Load comments
    const savedComments = localStorage.getItem('eventComments')
    if (savedComments) {
      const allComments = JSON.parse(savedComments)
      if (props.event?.id && allComments[props.event.id]) {
        comments.value = allComments[props.event.id]
      }
    }
    
    // Load rating
    const savedRatings = localStorage.getItem('eventRatings')
    if (savedRatings && props.event?.id) {
      const allRatings = JSON.parse(savedRatings)
      if (allRatings[props.event.id]) {
        rating.value = allRatings[props.event.id]
      }
    }
  } catch (e) {
    console.error('Error loading data from localStorage:', e)
  }
}

// Update state when event changes
watch(() => props.event, (newEvent) => {
  if (!newEvent) return
  
  // Check if event is in favorites
  if (favorites.value) {
    isFavorite.value = favorites.value.some(fav => fav?.id === newEvent.id)
  }
  
  // Check if user is subscribed to organizer
  if (subscriptions.value && newEvent.organizer) {
    isSubscribed.value = subscriptions.value.includes(newEvent.organizer)
  }
  
  // Load comments for this event
  if (process.client) {
    try {
      const savedComments = localStorage.getItem('eventComments')
      if (savedComments) {
        const allComments = JSON.parse(savedComments)
        comments.value = allComments[newEvent.id] || []
      } else {
        comments.value = []
      }
      
      // Load rating for this event
      const savedRatings = localStorage.getItem('eventRatings')
      if (savedRatings) {
        const allRatings = JSON.parse(savedRatings)
        rating.value = allRatings[newEvent.id] || 0
      } else {
        rating.value = 0
      }
    } catch (e) {
      console.error('Error loading data for event', e)
    }
  }
}, { immediate: true })

// Toggle event favorite status
const toggleFavorite = () => {
  if (!props.event) return
  
  const wasFavorite = isFavorite.value
  if (wasFavorite) {
    favorites.value = favorites.value.filter(fav => fav?.id !== props.event.id)
  } else {
    favorites.value.push(props.event)
  }
  
  isFavorite.value = !wasFavorite
  
  if (process.client) {
    localStorage.setItem('favorites', JSON.stringify(favorites.value))
  }
  
  emit('toggleFavorite', { favorites: favorites.value, isAdded: !wasFavorite })
}

// Toggle subscription to organizer
const toggleSubscribe = () => {
  if (!props.event?.organizer) return
  
  // Simulate server request
  showNotification.value = true
  notificationMessage.value = 'Обработка запроса...'
  
  setTimeout(() => {
    const wasSubscribed = isSubscribed.value
    
    if (wasSubscribed) {
      subscriptions.value = subscriptions.value.filter(org => org !== props.event.organizer)
      notificationMessage.value = `Вы отписались от организатора "${props.event.organizer}"`
    } else {
      subscriptions.value.push(props.event.organizer)
      notificationMessage.value = `Вы подписались на организатора "${props.event.organizer}"`
    }
    
    isSubscribed.value = !wasSubscribed
    
    if (process.client) {
      localStorage.setItem('subscriptions', JSON.stringify(subscriptions.value))
    }
    
    setTimeout(() => {
      showNotification.value = false
    }, 3000)
  }, 1000)
}

// Set event rating
const setRating = (value) => {
  if (!props.event?.id) return
  
  // Simulate server request
  showNotification.value = true
  notificationMessage.value = 'Сохранение оценки...'
  
  setTimeout(() => {
    rating.value = value
    
    if (process.client) {
      try {
        const savedRatings = localStorage.getItem('eventRatings')
        const allRatings = savedRatings ? JSON.parse(savedRatings) : {}
        
        allRatings[props.event.id] = value
        localStorage.setItem('eventRatings', JSON.stringify(allRatings))
        
        notificationMessage.value = `Вы поставили оценку ${value}/5`
      } catch (e) {
        console.error('Error saving rating', e)
        notificationMessage.value = 'Ошибка при сохранении оценки'
      }
    }
    
    setTimeout(() => {
      showNotification.value = false
    }, 2000)
  }, 1000)
}

// Add a comment to the event
const addComment = () => {
  if (!props.event?.id || !newComment.trim()) return
  
  // Create the comment
  const comment = {
    text: newComment.trim(),
    author: 'Вы',
    date: new Date().toISOString()
  }
  
  // Add to comments array
  comments.value.unshift(comment)
  newComment.value = ''
  
  // Save to localStorage
  if (process.client) {
    try {
      const savedComments = localStorage.getItem('eventComments')
      const allComments = savedComments ? JSON.parse(savedComments) : {}
      
      allComments[props.event.id] = comments.value
      localStorage.setItem('eventComments', JSON.stringify(allComments))
      
      // Show success notification
      showNotification.value = true
      notificationMessage.value = 'Комментарий добавлен'
      
      setTimeout(() => {
        showNotification.value = false
      }, 2000)
    } catch (e) {
      console.error('Error saving comment', e)
      showNotification.value = true
      notificationMessage.value = 'Ошибка при сохранении комментария'
      setTimeout(() => {
        showNotification.value = false
      }, 2000)
    }
  }
}

// Share event in social networks
const shareEvent = (network) => {
  if (!props.event?.id) return
  
  // Create share URL
  const eventUrl = `${window.location.origin}/event/${props.event.id}`
  const shareTitle = props.event.title
  const shareText = props.event.description?.substring(0, 100) || 'Посмотрите это мероприятие!'
  
  let shareUrl = ''
  
  switch (network) {
    case 'vk':
      shareUrl = `https://vk.com/share.php?url=${encodeURIComponent(eventUrl)}&title=${encodeURIComponent(shareTitle)}`
      break
    case 'telegram':
      shareUrl = `https://t.me/share/url?url=${encodeURIComponent(eventUrl)}&text=${encodeURIComponent(shareTitle)}`
      break
    case 'whatsapp':
      shareUrl = `https://wa.me/?text=${encodeURIComponent(shareTitle + ' ' + eventUrl)}`
      break
  }
  
  // Simulate share opening
  showNotification.value = true
  notificationMessage.value = 'Открываем окно для шаринга...'
  
  setTimeout(() => {
    showNotification.value = true
    notificationMessage.value = `Ссылка скопирована. Делитесь в ${network}!`
    
    // In real implementation we would open a window:
    // window.open(shareUrl, '_blank')
    
    setTimeout(() => {
      showNotification.value = false
    }, 2000)
  }, 1000)
}

// Copy event link to clipboard
const copyEventLink = () => {
  if (!props.event?.id) return
  
  const eventUrl = `${window.location.origin}/event/${props.event.id}`
  
  // Simulate copy to clipboard
  showNotification.value = true
  notificationMessage.value = 'Копирование ссылки...'
  
  setTimeout(() => {
    // In a real app, we'd use clipboard API:
    // navigator.clipboard.writeText(eventUrl)
    
    notificationMessage.value = 'Ссылка скопирована в буфер обмена!'
    
    setTimeout(() => {
      showNotification.value = false
    }, 2000)
  }, 1000)
}

// Format comment date
const formatCommentDate = (dateString) => {
  if (!dateString) return ''
  
  try {
    const date = new Date(dateString)
    const now = new Date()
    const diffMs = now - date
    const diffSeconds = Math.floor(diffMs / 1000)
    const diffMinutes = Math.floor(diffSeconds / 60)
    const diffHours = Math.floor(diffMinutes / 60)
    const diffDays = Math.floor(diffHours / 24)
    
    if (diffSeconds < 60) {
      return 'только что'
    } else if (diffMinutes < 60) {
      return `${diffMinutes} мин. назад`
    } else if (diffHours < 24) {
      return `${diffHours} ч. назад`
    } else {
      return date.toLocaleDateString('ru-RU')
    }
  } catch {
    return dateString
  }
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
  color: white;
  transition: all 0.3s ease;
}

.action-btn:hover {
  background-color: #5b54e0;
}

.cancel-btn {
  background-color: #FF6584;
  color: white;
  transition: all 0.3s ease;
}

.cancel-btn:hover {
  background-color: #e05574;
}

.register-btn {
  background-color: #35A7FF;
  color: white;
  transition: all 0.3s ease;
}

.register-btn:hover {
  background-color: #2b86cc;
}

@keyframes fadeUp {
  0% {
    opacity: 0;
    transform: translate(-50%, 20px);
  }
  100% {
    opacity: 1;
    transform: translate(-50%, 0);
  }
}

.animate-fade-up {
  animation: fadeUp 0.3s ease-out forwards;
}
</style>