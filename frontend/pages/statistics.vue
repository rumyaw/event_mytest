<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']">
    <Navbar />
    
    <!-- Фон с падающими звездами как на главной странице -->
    <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
      <div v-for="i in 30" :key="i" class="falling-star" :style="randomFallingStarStyle()" />
    </div>

    <div class="container mx-auto p-6 space-y-8 relative z-10 pt-20">
      <div class="text-center mb-8">
        <h1 class="text-3xl font-bold text-accent">Статистика мероприятий</h1>
        <p class="text-opacity-80">Подробная аналитика по всем событиям и организаторам</p>
      </div>

      <!-- Dashboard Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        <div class="p-6 rounded-xl shadow-lg events-section">
          <h3 class="text-sm font-medium text-opacity-70 mb-1">Текущие просмотры</h3>
          <div class="flex items-end gap-2">
            <p class="text-3xl font-bold text-accent">{{ currentViews }}</p>
            <p :class="['text-sm', viewsChange >= 0 ? 'text-green-500' : 'text-red-500']">
              {{ viewsChange >= 0 ? '+' : '' }}{{ viewsChange }}%
            </p>
          </div>
          <div class="flex items-center gap-2 text-sm mt-2">
            <div class="flex h-2 bg-gray-200 dark:bg-gray-700 rounded-full w-full overflow-hidden">
              <div class="bg-gradient-to-r from-purple-500 to-blue-500 h-full rounded-full" :style="`width: ${viewsPercentage}%`"></div>
            </div>
            <span class="text-xs">{{ viewsPercentage }}%</span>
          </div>
        </div>
        
        <div class="p-6 rounded-xl shadow-lg events-section">
          <h3 class="text-sm font-medium text-opacity-70 mb-1">Общее количество просмотров</h3>
          <div class="flex items-end gap-2">
            <p class="text-3xl font-bold text-accent">{{ totalViews }}</p>
            <p :class="['text-sm', totalViewsChange >= 0 ? 'text-green-500' : 'text-red-500']">
              {{ totalViewsChange >= 0 ? '+' : '' }}{{ totalViewsChange }}%
            </p>
          </div>
          <div class="flex items-center gap-2 text-sm mt-2">
            <div class="flex h-2 bg-gray-200 dark:bg-gray-700 rounded-full w-full overflow-hidden">
              <div class="bg-gradient-to-r from-blue-500 to-cyan-500 h-full rounded-full" :style="`width: ${totalViewsPercentage}%`"></div>
            </div>
            <span class="text-xs">{{ totalViewsPercentage }}%</span>
          </div>
        </div>
        
        <div class="p-6 rounded-xl shadow-lg events-section">
          <h3 class="text-sm font-medium text-opacity-70 mb-1">Записавшиеся пользователи</h3>
          <div class="flex items-end gap-2">
            <p class="text-3xl font-bold text-accent">{{ registeredUsers }}</p>
            <p :class="['text-sm', registeredUsersChange >= 0 ? 'text-green-500' : 'text-red-500']">
              {{ registeredUsersChange >= 0 ? '+' : '' }}{{ registeredUsersChange }}%
            </p>
          </div>
          <div class="flex items-center gap-2 text-sm mt-2">
            <div class="flex h-2 bg-gray-200 dark:bg-gray-700 rounded-full w-full overflow-hidden">
              <div class="bg-gradient-to-r from-green-500 to-emerald-500 h-full rounded-full" :style="`width: ${registeredUsersPercentage}%`"></div>
            </div>
            <span class="text-xs">{{ registeredUsersPercentage }}%</span>
          </div>
        </div>
        
        <div class="p-6 rounded-xl shadow-lg events-section">
          <h3 class="text-sm font-medium text-opacity-70 mb-1">В избранном</h3>
          <div class="flex items-end gap-2">
            <p class="text-3xl font-bold text-accent">{{ favoriteCount }}</p>
            <p :class="['text-sm', favoritesChange >= 0 ? 'text-green-500' : 'text-red-500']">
              {{ favoritesChange >= 0 ? '+' : '' }}{{ favoritesChange }}%
            </p>
          </div>
          <div class="flex items-center gap-2 text-sm mt-2">
            <div class="flex h-2 bg-gray-200 dark:bg-gray-700 rounded-full w-full overflow-hidden">
              <div class="bg-gradient-to-r from-yellow-500 to-amber-500 h-full rounded-full" :style="`width: ${favoritesPercentage}%`"></div>
            </div>
            <span class="text-xs">{{ favoritesPercentage }}%</span>
          </div>
        </div>
      </div>

      <!-- Statistics Content -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Main content column -->
        <div class="lg:col-span-2 space-y-6">
          <!-- Audience Engagement Chart -->
          <div class="p-6 rounded-xl shadow-lg events-section">
            <div class="flex justify-between items-center mb-6">
              <h2 class="text-xl font-bold text-accent">Посещаемость мероприятий</h2>
              <div class="flex gap-2">
                <button 
                  v-for="period in chartPeriods" 
                  :key="period.id"
                  @click="selectedChartPeriod = period.id"
                  :class="['px-3 py-1 text-xs rounded-lg transition-colors', 
                    selectedChartPeriod === period.id ? 'bg-blue-500 text-white' : 'bg-opacity-20 hover:bg-opacity-30']"
                >
                  {{ period.label }}
                </button>
              </div>
            </div>
            <div class="h-64 relative chart-placeholder p-4">
              <!-- Имитация графика с колонками -->
              <div class="absolute inset-0 flex items-end justify-between px-8 pb-10">
                <div v-for="(column, i) in chartColumns" :key="i" class="h-full flex items-end justify-center" style="width: 8%;">
                  <div 
                    class="bg-gradient-to-t from-blue-500 to-purple-500 rounded-t-md w-full transition-all duration-1000"
                    :style="`height: ${column}%`">
                  </div>
                </div>
              </div>
              
              <!-- Подписи по оси X -->
              <div class="absolute bottom-2 left-8 right-8 flex justify-between">
                <div v-for="(label, i) in chartLabels" :key="i" class="text-xs text-center" style="width: 8%;">
                  {{ label }}
                </div>
              </div>
              
              <!-- Подписи по оси Y -->
              <div class="absolute left-2 top-0 bottom-10 flex flex-col justify-between">
                <div v-for="value in 6" :key="value" class="text-xs">
                  {{ (100 - (value - 1) * 20) }}
                </div>
              </div>
              
              <!-- Имитация осей -->
              <div class="absolute bottom-8 left-0 right-0 h-px bg-gray-300 dark:bg-gray-600"></div>
              <div class="absolute top-0 bottom-0 left-6 w-px bg-gray-300 dark:bg-gray-600"></div>
              
              <!-- Линии сетки -->
              <div v-for="line in 5" :key="line" class="absolute left-6 right-0 h-px bg-gray-200 dark:bg-gray-700 opacity-30" 
                   :style="`bottom: ${8 + (line * 20 * (56/100))}%`">
              </div>
            </div>
          </div>

          <!-- Popular Events Table -->
          <div class="p-6 rounded-xl shadow-lg events-section">
            <h2 class="text-xl font-bold text-accent mb-4">Популярные мероприятия</h2>
            <div class="overflow-x-auto">
              <table class="w-full">
                <thead>
                  <tr>
                    <th class="text-left pb-3 text-sm font-medium text-opacity-70">Событие</th>
                    <th class="text-left pb-3 text-sm font-medium text-opacity-70">Организатор</th>
                    <th class="text-left pb-3 text-sm font-medium text-opacity-70">Просмотры</th>
                    <th class="text-left pb-3 text-sm font-medium text-opacity-70">Участники</th>
                    <th class="text-left pb-3 text-sm font-medium text-opacity-70">В избранном</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="event in popularEvents" :key="event.id" class="hover:bg-opacity-5 hover:bg-blue-500 transition-colors">
                    <td class="py-3 text-sm">{{ event.title }}</td>
                    <td class="py-3 text-sm">{{ event.organizer }}</td>
                    <td class="py-3 text-sm">{{ event.views }}</td>
                    <td class="py-3 text-sm">{{ event.participants }}</td>
                    <td class="py-3 text-sm">{{ event.favorites }}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Right column with additional stats -->
        <div class="space-y-6">
          <!-- Top Organizers -->
          <div class="p-6 rounded-xl shadow-lg events-section">
            <h2 class="text-xl font-bold text-accent mb-4">Лучшие организаторы</h2>
            <div class="space-y-4">
              <div v-for="organizer in topOrganizers" :key="organizer.id" class="flex items-center gap-4">
                <div class="flex-shrink-0 h-12 w-12 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-white font-bold">
                  {{ organizer.initials }}
                </div>
                <div class="flex-1 flex flex-col">
                  <span class="font-medium">{{ organizer.name }}</span>
                  <span class="text-xs">{{ organizer.events }} мероприятий</span>
                  <div class="w-full mt-1 bg-gray-200 dark:bg-gray-700 rounded-full h-2">
                    <div class="bg-gradient-to-r from-purple-500 to-blue-500 h-2 rounded-full" :style="`width: ${organizer.rating}%`"></div>
                  </div>
                </div>
                <div class="text-xl font-bold text-accent">{{ organizer.rating }}%</div>
              </div>
            </div>
          </div>

          <!-- Demographics & Categories -->
          <div class="p-6 rounded-xl shadow-lg events-section">
            <h2 class="text-xl font-bold text-accent mb-4">Категории</h2>
            <div class="space-y-3">
              <div v-for="(category, index) in categories" :key="index" class="flex flex-col">
                <div class="flex justify-between mb-1">
                  <span class="text-sm">{{ category.name }}</span>
                  <span class="text-sm font-medium">{{ category.percentage }}%</span>
                </div>
                <div class="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2">
                  <div 
                    class="h-2 rounded-full" 
                    :class="category.colorClass"
                    :style="`width: ${category.percentage}%`"
                  ></div>
                </div>
              </div>
            </div>
          </div>
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
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'
import Navbar from '@/components/Navbar.vue'

