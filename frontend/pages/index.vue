<!-- pages/index.vue -->
<template>
  <div :class="['min-h-screen relative', isDarkMode ? 'dark-theme' : 'light-theme']">
    <Navbar />

    <!-- Уведомления -->
    <div
      v-if="showSuccess"
      :class="['fixed top-[64px] right-6 text-[#FFFFFF] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-[150] transition-colors', 
               isDarkMode ? 'bg-[#5A52D9]' : 'bg-[#6C63FF]']"
    >
      ✅ Добавлено в избранное!
    </div>

    <div
      v-if="showRemoveSuccess"
      :class="['fixed top-[64px] right-6 text-[#FFFFFF] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-[150] transition-colors',
               isDarkMode ? 'bg-[#E05574]' : 'bg-[#FF6584]']"
    >
      ❌ Удалено из избранного!
    </div>

    <!-- Фон с падающими звездами -->
    <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
      <div v-for="i in 30" :key="i" class="falling-star" :style="randomFallingStarStyle()" />
    </div>

    <div class="container mx-auto p-6 space-y-8 relative z-10 pt-20">
      <!-- Заголовок и краткое описание -->
      <div class="text-center space-y-2">
        <h1 class="text-3xl font-bold text-accent">Мероприятия на карте</h1>
        <p class="text-opacity-80">Найди интересные события рядом с тобой или в любой точке мира!</p>
      </div>

      <!-- Карта -->
      <div class="rounded-2xl shadow-lg overflow-hidden h-[500px]">
        <ClientOnly>
          <Map v-if="isClient" :events="events" :selected-event-id="selectedEventId" @update:selected="onUpdateSelected" />
          <template #fallback>
            <div class="h-full flex items-center justify-center text-opacity-80">Загрузка карты...</div>
          </template>
        </ClientOnly>
      </div>

      <!-- Секция с мероприятиями -->
      <div class="grid grid-cols-1 gap-8">
        <div>
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-2xl font-bold">Все мероприятия</h2>
            <div class="flex items-center gap-2">
              <button @click="toggleViewMode" class="p-2 rounded-lg hover:bg-opacity-20 bg-opacity-10 transition-colors">
                <svg v-if="viewMode === 'grid'" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
                </svg>
              </button>
              <button @click="toggleTheme" class="p-2 rounded-lg hover:bg-opacity-20 bg-opacity-10 transition-colors">
                <svg v-if="isDarkMode" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
                </svg>
              </button>
              <button class="action-btn p-2 px-4 rounded-lg flex items-center gap-2">
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                </svg>
                Создать мероприятие
              </button>
            </div>
          </div>
          
          <!-- Компонент EventsGrid для отображения событий -->
          <EventsGrid :events="events" @select-event="showEventDetails" />
        </div>
      </div>

      <!-- Рекомендации и активность -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="col-span-2 space-y-4 events-section p-4 rounded-2xl shadow-lg">
          <h2 class="text-xl font-semibold text-accent">Рекомендуемые события</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div
              v-for="event in popularEvents"
              :key="event.id"
              :id="`event-${event.id}`"
              class="p-4 rounded-xl event-card shadow-md cursor-pointer transition-all duration-300 hover:shadow-lg"
              @click="showEventDetails(event)"
            >
              <h3 class="text-lg font-semibold">{{ event.title }}</h3>
              <p class="text-sm opacity-80 mt-1">{{ formatDate(event.date) || 'Дата неизвестна' }}</p>
              <div class="flex mt-2">
                <span v-for="tag in event.tags" :key="tag" 
                      :class="['mr-2 px-2 py-1 rounded-full text-xs', 
                              isDarkMode ? 'bg-purple bg-opacity-50' : 'bg-purple bg-opacity-30']">
                  {{ tag }}
                </span>
              </div>
            </div>
          </div>
        </div>

        <div class="col-span-1 space-y-4">
          <div class="p-4 rounded-2xl shadow-lg events-section">
            <h3 class="text-lg font-semibold text-accent">Популярное</h3>
            <ul class="space-y-2">
              <li v-for="(event, index) in popularEvents" :key="index" class="text-sm hover:text-accent cursor-pointer transition-colors">
                {{ event.title }}
              </li>
            </ul>
          </div>
          <div class="p-4 rounded-2xl shadow-lg events-section">
            <h3 class="text-lg font-semibold text-accent">Быстрые действия</h3>
            <div class="space-y-2">
              <button class="w-full action-btn p-2 rounded-xl flex items-center justify-center gap-2">
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
                </svg>
                Добавить событие
              </button>
              <button class="w-full cancel-btn p-2 rounded-xl flex items-center justify-center gap-2">
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v8m4-8v8m4-8v8" />
                </svg>
                Поделиться картой
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Детали события -->
      <EventDetails
        v-if="selectedEvent"
        :event="selectedEvent"
        :is-dark-mode="isDarkMode"
        @close="closeEventDetails"
        @toggleFavorite="handleToggleFavorite"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import Navbar from '~/components/Navbar.vue'
