<template>
  <div class="fixed right-4 top-20 z-50 transition-all duration-300" :class="[isOpen ? 'translate-x-0' : 'translate-x-[105%]']">
    <div :class="['p-4 rounded-lg shadow-xl w-80 max-h-[80vh] overflow-y-auto', 
                 isDarkMode ? 'dark-theme-bg' : 'light-theme-bg']">
      
      <!-- Header -->
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-semibold">Уведомления</h3>
        <div class="flex gap-2">
          <button @click="showSubscriptions = !showSubscriptions" class="p-1 rounded-full hover:bg-opacity-10 hover:bg-black">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z" clip-rule="evenodd" />
            </svg>
          </button>
          <button @click="markAllAsRead" class="p-1 rounded-full hover:bg-opacity-10 hover:bg-black">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
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
      <div v-if="!showSubscriptions">
        <!-- Tabs -->
        <div class="flex mb-4 border-b" :class="isDarkMode ? 'border-gray-700' : 'border-gray-300'">
          <button 
            @click="activeTab = 'all'" 
            class="px-3 py-2 text-sm font-medium transition-colors -mb-px"
            :class="activeTab === 'all' ? 
              (isDarkMode ? 'text-accent border-b-2 border-accent' : 'text-accent border-b-2 border-accent') : 
              (isDarkMode ? 'text-gray-400 hover:text-gray-200' : 'text-gray-500 hover:text-gray-800')"
          >
            Все
          </button>
          <button 
            @click="activeTab = 'unread'" 
            class="px-3 py-2 text-sm font-medium transition-colors -mb-px"
            :class="activeTab === 'unread' ? 
              (isDarkMode ? 'text-accent border-b-2 border-accent' : 'text-accent border-b-2 border-accent') : 
              (isDarkMode ? 'text-gray-400 hover:text-gray-200' : 'text-gray-500 hover:text-gray-800')"
          >
            Непрочитанные
          </button>
        </div>
        
        <!-- Notifications list -->
        <div class="space-y-3">
          <TransitionGroup name="notification-list">
            <div 
              v-for="notification in filteredNotifications" 
              :key="notification.id"
              :class="['p-3 rounded-lg text-sm relative transition-all', 
                     notification.read ? 'opacity-70' : 'opacity-100',
                     isDarkMode ? 'bg-gray-800 hover:bg-gray-700' : 'bg-gray-100 hover:bg-gray-200']"
            >
              <!-- Notification badge -->
              <div 
                v-if="!notification.read" 
                class="absolute w-2 h-2 rounded-full bg-accent top-3 right-3"
              ></div>
              
              <!-- Notification content -->
              <div class="flex items-start gap-3">
                <div class="flex-shrink-0">
                  <div :class="['p-2 rounded-full', getNotificationTypeClass(notification.type)]">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                      <path v-if="notification.type === 'event'" fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd" />
                      <path v-else-if="notification.type === 'comment'" fill-rule="evenodd" d="M18 10c0 3.866-3.582 7-8 7a8.841 8.841 0 01-4.083-.98L2 17l1.338-3.123C2.493 12.767 2 11.434 2 10c0-3.866 3.582-7 8-7s8 3.134 8 7zM7 9H5v2h2V9zm8 0h-2v2h2V9zM9 9h2v2H9V9z" clip-rule="evenodd" />
                      <path v-else-if="notification.type === 'subscription'" d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" />
                      <path v-else d="M10 2a6 6 0 00-6 6v3.586l-.707.707A1 1 0 004 14h12a1 1 0 00.707-1.707L16 11.586V8a6 6 0 00-6-6zM10 18a3 3 0 01-3-3h6a3 3 0 01-3 3z" />
                    </svg>
                  </div>
                </div>
                <div class="flex-1">
                  <div class="font-medium">{{ notification.title }}</div>
                  <p>{{ notification.message }}</p>
                  <div class="text-xs mt-1 opacity-60">{{ formatDate(notification.date) }}</div>
                  
                  <!-- Action buttons if needed -->
                  <div v-if="notification.actionable" class="mt-2 flex gap-2">
                    <button 
                      @click="handleNotificationAction(notification)" 
                      class="px-3 py-1 text-xs rounded-lg bg-accent bg-opacity-20 text-accent hover:bg-opacity-30 transition-colors"
                    >
                      {{ notification.actionLabel || 'Смотреть' }}
                    </button>
                    <button 
                      @click="dismissNotification(notification.id)" 
                      class="px-3 py-1 text-xs rounded-lg bg-gray-500 bg-opacity-20 hover:bg-opacity-30 transition-colors"
                    >
                      Закрыть
                    </button>
                  </div>
                </div>
              </div>
            </div>
            
            <div v-if="filteredNotifications.length === 0" key="empty" class="text-center py-8 opacity-60">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-2 opacity-40" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
              </svg>
              <p>Нет уведомлений</p>
            </div>
          </TransitionGroup>
        </div>
      </div>
      
      <!-- Subscriptions panel -->
      <div v-else>
        <div class="mb-4 pb-2 border-b flex justify-between items-center" 
             :class="isDarkMode ? 'border-gray-700' : 'border-gray-300'">
          <h4 class="font-medium">Ваши подписки</h4>
          <div class="text-xs">
            <label class="flex items-center">
              <input type="checkbox" v-model="notificationsEnabled" class="mr-2 accent-purple">
              Уведомления
            </label>
          </div>
        </div>
        
        <!-- Subscriptions list -->
        <div class="space-y-3">
          <div v-for="sub in subscriptions" :key="sub.id"
               :class="['p-3 rounded-lg flex justify-between items-center', 
                       isDarkMode ? 'bg-gray-800' : 'bg-gray-100']">
            <div>
              <div class="font-medium">{{ sub.name }}</div>
              <div class="text-xs opacity-70">{{ sub.eventsCount }} мероприятий</div>
            </div>
            <div class="flex gap-2">
              <button @click="toggleSubscription(sub.id)" :class="[
                'px-3 py-1 text-xs rounded-full transition-colors',
                sub.isSubscribed 
                  ? 'bg-accent bg-opacity-20 text-accent' 
                  : isDarkMode ? 'bg-gray-700' : 'bg-gray-300'
              ]">
                {{ sub.isSubscribed ? 'Отписаться' : 'Подписаться' }}
              </button>
            </div>
          </div>
          
          <div v-if="subscriptions.length === 0" class="text-center py-8 opacity-60">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 mx-auto mb-2 opacity-40" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
            </svg>
            <p>У вас пока нет подписок</p>
          </div>
        </div>
        
        <!-- Discover organizers -->
        <div class="mt-6">
          <h4 class="font-medium mb-3">Популярные организаторы</h4>
          <div class="space-y-3">
            <div v-for="organizer in recommendedOrganizers" :key="organizer.id"
                 :class="['p-3 rounded-lg flex justify-between items-center', 
                         isDarkMode ? 'bg-gray-800' : 'bg-gray-100']">
              <div>
                <div class="font-medium">{{ organizer.name }}</div>
                <div class="text-xs opacity-70">{{ organizer.category }}</div>
              </div>
              <div>
                <button @click="subscribeToOrganizer(organizer.id)" :class="[
                  'px-3 py-1 text-xs rounded-full transition-colors',
                  isDarkMode ? 'bg-gray-700 hover:bg-gray-600' : 'bg-gray-300 hover:bg-gray-400'
                ]">
                  Подписаться
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      
    </div>
    
    <!-- Toggle button -->
    <button 
      @click="togglePanel"
      :class="['absolute -left-12 p-2 rounded-l-lg shadow-lg transition-colors',
              isDarkMode ? 'bg-gray-800 hover:bg-gray-700' : 'bg-white hover:bg-gray-100',
              isOpen ? 'hidden' : '']"
    >
      <div class="relative">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
        </svg>
        <div 
          v-if="unreadCount > 0"
          class="absolute -top-1 -right-1 w-4 h-4 flex items-center justify-center rounded-full bg-accent text-white text-xs"
        >
          {{ unreadCount > 9 ? '9+' : unreadCount }}
        </div>
      </div>
    </button>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';