// Theme state
const theme = useState('theme')
const isDarkMode = computed(() => theme.value === 'dark')

// Dashboard metrics with random initial values
const currentViews = ref(Math.floor(Math.random() * 100) + 50)
const totalViews = ref(Math.floor(Math.random() * 5000) + 1000)
const registeredUsers = ref(Math.floor(Math.random() * 200) + 50)
const favoriteCount = ref(Math.floor(Math.random() * 300) + 100)

// Change percentages (random values between -15 and +25)
const viewsChange = ref(Math.floor(Math.random() * 40) - 15)
const totalViewsChange = ref(Math.floor(Math.random() * 40) - 15)
const registeredUsersChange = ref(Math.floor(Math.random() * 40) - 15)
const favoritesChange = ref(Math.floor(Math.random() * 40) - 15)

// Percentage bars (calculated to look good in UI, not actual percentages)
const viewsPercentage = ref(Math.floor(Math.random() * 40) + 40)
const totalViewsPercentage = ref(Math.floor(Math.random() * 40) + 40)
const registeredUsersPercentage = ref(Math.floor(Math.random() * 40) + 40)
const favoritesPercentage = ref(Math.floor(Math.random() * 40) + 40)

// Chart data
const chartPeriods = [
  { id: 'day', label: 'День' },
  { id: 'week', label: 'Неделя' },
  { id: 'month', label: 'Месяц' },
  { id: 'year', label: 'Год' }
]
const selectedChartPeriod = ref('week')

