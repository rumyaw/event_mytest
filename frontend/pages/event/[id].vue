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
            <button 
              @click="toggleSubscribe"
              :class="['ml-3 text-xs px-3 py-1 rounded-full transition-colors',
                     isSubscribed ? 'bg-[#6C63FF]/20 text-[#6C63FF]' : 'bg-gray-200/30 hover:bg-[#6C63FF]/20']"
            >
              {{ isSubscribed ? 'Подписан' : 'Подписаться' }}
            </button>
          </div>
          <div class="flex items-center">
            <span class="mr-2">📍</span>
            <p>Формат: Офлайн | Место: Москва, ул. Ленина, 10</p>
          </div>
        </div>

        <div class="map-container mt-6 h-64 rounded-xl flex items-center justify-center border-2 border-dashed">
          <p>Здесь будет карта</p>
        </div>

        <!-- Rating section -->
        <div class="p-6 border-t border-opacity-10" :class="isDarkMode ? 'border-white' : 'border-gray-800'">
          <h2 class="text-xl font-semibold mb-4">Оценка мероприятия</h2>
          
          <div class="flex items-center mb-6">
            <div class="flex">
              <button v-for="star in 5" :key="`rating-star-${star}`" 
                      @click="setRating(star)" 
                      class="focus:outline-none">
                <svg class="w-8 h-8" :class="star <= userRating ? 'text-yellow-400' : 'text-gray-400'"
                     xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z" />
                </svg>
              </button>
            </div>
            <div class="ml-3">
              <span class="text-lg font-medium">{{ eventRating.toFixed(1) }}</span>
              <span class="text-sm opacity-70 ml-1">({{ ratingCount }} оценок)</span>
            </div>
          </div>
        </div>
        
        <!-- Comments section -->
        <div class="p-6 border-t border-opacity-10" :class="isDarkMode ? 'border-white' : 'border-gray-800'">
          <h2 class="text-xl font-semibold mb-4">Комментарии ({{ comments.length }})</h2>
          
          <!-- Comment form -->
          <div class="mb-6">
            <textarea v-model="newComment" 
                      placeholder="Оставьте свой комментарий..." 
                      :class="['w-full p-3 rounded-xl transition-colors resize-none',
                              isDarkMode ? 'bg-[#3A3A50] focus:bg-[#434360]' : 'bg-gray-100 focus:bg-white']"
                      rows="3"></textarea>
            <div class="flex justify-end mt-2">
              <button @click="addComment" 
                      :disabled="!newComment.trim()"
                      :class="['px-4 py-2 rounded-lg transition-colors',
                              newComment.trim() ? 'bg-accent text-white' : 'bg-gray-400 opacity-50 cursor-not-allowed']">
                Отправить
              </button>
            </div>
          </div>
          
          <!-- Comments list -->
          <div v-if="comments.length > 0" class="space-y-4 max-h-96 overflow-y-auto pr-2">
            <div v-for="(comment, index) in comments" :key="index"
                 :class="['p-4 rounded-xl transition-colors', 
                         isDarkMode ? 'bg-[#3A3A50]' : 'bg-gray-100']">
              <div class="flex justify-between mb-2">
                <span class="font-semibold">{{ comment.author }}</span>
                <span class="text-sm opacity-60">{{ formatCommentDate(comment.date) }}</span>
              </div>
              <p>{{ comment.text }}</p>
              
              <!-- Comment actions -->
              <div class="flex items-center mt-3 text-sm">
                <button @click="likeComment(index)" class="flex items-center opacity-70 hover:opacity-100 transition-opacity">
                  <svg :class="['w-4 h-4 mr-1', comment.liked ? 'text-red-500' : '']" 
                       xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M3.172 5.172a4 4 0 015.656 0L10 6.343l1.172-1.171a4 4 0 115.656 5.656L10 17.657l-6.828-6.829a4 4 0 010-5.656z" clip-rule="evenodd" />
                  </svg>
                  {{ comment.likes }}
                </button>
                <button @click="showReplyForm = index" class="ml-4 opacity-70 hover:opacity-100 transition-opacity">
                  Ответить
                </button>
                <button v-if="comment.isMine" @click="deleteComment(index)" class="ml-auto opacity-70 hover:opacity-100 transition-opacity text-red-400">
                  Удалить
                </button>
              </div>
              
              <!-- Reply form -->
              <div v-if="showReplyForm === index" class="mt-3">
                <div class="flex">
                  <input v-model="replyText" 
                         :class="['flex-1 p-2 text-sm rounded-l-lg',
                                isDarkMode ? 'bg-[#2A2A40]' : 'bg-white border']"
                         placeholder="Ответить на комментарий...">
                  <button @click="addReply(index)" 
                          :disabled="!replyText.trim()"
                          :class="['px-3 py-1 rounded-r-lg transition-colors',
                                  replyText.trim() ? 'bg-accent text-white' : 'bg-gray-400 opacity-50']">
                    Отправить
                  </button>
                </div>
              </div>
              
              <!-- Replies -->
              <div v-if="comment.replies && comment.replies.length > 0" class="mt-3 ml-4 pl-3 border-l-2 border-gray-300 border-opacity-30 space-y-3">
                <div v-for="(reply, replyIndex) in comment.replies" :key="`reply-${index}-${replyIndex}`"
                     class="text-sm">
                  <div class="flex justify-between">
                    <span class="font-semibold">{{ reply.author }}</span>
                    <span class="opacity-60">{{ formatCommentDate(reply.date) }}</span>
                  </div>
                  <p>{{ reply.text }}</p>
                </div>
              </div>
            </div>
          </div>
          
          <div v-else class="text-center py-8 opacity-60">
            <p>Будьте первым, кто оставит комментарий</p>
          </div>
        </div>
        
        <!-- Sharing section -->
        <div class="p-6 border-t border-opacity-10" :class="isDarkMode ? 'border-white' : 'border-gray-800'">
          <h2 class="text-xl font-semibold mb-4">Поделиться</h2>
          
          <div class="flex flex-wrap gap-3">
            <button @click="shareEvent('vk')" 
                    class="flex items-center px-4 py-2 rounded-lg transition-colors bg-[#4C75A3]/20 hover:bg-[#4C75A3]/30 text-[#4C75A3]">
              <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
                <path d="M21.547 7h-3.29a.743.743 0 0 0-.655.392s-1.312 2.416-1.734 3.23c-1.43 2.85-2.053 3.377-2.35 3.233-.565-.237-.494-1.9-.494-2.826V7.183c0-.575-.213-.783-.778-.783h-3.643c-.421 0-.678.267-.678.584 0 .541.82.717.893 2.26v3.253c.038.967-.166 1.179-.557 1.179-1.012 0-3.572-3.468-5.08-7.46C3.113 6.494 2.95 6 2.357 6H.803C.136 6 0 6.416 0 6.761c0 .6.096 3.627 4.537 7.85C7.237 17.056 10.775 18 13.927 18c1.82 0 2.112-.167 2.112-.908v-2.058c0-.67.142-.804.614-.804.37 0 1.002.176 2.476 1.607C20.857 17.446 21.046 18 21.895 18h3.291c.667 0 .998-.271.808-.817-.207-.632-2.436-2.855-2.539-2.988-.347-.407-.247-.631 0-1.024.173-.291 1.16-1.72 1.798-2.733.515-.814.902-1.493 1.097-1.926.184-.402.164-.662-.45-.902"></path>
              </svg>
              ВКонтакте
            </button>
            
            <button @click="shareEvent('telegram')" 
                    class="flex items-center px-4 py-2 rounded-lg transition-colors bg-[#0088cc]/20 hover:bg-[#0088cc]/30 text-[#0088cc]">
              <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm4.64 6.8c-.15 1.58-.8 5.42-1.13 7.19-.14.75-.42 1-.68 1.03-.58.05-1.02-.38-1.58-.75-.88-.58-1.38-.94-2.23-1.5-.99-.65-.35-1.01.22-1.59.15-.15 2.71-2.48 2.76-2.69.01-.03.01-.14-.07-.2-.08-.06-.19-.04-.27-.02-.12.03-1.89 1.2-5.31 3.53-.5.34-.96.51-1.37.5-.45-.02-1.32-.26-1.97-.47-.8-.26-1.43-.4-1.37-.85.03-.22.27-.45.74-.68 2.87-1.25 4.79-2.07 5.76-2.46 2.74-1.12 3.31-1.31 3.69-1.32.08 0 .27.02.39.12.1.08.13.19.14.27-.01.06.01.24 0 .38z"/>
              </svg>
              Telegram
            </button>
            
            <button @click="shareEvent('whatsapp')" 
                    class="flex items-center px-4 py-2 rounded-lg transition-colors bg-[#25D366]/20 hover:bg-[#25D366]/30 text-[#25D366]">
              <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="currentColor">
                <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
              </svg>
              WhatsApp
            </button>
            
            <button @click="copyEventLink" 
                    class="flex items-center px-4 py-2 rounded-lg transition-colors bg-gray-600/20 hover:bg-gray-600/30">
              <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
              </svg>
              Копировать ссылку
            </button>
          </div>
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

    <!-- Notification toast for share -->
    <div v-if="showToast" 
         class="fixed bottom-20 left-1/2 transform -translate-x-1/2 px-4 py-2 rounded-lg bg-accent text-white shadow-lg fade-in">
      {{ toastMessage }}
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue'
import Navbar from '~/components/Navbar.vue'