// State
const isOpen = ref(false);
const activeTab = ref('all');
const notifications = ref([]);
const subscriptions = ref([]);
const recommendedOrganizers = ref([]);
const showSubscriptions = ref(false);
const notificationsEnabled = ref(true);

// Props and emit
const props = defineProps({
  initialOpen: {
    type: Boolean,
    default: false
  }
});

const emits = defineEmits(['close', 'notificationAction']);

// Theme
const theme = useState('theme', () => 'dark');
const isDarkMode = computed(() => theme.value === 'dark');

// Computed
const filteredNotifications = computed(() => {
  if (activeTab.value === 'unread') {
    return notifications.value.filter(notification => !notification.read);
  }
  return notifications.value;
});

const unreadCount = computed(() => {
  return notifications.value.filter(notification => !notification.read).length;
});

// Methods
function togglePanel() {
  isOpen.value = !isOpen.value;
}

function closePanel() {
  isOpen.value = false;
  emits('close');
}

function markAllAsRead() {
  notifications.value = notifications.value.map(notification => {
    return { ...notification, read: true };
  });
}

function dismissNotification(id) {
  notifications.value = notifications.value.filter(notification => notification.id !== id);
}

function handleNotificationAction(notification) {
  // Mark as read
  const index = notifications.value.findIndex(n => n.id === notification.id);
  if (index !== -1) {
    notifications.value[index].read = true;
  }
  
  // Emit action event
  emits('notificationAction', notification);
}