import Map from '~/components/Map.vue'
import EventsGrid from '~/components/EventsGrid.vue'
import EventDetails from '~/components/EventDetails.vue'

// Состояния
const theme = useState('theme')
const isDarkMode = computed(() => theme.value === 'dark')
const events = ref([])
const selectedEventId = ref(null)
const selectedEvent = ref(null)
const isClient = ref(false)
const showSuccess = ref(false)
const showRemoveSuccess = ref(false)
const viewMode = ref('grid') // grid или list

// Инициализация
onMounted(() => {
  isClient.value = true
  generateEvents()
  loadFavorites()
})

// Генерация тестовых данных
function generateEvents() {
  const categories = ['tech', 'music', 'art', 'business', 'education', 'gaming', 'science', 'crypto']
  const tags = [
    ['web', 'frontend', 'backend', 'fullstack', 'devops'],
    ['blockchain', 'crypto', 'nft', 'web3', 'defi'],
    ['ai', 'machine learning', 'neural networks', 'data science', 'robotics'],
    ['design', 'ui/ux', 'product', 'графический дизайн', 'motion'],
    ['mobile', 'ios', 'android', 'react native', 'flutter'],
    ['business', 'startup', 'инвестиции', 'предпринимательство', 'маркетинг'],
    ['cybersecurity', 'облачные технологии', 'big data', 'iot', 'архитектура'],
    ['education', 'курсы', 'карьера', 'менторство', 'soft skills'],
    ['музыка', 'живые выступления', 'электронная', 'классическая', 'джаз'],
    ['gaming', 'киберспорт', 'game dev', 'vr/ar', 'метавселенная']
  ]
  
  const locations = [
    { city: 'Москва', venue: 'Цифровой Деловой Центр' },
    { city: 'Санкт-Петербург', venue: 'Новая Голландия' },
    { city: 'Казань', venue: 'Смарт-Сити Хаб' },
    { city: 'Екатеринбург', venue: 'Конгресс-центр "Высоцкий"' },
    { city: 'Новосибирск', venue: 'Технопарк Академгородка' },
    { city: 'Сочи', venue: 'Олимпийский медиа-центр' },
    { city: 'Калининград', venue: 'Куб' },
    { city: 'Владивосток', venue: 'Океанариум ДВФУ' },
    { city: 'Нижний Новгород', venue: 'Технопарк "Анкудиновка"' },
    { city: 'Минск', venue: 'ВЕБ парк' },
    { city: 'Берлин', venue: 'Station Berlin' },
    { city: 'Амстердам', venue: 'The Edge' },
    { city: 'Дубай', venue: 'Museum of the Future' },
    { city: 'Сингапур', venue: 'Marina Bay Sands' },
    { city: 'Сан-Франциско', venue: 'Moscone Center' }
  ]
  
  const titles = [
    'DevConf 2023',
    'AI Summit',
    'Blockchain Forum',
    'UX/UI Masterclass',
    'Web3 Conference',
    'Digital Marketing Forum',
    'Data Science Day',
    'JavaScript Meetup',
    'Startup Pitch Night',
    'Tech Leadership Summit',
    'Crypto Expo',
    'Product Management Workshop',
    'VR/AR Exhibition',
    'Mobile Dev Conference',
    'Cybersecurity Symposium',
    'Cloud Computing Summit',
    'Game Developers Meet',
    'NFT Art Festival',
    'IoT Conference',
    'Music Tech Fusion'
  ]
  
  const descriptions = [
    'Глобальная конференция, собирающая лидеров индустрии для обсуждения последних трендов в разработке ПО и инновационных решений.',
    'Погружение в мир искусственного интеллекта с практическими демонстрациями, кейс-стади и выступлениями от ведущих экспертов области.',
    'Интенсивный воркшоп по современным технологиям блокчейн с фокусом на практическое применение в бизнесе и финансах.',
    'Уникальное мероприятие для дизайнеров всех уровней, сочетающее теоретические знания с интерактивными практическими сессиями.',
    'Форум для обсуждения будущего web3-технологий, децентрализованных приложений и метавселенных в контексте меняющегося цифрового ландшафта.',
    'Комплексная образовательная программа, охватывающая все аспекты современного цифрового маркетинга от аналитики до многоканального продвижения.',
    'Специализированное мероприятие для data scientists и аналитиков с глубоким погружением в методы анализа больших данных и машинное обучение.',
    'Ежемесячная встреча энтузиастов JavaScript для обмена опытом, совместного решения сложных задач и нетворкинга в неформальной обстановке.',
    'Вечер для стартаперов и инвесторов, где инновационные проекты презентуют свои идеи и бизнес-модели потенциальным партнерам и менторам.',
    'Серия дискуссионных панелей и мастер-классов для технических руководителей всех уровней по развитию команд и управлению технологическими процессами.'
  ]

  // Speakers database
  const speakersDatabase = [
    { id: 1, name: 'Александр Волков', position: 'AI Research Lead, DeepMind', avatar: 'https://randomuser.me/api/portraits/men/32.jpg' },
    { id: 2, name: 'Мария Соколова', position: 'Product Director, TechVision', avatar: 'https://randomuser.me/api/portraits/women/44.jpg' },
    { id: 3, name: 'Илья Морозов', position: 'Blockchain Developer, CryptoLab', avatar: 'https://randomuser.me/api/portraits/men/68.jpg' },
    { id: 4, name: 'Екатерина Ковалева', position: 'UX/UI Lead, DesignForce', avatar: 'https://randomuser.me/api/portraits/women/65.jpg' },
    { id: 5, name: 'Сергей Игнатов', position: 'CTO, InnovaTech', avatar: 'https://randomuser.me/api/portraits/men/41.jpg' },
    { id: 6, name: 'Анастасия Орлова', position: 'Frontend Architect, WebPioneer', avatar: 'https://randomuser.me/api/portraits/women/33.jpg' },
    { id: 7, name: 'Павел Дмитриев', position: 'DevOps Lead, CloudMasters', avatar: 'https://randomuser.me/api/portraits/men/22.jpg' },
    { id: 8, name: 'Ольга Зайцева', position: 'Mobile Dev Expert, AppCraft', avatar: 'https://randomuser.me/api/portraits/women/28.jpg' },
    { id: 9, name: 'Артём Соловьёв', position: 'ML Engineer, AILabs', avatar: 'https://randomuser.me/api/portraits/men/18.jpg' },
    { id: 10, name: 'Ирина Белова', position: 'Security Specialist, SecureNet', avatar: 'https://randomuser.me/api/portraits/women/17.jpg' },
    { id: 11, name: 'Владимир Козлов', position: 'Backend Developer, DataSystems', avatar: 'https://randomuser.me/api/portraits/men/55.jpg' },
    { id: 12, name: 'Наталья Романова', position: 'Data Scientist, AnalyticsPro', avatar: 'https://randomuser.me/api/portraits/women/8.jpg' }
  ]

  // Celebrities database
  const celebritiesDatabase = [
    { id: 1, name: 'Тимур Бекмамбетов', type: 'режиссер', avatar: 'https://randomuser.me/api/portraits/men/71.jpg' },
    { id: 2, name: 'Ксения Собчак', type: 'телеведущая', avatar: 'https://randomuser.me/api/portraits/women/79.jpg' },
    { id: 3, name: 'Илон Маск', type: 'предприниматель', avatar: 'https://randomuser.me/api/portraits/men/44.jpg' },
    { id: 4, name: 'Настя Ивлеева', type: 'блогер', avatar: 'https://randomuser.me/api/portraits/women/91.jpg' },
    { id: 5, name: 'Павел Дуров', type: 'IT-предприниматель', avatar: 'https://randomuser.me/api/portraits/men/9.jpg' },
    { id: 6, name: 'Юрий Дудь', type: 'журналист', avatar: 'https://randomuser.me/api/portraits/men/36.jpg' },
    { id: 7, name: 'Ольга Бузова', type: 'певица', avatar: 'https://randomuser.me/api/portraits/women/62.jpg' },
    { id: 8, name: 'Хабиб Нурмагомедов', type: 'спортсмен', avatar: 'https://randomuser.me/api/portraits/men/81.jpg' }
  ]

  // Organizers database
  const organizersDatabase = [
    'Event SpheRRe',
    'DigitalHorizon',
    'TechConnect',
    'NextGen Events',
    'FutureVision',
    'InnovateNow',
    'CreativeMinds',
    'ArtSpace Gallery',
    'CyberSummit',
    'Web3 Community',
    'AI Consortium',
    'MusicWave',
    'SciTech Hub'
  ]

  const newEvents = []
  for (let i = 0; i < 25; i++) {
    const randomDate = new Date()
    randomDate.setDate(randomDate.getDate() + Math.floor(Math.random() * 30))
    randomDate.setHours(Math.floor(Math.random() * 12) + 9, Math.floor(Math.random() * 60))

    const lat = 41.1850968 + Math.random() * (82.0586232 - 41.1850968)
    const lng = 19.6389 + Math.random() * (180 - 19.6389)
    const category = categories[Math.floor(Math.random() * categories.length)]
    const eventTags = tags[Math.floor(Math.random() * tags.length)]
    const location = locations[Math.floor(Math.random() * locations.length)]
    const title = titles[Math.floor(Math.random() * titles.length)]
    const description = descriptions[Math.floor(Math.random() * descriptions.length)]
    const organizer = organizersDatabase[Math.floor(Math.random() * organizersDatabase.length)]

    // Randomly assign speakers (between 2-5)
    const speakers = []
    const speakerCount = Math.floor(Math.random() * 4) + 2
    const shuffledSpeakers = [...speakersDatabase].sort(() => 0.5 - Math.random())
    for (let j = 0; j < speakerCount; j++) {
      if (shuffledSpeakers[j]) {
        speakers.push(shuffledSpeakers[j])
      }
    }

    // Randomly decide if event has celebrities (45% chance)
    let celebrities = []
    if (Math.random() < 0.45) {
      const celebrityCount = Math.floor(Math.random() * 3) + 1
      const shuffledCelebrities = [...celebritiesDatabase].sort(() => 0.5 - Math.random())
      for (let j = 0; j < celebrityCount; j++) {
        if (shuffledCelebrities[j]) {
          celebrities.push(shuffledCelebrities[j])
        }
      }
    }

    newEvents.push({
      id: i + 1,
      title: `${title} ${i + 1}`,
      date: randomDate.toISOString(),
      lat,
      lng,
      category,
      tags: eventTags,
      location,
      description,
      organizer,
      price: Math.floor(Math.random() * 500) * 10,
      speakers,
      celebrities
    })
  }
  
  events.value = newEvents
}