// Подписи для оси X (дни недели)
const chartLabels = ref(['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс', 'Пн', 'Вт', 'Ср'])

// Simulate 10 columns for the chart with random values
const chartColumns = ref(Array.from({ length: 10 }, () => Math.floor(Math.random() * 80) + 10))

// Popular events with random data
const popularEvents = ref([
  { id: 1, title: 'Фестиваль музыки "Мелодия лета"', organizer: 'ИвентПро', views: 1842, participants: 256, favorites: 124 },
  { id: 2, title: 'Выставка современного искусства', organizer: 'Арт-Галерея', views: 1254, participants: 187, favorites: 98 },
  { id: 3, title: 'Мастер-класс по программированию', organizer: 'TechHub', views: 987, participants: 145, favorites: 76 },
  { id: 4, title: 'Конференция маркетологов', organizer: 'МаркетингПлюс', views: 864, participants: 112, favorites: 62 },
  { id: 5, title: 'Спортивный марафон "Здоровье"', organizer: 'СпортТайм', views: 743, participants: 98, favorites: 54 }
])

// Top organizers with random data
const topOrganizers = ref([
  { id: 1, name: 'ИвентПро', initials: 'ИП', events: 15, rating: 92 },
  { id: 2, name: 'Арт-Галерея', initials: 'АГ', events: 12, rating: 87 },
  { id: 3, name: 'TechHub', initials: 'TH', events: 10, rating: 84 },
  { id: 4, name: 'МаркетингПлюс', initials: 'МП', events: 8, rating: 78 }
])

// Categories with random data
const categories = ref([
  { name: 'Музыка и концерты', percentage: 28, colorClass: 'bg-gradient-to-r from-blue-500 to-cyan-500' },
  { name: 'Выставки и искусство', percentage: 22, colorClass: 'bg-gradient-to-r from-purple-500 to-pink-500' },
  { name: 'IT и технологии', percentage: 18, colorClass: 'bg-gradient-to-r from-green-500 to-emerald-500' },
  { name: 'Спорт и здоровье', percentage: 16, colorClass: 'bg-gradient-to-r from-yellow-500 to-amber-500' },
  { name: 'Образование', percentage: 10, colorClass: 'bg-gradient-to-r from-red-500 to-orange-500' },
  { name: 'Другое', percentage: 6, colorClass: 'bg-gradient-to-r from-gray-500 to-slate-500' }
])

