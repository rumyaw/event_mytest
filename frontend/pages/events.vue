<template>
  <div :class="['min-h-screen', isDarkMode ? 'dark-theme' : 'light-theme']">
    <Navbar />
    
    <div class="container mx-auto p-6 pt-24">
      <div class="flex flex-col md:flex-row justify-between items-center mb-8">
        <h1 class="text-3xl font-bold">Мероприятия</h1>
        
        <div class="mt-4 md:mt-0 flex gap-3">
          <button class="action-btn p-2 rounded-xl flex items-center justify-center gap-2">
            <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
            Создать мероприятие
          </button>
          <button @click="toggleTheme" class="p-2 rounded-xl shadow-sm border">
            <svg v-if="isDarkMode" class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364 6.364l-.707-.707M6.343 6.343l-.707-.707m12.728 0l-.707.707M6.343 17.657l-.707.707M16 12a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            <svg v-else class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
            </svg>
          </button>
        </div>
      </div>
      
      <EventsGrid :events="events" @select-event="showEventDetails" />
      
      <EventDetails
        v-if="selectedEvent"
        :event="selectedEvent"
        :is-dark-mode="isDarkMode"
        @close="closeEventDetails"
        @toggleFavorite="toggleFavorite"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import Navbar from '~/components/Navbar.vue';
import EventsGrid from '~/components/EventsGrid.vue';
import EventDetails from '~/components/EventDetails.vue';

const events = ref([]);
const selectedEvent = ref(null);
const showSuccess = ref(false);
const favorites = ref([]);

const theme = useState('theme');
const isDarkMode = computed(() => theme.value === 'dark');

onMounted(() => {
  generateEvents();
  loadFavorites();
});

function generateEvents() {
  const categories = ['tech', 'music', 'art'];
  const tags = [
    ['веб', 'конференция', 'программирование'],
    ['концерт', 'живая музыка', 'джаз'],
    ['выставка', 'современное искусство', 'галерея'],
    ['нетворкинг', 'бизнес', 'стартапы'],
    ['семинар', 'образование', 'мастер-класс']
  ];
  
  const locations = [
    { city: 'Москва', venue: 'Центр "Космос"' },
    { city: 'Санкт-Петербург', venue: 'Манеж' },
    { city: 'Казань', venue: 'IT Park' },
    { city: 'Екатеринбург', venue: 'Ельцин Центр' },
    { city: 'Новосибирск', venue: 'Технопарк' }
  ];
  
  const titles = [
    'DevConf 2023',
    'Джазовый вечер',
    'Выставка цифрового искусства',
    'Нетворкинг для IT-специалистов',
    'Мастер-класс по UI/UX дизайну',
    'Киберспортивный турнир',
    'Конференция по искусственному интеллекту',
    'Фестиваль электронной музыки',
    'Открытый форум разработчиков',
    'Хакатон: Умный город',
    'Встреча любителей блокчейна',
    'Дизайн-маркет',
    'Архитектурная выставка'
  ];
  
  const descriptions = [
    'Встреча профессионалов индустрии для обсуждения последних тенденций разработки.',
    'Атмосферный вечер с живой музыкой от лучших музыкантов города.',
    'Уникальная выставка, сочетающая традиционное и цифровое искусство.',
    'Возможность найти единомышленников и расширить профессиональные связи.',
    'Интерактивный мастер-класс с опытными спикерами и практическими упражнениями.',
    'Соревнование для киберспортсменов с ценными призами и онлайн-трансляцией.'
  ];

  const newEvents = [];
  for (let i = 0; i < 15; i++) {
    const randomDate = new Date();
    randomDate.setDate(randomDate.getDate() + Math.floor(Math.random() * 30));
    randomDate.setHours(Math.floor(Math.random() * 12) + 9, Math.floor(Math.random() * 60));

    const category = categories[Math.floor(Math.random() * categories.length)];
    const eventTags = tags[Math.floor(Math.random() * tags.length)];
    const location = locations[Math.floor(Math.random() * locations.length)];
    const title = titles[Math.floor(Math.random() * titles.length)];
    const description = descriptions[Math.floor(Math.random() * descriptions.length)];

    newEvents.push({
      id: i + 1,
      title: `${title} ${i + 1}`,
      date: randomDate.toISOString(),
      category,
      tags: eventTags,
      location,
      description,
      organizer: 'Event SpheRRe',
      price: Math.floor(Math.random() * 500) * 10
    });
  }

  events.value = newEvents;
}

function loadFavorites() {
  if (process.client) {
    const saved = localStorage.getItem('favorites');
    if (saved) {
      favorites.value = JSON.parse(saved);
    }
  }
}

function saveFavorites() {
  if (process.client) {
    localStorage.setItem('favorites', JSON.stringify(favorites.value));
  }
}

function toggleFavorite(eventId) {
  const index = favorites.value.indexOf(eventId);
  
  if (index === -1) {
    favorites.value.push(eventId);
    showSuccess.value = true;
  } else {
    favorites.value.splice(index, 1);
  }
  
  saveFavorites();
  
  setTimeout(() => {
    showSuccess.value = false;
  }, 3000);
}

function showEventDetails(event) {
  selectedEvent.value = event;
}

function closeEventDetails() {
  selectedEvent.value = null;
}

const toggleTheme = () => {
  theme.value = isDarkMode.value ? 'light' : 'dark';
};
</script>

<style scoped>
.action-btn {
  background: var(--purple);
  color: #FFFFFF;
  transition: all 0.2s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.action-btn:hover {
  transform: translateY(-1px);
  box-shadow: 0 4px 8px rgba(108, 99, 255, 0.3);
}
</style> 