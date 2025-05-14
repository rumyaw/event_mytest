<template>
  <div class="fixed right-4 top-20 z-50 transition-all duration-300" :class="[isOpen ? 'translate-x-0' : 'translate-x-[105%]']">
    <div :class="['p-4 rounded-lg shadow-xl w-[450px] max-h-[85vh] overflow-y-auto', 
                 isDarkMode ? 'dark-theme-bg' : 'light-theme-bg']">
      
      <!-- Header -->
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Рекомендации</h3>
        <div class="flex gap-2">
          <button @click="showPreferences = !showPreferences" class="p-1 rounded-full hover:bg-opacity-10 hover:bg-black">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            </svg>
          </button>
          <button @click="closePanel" class="p-1 rounded-full hover:bg-opacity-10 hover:bg-black">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
      
      <!-- Main content -->
      <div v-if="!showPreferences">
        <!-- Tabs -->
        <div class="flex mb-4 border-b" :class="isDarkMode ? 'border-gray-700' : 'border-gray-300'">
          <button 
            @click="activeTab = 'personal'" 
            class="px-4 py-2 text-sm font-medium transition-colors -mb-px"
            :class="activeTab === 'personal' ? 
              (isDarkMode ? 'text-accent border-b-2 border-accent' : 'text-accent border-b-2 border-accent') : 
              (isDarkMode ? 'text-gray-400 hover:text-gray-200' : 'text-gray-500 hover:text-gray-800')"
          >
            Для вас
          </button>
          <button 
            @click="activeTab = 'nearby'" 
            class="px-4 py-2 text-sm font-medium transition-colors -mb-px"
            :class="activeTab === 'nearby' ? 
              (isDarkMode ? 'text-accent border-b-2 border-accent' : 'text-accent border-b-2 border-accent') : 
              (isDarkMode ? 'text-gray-400 hover:text-gray-200' : 'text-gray-500 hover:text-gray-800')"
          >
            Рядом
          </button>
          <button 
            @click="activeTab = 'popular'" 
            class="px-4 py-2 text-sm font-medium transition-colors -mb-px"
            :class="activeTab === 'popular' ? 
              (isDarkMode ? 'text-accent border-b-2 border-accent' : 'text-accent border-b-2 border-accent') : 
              (isDarkMode ? 'text-gray-400 hover:text-gray-200' : 'text-gray-500 hover:text-gray-800')"
          >
            Популярные
          </button>
        </div>
        
        <!-- Recommendations list -->
        <div class="space-y-4">
          <TransitionGroup name="recommendation-list">
            <div 
              v-for="recommendation in filteredRecommendations" 
              :key="recommendation.id"
              :class="['p-4 rounded-lg relative transition-all cursor-pointer', 
                     isDarkMode ? 'bg-gray-800 hover:bg-gray-700' : 'bg-gray-100 hover:bg-gray-200']"
              @click="viewEventDetails(recommendation.event)"
            >
              <!-- Recommendation badge -->
              <div 
                class="absolute px-2 py-1 rounded-full text-xs top-4 right-4 font-medium shadow-sm"
                :class="getRecommendationTypeClass(recommendation.type)"
              >
                {{ recommendationTypeLabels[recommendation.type] }}
              </div>
              
              <!-- Recommendation content -->
              <div class="flex items-start gap-4">
                <div class="flex-shrink-0">
                  <div :class="['p-2 rounded-full', recommendation.type === 'location' ? 'bg-blue-600 text-white' : getRecommendationTypeClass(recommendation.type)]">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" viewBox="0 0 20 20" fill="currentColor">
                      <path v-if="recommendation.type === 'interest'" fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd" />
                      <path v-else-if="recommendation.type === 'location'" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" />
                      <path v-else-if="recommendation.type === 'history'" d="M9 4.804A7.968 7.968 0 005.5 4c-1.255 0-2.443.29-3.5.804v10A7.969 7.969 0 015.5 14c1.669 0 3.218.51 4.5 1.385A7.962 7.962 0 0114.5 14c1.255 0 2.443.29 3.5.804v-10A7.968 7.968 0 0014.5 4c-1.255 0-2.443.29-3.5.804V12a1 1 0 11-2 0V4.804z" />
                      <path v-else d="M5 3a2 2 0 012-2h10a2 2 0 012 2v16a2 2 0 01-2 2H7a2 2 0 01-2-2V3zm2 0v16h10V3H7zm5 5a1 1 0 100-2 1 1 0 000 2zm-1 5a1 1 0 112 0v4a1 1 0 11-2 0v-4z" />
                    </svg>
                  </div>
                </div>
                <div class="flex-1 pr-6">
                  <div class="font-medium text-base mb-1">{{ recommendation.event.title }}</div>
                  <p class="text-sm mt-1">{{ formatDate(recommendation.event.date) }}</p>
                  <p class="mt-2">{{ recommendation.reason }}</p>
                  
                  <!-- Tags -->
                  <div class="flex flex-wrap gap-2 mt-3">
                    <span v-for="tag in recommendation.event.tags" :key="tag" 
                          :class="['px-2 py-1 rounded-full text-xs font-medium', 
                                  isDarkMode ? 'bg-indigo-700 text-white' : 'bg-indigo-100 text-indigo-800']">
                      {{ tag }}
                    </span>
                  </div>
                  
                  <!-- Action buttons -->
                  <div class="mt-3 flex gap-3">
                    <button 
                      @click.stop="addToFavorites(recommendation.event)" 
                      class="px-4 py-1.5 text-xs rounded-lg bg-accent text-white hover:bg-opacity-90 transition-colors"
                    >
                      В избранное
                    </button>
                    <button 
                      @click.stop="dismissRecommendation(recommendation.id)" 
                      class="px-4 py-1.5 text-xs rounded-lg bg-gray-500 bg-opacity-90 text-white hover:bg-opacity-100 transition-colors"
                    >
                      Скрыть
                    </button>
                  </div>
                </div>
              </div>
            </div>
            
            <div v-if="filteredRecommendations.length === 0" key="empty" class="text-center py-12 opacity-60">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16 mx-auto mb-4 opacity-40" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
              <p class="text-lg">Нет рекомендаций</p>
            </div>
          </TransitionGroup>
        </div>
      </div>
      
      <!-- Preferences panel -->
      <div v-else>
        <div class="mb-4 pb-2 border-b flex justify-between items-center" 
             :class="isDarkMode ? 'border-gray-700' : 'border-gray-300'">
          <h4 class="font-medium">Настройки рекомендаций</h4>
        </div>
        
        <!-- User interests -->
        <div class="mb-5">
          <h5 class="font-medium mb-3">Ваши интересы</h5>
          <div class="flex flex-wrap gap-2 mb-4">
            <div 
              v-for="interest in interests" 
              :key="interest.id"
              @click="toggleInterest(interest.id)"
              :class="['px-3 py-1.5 rounded-full text-sm cursor-pointer transition-colors', 
                      interest.selected ? 
                        'bg-accent text-white' : 
                        (isDarkMode ? 'bg-gray-700 text-gray-300' : 'bg-gray-200 text-gray-700')]"
            >
              {{ interest.name }}
            </div>
            <div 
              class="px-3 py-1.5 rounded-full text-sm cursor-pointer transition-colors"
              :class="isDarkMode ? 'bg-gray-700 text-gray-300' : 'bg-gray-200 text-gray-700'"
            >
              <span>+</span>
            </div>
          </div>
        </div>
        
        <!-- Recommendation settings -->
        <div class="space-y-4">
          <div :class="['p-4 rounded-lg', isDarkMode ? 'bg-gray-800' : 'bg-gray-100']">
            <h5 class="font-medium mb-3">Факторы рекомендаций</h5>
            
            <div class="space-y-3">
              <div class="flex items-center justify-between">
                <label class="flex items-center text-sm">
                  <div class="relative flex items-center">
                    <input type="checkbox" v-model="recommendationSettings.history" class="sr-only peer" />
                    <div class="w-5 h-5 bg-gray-600 rounded-sm peer-checked:bg-accent peer-checked:flex peer-checked:items-center peer-checked:justify-center mr-2">
                      <svg v-if="recommendationSettings.history" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-white" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                    </div>
                  </div>
                  История просмотров
                </label>
                <span class="text-xs opacity-60">Высокий приоритет</span>
              </div>
              
              <div class="flex items-center justify-between">
                <label class="flex items-center text-sm">
                  <div class="relative flex items-center">
                    <input type="checkbox" v-model="recommendationSettings.favorites" class="sr-only peer" />
                    <div class="w-5 h-5 bg-gray-600 rounded-sm peer-checked:bg-accent peer-checked:flex peer-checked:items-center peer-checked:justify-center mr-2">
                      <svg v-if="recommendationSettings.favorites" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-white" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                    </div>
                  </div>
                  Избранные мероприятия
                </label>
                <span class="text-xs opacity-60">Высокий приоритет</span>
              </div>
              
              <div class="flex items-center justify-between">
                <label class="flex items-center text-sm">
                  <div class="relative flex items-center">
                    <input type="checkbox" v-model="recommendationSettings.interests" class="sr-only peer" />
                    <div class="w-5 h-5 bg-gray-600 rounded-sm peer-checked:bg-accent peer-checked:flex peer-checked:items-center peer-checked:justify-center mr-2">
                      <svg v-if="recommendationSettings.interests" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-white" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                    </div>
                  </div>
                  Интересы
                </label>
                <span class="text-xs opacity-60">Средний приоритет</span>
              </div>
              
              <div class="flex items-center justify-between">
                <label class="flex items-center text-sm">
                  <div class="relative flex items-center">
                    <input type="checkbox" v-model="recommendationSettings.location" class="sr-only peer" />
                    <div class="w-5 h-5 bg-gray-600 rounded-sm peer-checked:bg-accent peer-checked:flex peer-checked:items-center peer-checked:justify-center mr-2">
                      <svg v-if="recommendationSettings.location" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-white" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                    </div>
                  </div>
                  Геолокация
                </label>
                <span class="text-xs opacity-60">Низкий приоритет</span>
              </div>
              
              <div class="flex items-center justify-between">
                <label class="flex items-center text-sm">
                  <div class="relative flex items-center">
                    <input type="checkbox" v-model="recommendationSettings.rating" class="sr-only peer" />
                    <div class="w-5 h-5 bg-gray-600 rounded-sm peer-checked:bg-accent peer-checked:flex peer-checked:items-center peer-checked:justify-center mr-2">
                      <svg v-if="recommendationSettings.rating" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-white" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                      </svg>
                    </div>
                  </div>
                  Рейтинг мероприятий
                </label>
                <span class="text-xs opacity-60">Низкий приоритет</span>
              </div>
            </div>
          </div>
          
          <div :class="['p-4 rounded-lg', isDarkMode ? 'bg-gray-800' : 'bg-gray-100']">
            <h5 class="font-medium mb-3">Ваше местоположение</h5>
            
            <div class="flex items-center justify-between mb-3">
              <label class="flex items-center text-sm">
                <div class="relative flex items-center">
                  <input type="checkbox" v-model="useCurrentLocation" class="sr-only peer" />
                  <div class="w-5 h-5 bg-gray-600 rounded-sm peer-checked:bg-accent peer-checked:flex peer-checked:items-center peer-checked:justify-center mr-2">
                    <svg v-if="useCurrentLocation" xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-white" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                  </div>
                </div>
                Использовать текущее местоположение
              </label>
            </div>
            
            <div v-if="!useCurrentLocation" class="space-y-2">
              <select class="w-full p-2 rounded-lg text-sm"
                     :class="isDarkMode ? 'bg-gray-700 text-gray-200' : 'bg-gray-200 text-gray-700'">
                <option>Москва</option>
                <option>Санкт-Петербург</option>
                <option>Казань</option>
                <option>Новосибирск</option>
                <option>Екатеринбург</option>
              </select>
            </div>
          </div>
          
          <button @click="savePreferences" class="w-full p-3 rounded-lg bg-accent text-white font-medium">
            Сохранить настройки
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';