const isDarkMode = ref(true)
const toggleTheme = () => {
  isDarkMode.value = !isDarkMode.value
}

// Event details
const route = useRoute();
const eventId = computed(() => route.params.id);

// Subscription state
const isSubscribed = ref(false);
const toggleSubscribe = () => {
  isSubscribed.value = !isSubscribed.value;
  if (isSubscribed.value) {
    showToastMessage('Вы подписались на организатора "ТехноСфера"');
  }
};

// Rating system
const userRating = ref(0);
const allRatings = ref([4.2, 4.7, 3.8, 4.5, 5, 4.1, 4.3]);
const ratingCount = computed(() => allRatings.value.length + (userRating.value > 0 ? 1 : 0));
const eventRating = computed(() => {
  if (userRating.value > 0) {
    return (allRatings.value.reduce((a, b) => a + b, 0) + userRating.value) / ratingCount.value;
  }
  return allRatings.value.reduce((a, b) => a + b, 0) / allRatings.value.length;
});

function setRating(rating) {
  userRating.value = rating;
  showToastMessage('Спасибо за вашу оценку!');
}

// Comments system
const comments = ref([
  {
    author: 'Александр',
    text: 'Очень жду эту конференцию! Надеюсь, будут интересные темы по искусственному интеллекту.',
    date: new Date(Date.now() - 1000 * 60 * 60 * 24 * 2).toISOString(),
    likes: 5,
    liked: false,
    isMine: false,
    replies: [
      {
        author: 'Мария',
        text: 'Согласна, особенно интересует применение ИИ в медицине.',
        date: new Date(Date.now() - 1000 * 60 * 60 * 24).toISOString()
      }
    ]
  },
  {
    author: 'Ирина',
    text: 'А будет ли онлайн-трансляция для тех, кто не сможет приехать в Москву?',
    date: new Date(Date.now() - 1000 * 60 * 60 * 4).toISOString(),
    likes: 2,
    liked: false,
    isMine: false,
    replies: []
  }
]);
const newComment = ref('');
const showReplyForm = ref(null);
const replyText = ref('');