function getNotificationTypeClass(type) {
  const classes = {
    'event': 'bg-purple-500 bg-opacity-20 text-purple-500',
    'comment': 'bg-green-500 bg-opacity-20 text-green-500',
    'subscription': 'bg-blue-500 bg-opacity-20 text-blue-500',
    'system': 'bg-gray-500 bg-opacity-20 text-gray-500'
  };
  
  return classes[type] || classes.system;
}

function formatDate(dateString) {
  const date = new Date(dateString);
  const now = new Date();
  const diffMs = now - date;
  const diffMins = Math.floor(diffMs / (1000 * 60));
  const diffHours = Math.floor(diffMs / (1000 * 60 * 60));
  const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));
  
  if (diffMins < 60) {
    return `${diffMins} мин. назад`;
  } else if (diffHours < 24) {
    return `${diffHours} ч. назад`;
  } else if (diffDays < 7) {
    return `${diffDays} дн. назад`;
  } else {
    return date.toLocaleDateString('ru-RU');
  }
}

function toggleSubscription(organizerId) {
  const index = subscriptions.value.findIndex(sub => sub.id === organizerId);
  if (index !== -1) {
    subscriptions.value[index].isSubscribed = !subscriptions.value[index].isSubscribed;
    
    // Add notification about subscription change
    if (subscriptions.value[index].isSubscribed) {
      addNotification({
        id: Date.now(),
        type: 'subscription',
        title: 'Новая подписка',
        message: `Вы подписались на организатора "${subscriptions.value[index].name}"`,
        date: new Date().toISOString(),
        read: false
      });
    }
  }
}

function subscribeToOrganizer(organizerId) {
  const organizer = recommendedOrganizers.value.find(org => org.id === organizerId);
  if (organizer) {
    const newSubscription = {
      id: organizer.id,
      name: organizer.name,
      eventsCount: organizer.eventsCount || Math.floor(Math.random() * 10) + 1,
      isSubscribed: true
    };
    
    subscriptions.value.push(newSubscription);
    recommendedOrganizers.value = recommendedOrganizers.value.filter(org => org.id !== organizerId);
    
    // Add notification
    addNotification({
      id: Date.now(),
      type: 'subscription',
      title: 'Новая подписка',
      message: `Вы подписались на организатора "${organizer.name}"`,
      date: new Date().toISOString(),
      read: false
    });
  }
}

function addNotification(notification) {
  notifications.value.unshift(notification);
}

// Init
onMounted(() => {
  isOpen.value = props.initialOpen;
  
  // Mock data
  mockData();
});

function mockData() {
  // Sample notifications
  notifications.value = [
    {
      id: 1,
      type: 'event',
      title: 'Новое мероприятие',
      message: 'Организатор "ТехноСфера" добавил новое мероприятие "IT Конференция 2025"',
      date: new Date(Date.now() - 1000 * 60 * 30).toISOString(), // 30 mins ago
      read: false,
      actionable: true,
      actionLabel: 'Посмотреть'
    },
    {
      id: 2,
      type: 'comment',
      title: 'Новый комментарий',
      message: 'Пользователь johndoe прокомментировал мероприятие "Музыкальный фестиваль"',
      date: new Date(Date.now() - 1000 * 60 * 60 * 2).toISOString(), // 2 hours ago
      read: true,
      actionable: true
    },
    {
      id: 3,
      type: 'subscription',
      title: 'Обновление подписки',
      message: 'Организатор "АртГалерея" добавил новое мероприятие',
      date: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString(), // 1 day ago
      read: false,
      actionable: true
    },
    {
      id: 4,
      type: 'system',
      title: 'Добро пожаловать!',
      message: 'Спасибо за регистрацию на платформе',
      date: new Date(Date.now() - 1000 * 60 * 60 * 24 * 3).toISOString(), // 3 days ago
      read: true
    }
  ];
  
  // Sample subscriptions
  subscriptions.value = [
    {
      id: 101,
      name: 'ТехноСфера',
      eventsCount: 15,
      isSubscribed: true
    },
    {
      id: 102,
      name: 'АртГалерея',
      eventsCount: 8,
      isSubscribed: true
    },
    {
      id: 103,
      name: 'МузыкаПлюс',
      eventsCount: 12,
      isSubscribed: false
    }
  ];
  
  // Sample recommended organizers
  recommendedOrganizers.value = [
    {
      id: 201,
      name: 'Спорт-Ивент',
      category: 'Спортивные мероприятия',
      eventsCount: 7
    },
    {
      id: 202,
      name: 'Научный Центр',
      category: 'Образовательные мероприятия',
      eventsCount: 4
    },
    {
      id: 203,
      name: 'GameDev Сообщество',
      category: 'Игровая индустрия',
      eventsCount: 9
    }
  ];
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

/* Notification list transitions */
.notification-list-enter-active,
.notification-list-leave-active {
  transition: all 0.3s ease;
}

.notification-list-enter-from,
.notification-list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.notification-list-move {
  transition: transform 0.3s ease;
}
</style>