// State
const props = defineProps({
  show: {
    type: Boolean,
    default: false
  }
});

const emits = defineEmits(['close', 'viewEvent', 'addToFavorites']);

const isOpen = computed(() => props.show);
const activeTab = ref('personal');
const recommendations = ref([]);
const showPreferences = ref(false);
const interests = ref([]);
const useCurrentLocation = ref(true);
const recommendationSettings = ref({
  history: true,
  favorites: true,
  interests: true,
  location: true,
  rating: true
});

// Theme
const theme = useState('theme', () => 'dark');
const isDarkMode = computed(() => theme.value === 'dark');

// Computed
const filteredRecommendations = computed(() => {
  switch (activeTab.value) {
    case 'personal':
      return recommendations.value.filter(rec => rec.type === 'interest' || rec.type === 'history');
    case 'nearby':
      return recommendations.value.filter(rec => rec.type === 'location');
    case 'popular':
      return recommendations.value.filter(rec => rec.type === 'rating');
    default:
      return recommendations.value;
  }
});

const recommendationsCount = computed(() => {
  return recommendations.value.length;
});

const recommendationTypeLabels = {
  'interest': 'По интересам',
  'history': 'История просмотров',
  'location': 'Рядом с вами',
  'rating': 'Популярное'
};

// Methods
function togglePanel() {
  showPreferences.value = false;
}