function addComment() {
  if (!newComment.value.trim()) return;
  
  comments.value.unshift({
    author: 'Вы',
    text: newComment.value,
    date: new Date().toISOString(),
    likes: 0,
    liked: false,
    isMine: true,
    replies: []
  });
  
  newComment.value = '';
  showToastMessage('Комментарий добавлен');
}

function deleteComment(index) {
  comments.value.splice(index, 1);
  showToastMessage('Комментарий удален');
}

function likeComment(index) {
  const comment = comments.value[index];
  comment.liked = !comment.liked;
  comment.likes += comment.liked ? 1 : -1;
}

function addReply(commentIndex) {
  if (!replyText.value.trim()) return;
  
  comments.value[commentIndex].replies.push({
    author: 'Вы',
    text: replyText.value,
    date: new Date().toISOString()
  });
  
  replyText.value = '';
  showReplyForm.value = null;
  showToastMessage('Ответ добавлен');
}

function formatCommentDate(dateString) {
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

// Toast notification
const showToast = ref(false);
const toastMessage = ref('');

function showToastMessage(message) {
  toastMessage.value = message;
  showToast.value = true;
  
  setTimeout(() => {
    showToast.value = false;
  }, 3000);
}

// Social sharing
function shareEvent(platform) {
  const eventUrl = `${window.location.origin}/event/${eventId.value}`;
  const title = 'IT Конференция 2025';
  const text = 'Самое масштабное IT событие 2025 года!';
  
  let shareUrl;
  
  switch(platform) {
    case 'telegram':
      shareUrl = `https://t.me/share/url?url=${encodeURIComponent(eventUrl)}&text=${encodeURIComponent(title)}`;
      break;
    case 'vk':
      shareUrl = `https://vk.com/share.php?url=${encodeURIComponent(eventUrl)}&title=${encodeURIComponent(title)}&description=${encodeURIComponent(text)}`;
      break;
    case 'whatsapp':
      shareUrl = `https://api.whatsapp.com/send?text=${encodeURIComponent(title + ' ' + eventUrl)}`;
      break;
  }
  
  if (shareUrl) {
    window.open(shareUrl, '_blank');
  }
}

function copyEventLink() {
  const eventUrl = `${window.location.origin}/event/${eventId.value}`;
  navigator.clipboard.writeText(eventUrl).then(() => {
    showToastMessage('Ссылка скопирована в буфер обмена');
  });
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