// Загрузка избранного
function loadFavorites() {
  if (process.client) {
    try {
      const saved = localStorage.getItem('favorites')
      if (saved) {
        favorites.value = JSON.parse(saved)
      }
    } catch (e) {
      console.error('Ошибка загрузки избранного:', e)
    }
  }
}

// Популярные события
const popularEvents = computed(() => events.value.slice(0, 4))

// Обновление выбранного события
const onUpdateSelected = (id) => {
  selectedEventId.value = id
  if (id) {
    const event = events.value.find(e => e.id === id)
    if (event) {
      selectedEvent.value = event
    }
  }
}

// Генерация анимации "падающих звезд"
const randomFallingStarStyle = () => {
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

// Переключение темы
const toggleTheme = () => {
  theme.value = isDarkMode.value ? 'light' : 'dark'
}

// Переключение режима отображения
const toggleViewMode = () => {
  viewMode.value = viewMode.value === 'grid' ? 'list' : 'grid'
}

// Отображение деталей события
const showEventDetails = (event) => {
  selectedEvent.value = event
}

// Закрытие деталей события
const closeEventDetails = () => {
  selectedEvent.value = null
}

// Обработка добавления в избранное
const handleToggleFavorite = ({ favorites, isAdded }) => {
  if (isAdded) {
    showSuccess.value = true
    setTimeout(() => showSuccess.value = false, 3000)
  } else {
    showRemoveSuccess.value = true
    setTimeout(() => showRemoveSuccess.value = false, 3000)
  }
}

// Форматирование даты
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

<style scoped>
/* Светлая тема */
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

/* Темная тема */
.dark-theme {
  background-color: #1A1A2E;
  color: #E4E4F0;
}

.dark-theme .events-section {
  background-color: #2A2A40;
}

.dark-theme .event-card {
  background-color: #3A3A50;
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
.input-field {
  padding: 12px 16px;
  border-radius: 12px;
  width: 100%;
  transition: all 0.3s ease;
}

.input-field:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(108, 99, 255, 0.3);
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
  background-color: var(--purple);
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