function closePanel() {
  showPreferences.value = false;
  emits('close');
}

function dismissRecommendation(id) {
  recommendations.value = recommendations.value.filter(recommendation => recommendation.id !== id);
}

function viewEventDetails(event) {
  // Mark as viewed in history
  if (!event.viewedAt) {
    event.viewedAt = new Date().toISOString();
    generateRecommendations(); // Regenerate recommendations based on new view
  }
  
  // Emit event
  emits('viewEvent', event);
  closePanel();
}

function getRecommendationTypeClass(type) {
  const classes = {
    'interest': 'bg-purple-600 bg-opacity-100 text-white',
    'history': 'bg-green-600 bg-opacity-100 text-white',
    'location': 'bg-blue-600 bg-opacity-100 text-white',
    'rating': 'bg-yellow-600 bg-opacity-100 text-white'
  };
  
  return classes[type] || classes.interest;
}

function formatDate(dateString) {
  const date = new Date(dateString);
  const now = new Date();
  const diffMs = date - now; // Future events have positive diffMs
  
  if (diffMs <= 0) {
    return 'Уже началось';
  }
  
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));
  
  if (diffDays === 0) {
    return 'Сегодня';
  } else if (diffDays === 1) {
    return 'Завтра';
  } else if (diffDays < 7) {
    return `Через ${diffDays} дн.`;
  } else {
    return date.toLocaleDateString('ru-RU', {
      day: 'numeric',
      month: 'short'
    });
  }
}