// Update chart data every 5 seconds
let chartUpdateInterval
let metricsUpdateInterval

onMounted(() => {
  // Обновление графика
  chartUpdateInterval = setInterval(() => {
    // Генерируем новые значения для столбцов
    chartColumns.value = chartColumns.value.map(() => Math.floor(Math.random() * 80) + 10)
    
    // Обновляем подписи в зависимости от выбранного периода
    if (selectedChartPeriod.value === 'day') {
      chartLabels.value = ['9:00', '10:00', '11:00', '12:00', '13:00', '14:00', '15:00', '16:00', '17:00', '18:00']
    } else if (selectedChartPeriod.value === 'week') {
      chartLabels.value = ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс', 'Пн', 'Вт', 'Ср']
    } else if (selectedChartPeriod.value === 'month') {
      chartLabels.value = ['1', '5', '10', '15', '20', '25', '30', '5', '10', '15']
    } else {
      chartLabels.value = ['Янв', 'Фев', 'Мар', 'Апр', 'Май', 'Июн', 'Июл', 'Авг', 'Сен', 'Окт']
    }
  }, 5000)
  
  // Обновление метрик в реальном времени
  metricsUpdateInterval = setInterval(() => {
    // Обновляем текущие просмотры (+/- случайное число от -5 до +10)
    const viewsDelta = Math.floor(Math.random() * 16) - 5
    currentViews.value = Math.max(20, currentViews.value + viewsDelta)
    
    // Обновляем общее количество просмотров (всегда растет)
    const totalDelta = Math.floor(Math.random() * 20) + 1
    totalViews.value += totalDelta
    
    // Случайно обновляем другие метрики
    if (Math.random() > 0.7) {
      registeredUsers.value += Math.floor(Math.random() * 3)
    }
    
    if (Math.random() > 0.8) {
      favoriteCount.value += Math.floor(Math.random() * 2)
    }
    
    // Обновляем процентные изменения
    viewsChange.value = Math.floor(Math.random() * 40) - 15
    
    // Случайно обновляем Top-5 популярных мероприятий
    if (Math.random() > 0.8) {
      popularEvents.value = popularEvents.value.map(event => ({
        ...event,
        views: event.views + Math.floor(Math.random() * 10),
        participants: event.participants + (Math.random() > 0.7 ? 1 : 0),
        favorites: event.favorites + (Math.random() > 0.8 ? 1 : 0)
      }))
      
      // Сортируем по просмотрам
      popularEvents.value.sort((a, b) => b.views - a.views)
    }
  }, 3000)
})

onUnmounted(() => {
  // Очищаем интервалы при размонтировании компонента
  clearInterval(chartUpdateInterval)
  clearInterval(metricsUpdateInterval)
})

function toggleTheme() {
  theme.value = isDarkMode.value ? 'light' : 'dark'
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

// Наблюдатель за изменением периода
watch(selectedChartPeriod, (newPeriod) => {
  // Обновляем подписи в зависимости от выбранного периода
  if (newPeriod === 'day') {
    chartLabels.value = ['9:00', '10:00', '11:00', '12:00', '13:00', '14:00', '15:00', '16:00', '17:00', '18:00']
  } else if (newPeriod === 'week') {
    chartLabels.value = ['Пн', 'Вт', 'Ср', 'Чт', 'Пт', 'Сб', 'Вс', 'Пн', 'Вт', 'Ср']
  } else if (newPeriod === 'month') {
    chartLabels.value = ['1', '5', '10', '15', '20', '25', '30', '5', '10', '15']
  } else {
    chartLabels.value = ['Янв', 'Фев', 'Мар', 'Апр', 'Май', 'Июн', 'Июл', 'Авг', 'Сен', 'Окт']
  }
  
  // Генерируем новые случайные значения для столбцов
  chartColumns.value = chartColumns.value.map(() => Math.floor(Math.random() * 80) + 10)
});
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

.light-theme .theme-toggle {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.light-theme .text-accent {
  color: #6C63FF;
}

/* Dark Theme */
.dark-theme {
  background-color: #1A1A2E;
  color: #E4E4F0;
}

.dark-theme .events-section {
  background-color: #2A2A40;
}

.dark-theme .theme-toggle {
  background-color: #6C63FF;
  color: #FFFFFF;
}

.dark-theme .text-accent {
  color: #6C63FF;
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

/* Chart placeholder */
.chart-placeholder {
  border-bottom: 1px solid rgba(128, 128, 128, 0.3);
  border-left: 1px solid rgba(128, 128, 128, 0.3);
  border-radius: 0.5rem;
  background-color: rgba(128, 128, 128, 0.05);
}
</style> 