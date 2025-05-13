<template>
  <div>
    <div class="mb-6 flex flex-col md:flex-row gap-4 items-center justify-between">
      <div class="w-full md:w-1/2">
        <div class="relative">
          <input
            v-model="searchQuery"
            type="text"
            :class="['w-full py-2 px-4 rounded-xl shadow-sm border transition-all duration-200', 
                    isDarkMode ? 'bg-gray-800 border-gray-700 text-white' : 'bg-white border-gray-300 text-gray-900']"
            placeholder="Поиск мероприятий..."
          />
          <button class="absolute right-3 top-2 text-accent">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </button>
        </div>
      </div>

      <div class="flex gap-3 items-center">
        <select
          v-model="categoryFilter"
          :class="['py-2 px-4 rounded-xl shadow-sm border transition-all duration-200', 
                  isDarkMode ? 'bg-gray-800 border-gray-700 text-white' : 'bg-white border-gray-300 text-gray-900']"
        >
          <option value="">Все категории</option>
          <option value="tech">Технологии</option>
          <option value="music">Музыка</option>
          <option value="art">Искусство</option>
        </select>
        
        <button
          @click="toggleView"
          :class="['p-2 rounded-xl shadow-sm border transition-all duration-200', 
                  isDarkMode ? 'bg-gray-800 border-gray-700 text-white' : 'bg-white border-gray-300 text-gray-900']"
        >
          <svg v-if="isGridView" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
          <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          </svg>
        </button>
      </div>
    </div>
    
    <div v-if="filteredEvents.length === 0" class="text-center py-10">
      <p class="text-lg opacity-70">Мероприятия не найдены</p>
      <button @click="resetFilters" class="mt-3 text-accent underline">Сбросить фильтры</button>
    </div>
    
    <div v-else>
      <div v-if="isGridView" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <EventCard
          v-for="event in filteredEvents"
          :key="event.id"
          :event="event"
          :id="`event-${event.id}`"
          @click="$emit('select-event', event)"
        />
      </div>
      
      <div v-else class="space-y-4">
        <div
          v-for="event in filteredEvents"
          :key="event.id"
          :id="`event-${event.id}`"
          :class="['p-4 rounded-lg shadow-md hover:shadow-lg transition cursor-pointer', 
                  isDarkMode ? 'bg-card-dark bg-opacity-70' : 'bg-card-light bg-opacity-90']"
          @click="$emit('select-event', event)"
        >
          <div class="flex justify-between items-start">
            <div>
              <h3 class="text-lg font-semibold">{{ event.title }}</h3>
              <p class="text-sm opacity-80 mt-1">{{ formatDate(event.date) || 'Дата неизвестна' }}</p>
              <p v-if="event.description" class="mt-2">{{ event.description }}</p>
            </div>
            
            <div class="text-xs px-3 py-1 rounded-full" :class="getCategoryClass(event.category)">
              {{ event.category || 'Без категории' }}
            </div>
          </div>
          
          <div class="mt-3 flex gap-2 flex-wrap">
            <span v-for="tag in event.tags" :key="tag" 
                  :class="['px-2 py-1 rounded-full text-xs', 
                          isDarkMode ? 'bg-purple bg-opacity-50' : 'bg-purple bg-opacity-30']">
              {{ tag }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import EventCard from './EventCard.vue'

const props = defineProps({
  events: { type: Array, required: true, default: () => [] }
})

defineEmits(['select-event'])

const searchQuery = ref('')
const categoryFilter = ref('')
const isGridView = ref(true)

const theme = useState('theme')
const isDarkMode = computed(() => theme.value === 'dark')

const filteredEvents = computed(() => {
  let filtered = props.events
  
  if (searchQuery.value) {
    const query = searchQuery.value.toLowerCase()
    filtered = filtered.filter(event => 
      event.title?.toLowerCase().includes(query) || 
      event.description?.toLowerCase().includes(query)
    )
  }
  
  if (categoryFilter.value) {
    filtered = filtered.filter(event => event.category === categoryFilter.value)
  }
  
  return filtered
})

const toggleView = () => {
  isGridView.value = !isGridView.value
}

const resetFilters = () => {
  searchQuery.value = ''
  categoryFilter.value = ''
}

function getCategoryClass(category) {
  const categories = {
    'tech': 'bg-blue-500 bg-opacity-70 text-white',
    'music': 'bg-purple-500 bg-opacity-70 text-white',
    'art': 'bg-pink-500 bg-opacity-70 text-white',
    'default': isDarkMode.value ? 'bg-gray-700 text-gray-200' : 'bg-gray-300 text-gray-700'
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
</script> 