function toggleInterest(interestId) {
  const index = interests.value.findIndex(interest => interest.id === interestId);
  if (index !== -1) {
    interests.value[index].selected = !interests.value[index].selected;
  }
}

function savePreferences() {
  showPreferences.value = false;
  generateRecommendations();
}

function addToFavorites(event) {
  event.isFavorite = true;
  emits('addToFavorites', event);
  generateRecommendations(); // Regenerate based on new favorite
}

// Init
onMounted(() => {
  // Mock data
  mockData();
  generateRecommendations();
});

function mockData() {
  // Sample interests
  interests.value = [
    { id: 1, name: 'Музыка', selected: true },
    { id: 2, name: 'Технологии', selected: true },
    { id: 3, name: 'Искусство', selected: false },
    { id: 4, name: 'Спорт', selected: false },
    { id: 5, name: 'Образование', selected: true },
    { id: 6, name: 'Бизнес', selected: false },
    { id: 7, name: 'Еда и напитки', selected: true }
  ];
}

function generateRecommendations() {
  // Sample events
  const events = [
    {
      id: 1,
      title: 'Концерт классической музыки',
      date: new Date(Date.now() + 1000 * 60 * 60 * 24 * 2).toISOString(), // 2 days ahead
      description: 'Насладитесь произведениями великих композиторов в исполнении симфонического оркестра.',
      category: 'music',
      tags: ['классика', 'симфония', 'оркестр'],
      location: { lat: 55.75, lng: 37.62, name: 'Консерватория им. П.И. Чайковского' },
      rating: 4.8,
      viewedAt: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString() // Viewed yesterday
    },
    {
      id: 2,
      title: 'DevConf 2025',
      date: new Date(Date.now() + 1000 * 60 * 60 * 24 * 5).toISOString(), // 5 days ahead
      description: 'Крупнейшая технологическая конференция года с участием ведущих экспертов в области программирования.',
      category: 'tech',
      tags: ['конференция', 'IT', 'технологии'],
      location: { lat: 55.78, lng: 37.64, name: 'Экспоцентр' },
      rating: 4.6,
      isFavorite: true
    },
    {
      id: 3,
      title: 'Выставка современного искусства',
      date: new Date(Date.now() + 1000 * 60 * 60 * 24 * 1).toISOString(), // Tomorrow
      description: 'Экспозиция работ молодых художников, представляющих различные направления современного искусства.',
      category: 'art',
      tags: ['выставка', 'искусство', 'галерея'],
      location: { lat: 55.74, lng: 37.61, name: 'Центр современного искусства' },
      rating: 4.2
    },
    {
      id: 4,
      title: 'Мастер-класс по UI/UX дизайну',
      date: new Date(Date.now() + 1000 * 60 * 60 * 24 * 3).toISOString(), // 3 days ahead
      description: 'Интенсивный мастер-класс для начинающих и практикующих дизайнеров интерфейсов.',
      category: 'tech',
      tags: ['дизайн', 'UI/UX', 'мастер-класс'],
      location: { lat: 55.76, lng: 37.63, name: 'Коворкинг "Рабочая станция"' },
      rating: 4.5
    },
    {
      id: 5,
      title: 'Фестиваль уличной еды',
      date: new Date(Date.now() + 1000 * 60 * 60 * 24 * 6).toISOString(), // 6 days ahead
      description: 'Гастрономический фестиваль с участием лучших фуд-траков и ресторанов города.',
      category: 'food',
      tags: ['еда', 'фестиваль', 'стритфуд'],
      location: { lat: 55.73, lng: 37.60, name: 'Парк Горького' },
      rating: 4.7
    }
  ];
  
  // Generate recommendations based on user preferences
  recommendations.value = [];
  
  // Recommendations based on interests
  if (recommendationSettings.value.interests) {
    const userInterests = interests.value.filter(interest => interest.selected).map(i => i.name.toLowerCase());
    
    events.forEach(event => {
      // Check if event matches user interests
      const hasMatchingInterest = event.tags.some(tag => 
        userInterests.some(interest => tag.toLowerCase().includes(interest.toLowerCase()) || interest.includes(tag.toLowerCase()))
      );
      
      if (hasMatchingInterest) {
        let matchingInterest = '';
        for (const tag of event.tags) {
          for (const interest of userInterests) {
            if (tag.toLowerCase().includes(interest.toLowerCase()) || interest.includes(tag.toLowerCase())) {
              matchingInterest = interest;
              break;
            }
          }
          if (matchingInterest) break;
        }
        
        recommendations.value.push({
          id: `interest-${event.id}`,
          type: 'interest',
          event: event,
          reason: `Соответствует вашему интересу: ${matchingInterest}`
        });
      }
    });
  }
  
  // Recommendations based on view history
  if (recommendationSettings.value.history) {
    const viewedEvents = events.filter(event => event.viewedAt);
    
    if (viewedEvents.length > 0) {
      // Get categories and tags from viewed events
      const viewedCategories = new Set(viewedEvents.map(e => e.category));
      const viewedTags = viewedEvents.flatMap(e => e.tags);
      
      // Find similar events
      events.forEach(event => {
        // Skip if already recommended or if it's a viewed event
        if (
          recommendations.value.some(r => r.event.id === event.id) || 
          viewedEvents.some(e => e.id === event.id)
        ) {
          return;
        }
        
        // Check if event is similar to what user viewed
        const hasSameCategory = viewedCategories.has(event.category);
        const hasSimilarTags = event.tags.some(tag => viewedTags.includes(tag));
        
        if (hasSameCategory || hasSimilarTags) {
          recommendations.value.push({
            id: `history-${event.id}`,
            type: 'history',
            event: event,
            reason: `Похоже на мероприятия, которые вы просматривали`
          });
        }
      });
    }
  }
  
  // Recommendations based on location
  if (recommendationSettings.value.location && useCurrentLocation.value) {
    // Pretend current location is Moscow center
    const userLocation = { lat: 55.75, lng: 37.62 };
    
    // Simple distance calculation (not accurate, just for demo)
    function calculateDistance(loc1, loc2) {
      return Math.sqrt(Math.pow(loc1.lat - loc2.lat, 2) + Math.pow(loc1.lng - loc2.lng, 2));
    }
    
    events.forEach(event => {
      // Skip if already recommended
      if (recommendations.value.some(r => r.event.id === event.id)) {
        return;
      }
      
      // Check if event is nearby
      if (event.location) {
        const distance = calculateDistance(userLocation, event.location);
        if (distance < 0.1) { // Arbitrary threshold
          recommendations.value.push({
            id: `location-${event.id}`,
            type: 'location',
            event: event,
            reason: `Находится недалеко от вас (${event.location.name})`
          });
        }
      }
    });
  }
  
  // Recommendations based on rating
  if (recommendationSettings.value.rating) {
    // Get high-rated events
    const highRatedEvents = events
      .filter(event => event.rating >= 4.5)
      .sort((a, b) => b.rating - a.rating);
    
    highRatedEvents.forEach(event => {
      // Skip if already recommended
      if (recommendations.value.some(r => r.event.id === event.id)) {
        return;
      }
      
      recommendations.value.push({
        id: `rating-${event.id}`,
        type: 'rating',
        event: event,
        reason: `Высокий рейтинг (${event.rating}/5)`
      });
    });
  }
  
  // Recommendations based on favorites
  if (recommendationSettings.value.favorites) {
    const favoriteEvents = events.filter(event => event.isFavorite);
    
    if (favoriteEvents.length > 0) {
      // Find similar events to favorites
      events.forEach(event => {
        // Skip if already recommended or if it's a favorite
        if (
          recommendations.value.some(r => r.event.id === event.id) || 
          favoriteEvents.some(e => e.id === event.id)
        ) {
          return;
        }
        
        // Check similarity
        for (const favorite of favoriteEvents) {
          const hasSameCategory = event.category === favorite.category;
          const hasSimilarTags = event.tags.some(tag => favorite.tags.includes(tag));
          
          if (hasSameCategory || hasSimilarTags) {
            recommendations.value.push({
              id: `favorite-${event.id}`,
              type: 'interest', // Group with interests
              event: event,
              reason: `Похоже на мероприятия в вашем избранном`
            });
            break;
          }
        }
      });
    }
  }
}
</script>

<style scoped>
.dark-theme-bg {
  background-color: #2A2A40;
  color: #E4E4F0;
}

.light-theme-bg {
  background-color: white;
  color: #1A1A2E;
}

/* Recommendation list transitions */
.recommendation-list-enter-active,
.recommendation-list-leave-active {
  transition: all 0.3s ease;
}

.recommendation-list-enter-from,
.recommendation-list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.recommendation-list-move {
  transition: transform 0.3s ease;
}

.text-accent {
  color: #6C63FF;
}

.bg-accent {
  background-color: #6C63FF;
}

.bg-purple-500 {
  background-color: #8B5CF6;
}

.bg-purple-400 {
  background-color: #A78BFA;
}

.accent-purple {
  accent-color: #8B5CF6;
}
</style> 