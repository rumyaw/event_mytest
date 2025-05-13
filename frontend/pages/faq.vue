<template>
<Navbar />
    <div class="faq-container" style="padding-top: 80px">
        <div class="faq-header">
        <h1>Часто задаваемые вопросы</h1>
        <p>Найдите ответы на популярные вопросы или свяжитесь с нами</p>
        </div>

        <div class="faq-content">
        <div class="faq-list">
            <div 
            v-for="(item, index) in faqItems" 
            :key="index" 
            class="faq-item"
            :class="{ 'active': activeIndex === index }"
            @click="toggleFaq(index)"
            >
            <div class="faq-question">
                <h3>{{ item.question }}</h3>
                <svg 
                xmlns="http://www.w3.org/2000/svg" 
                width="20" 
                height="20" 
                viewBox="0 0 24 24" 
                fill="none" 
                stroke="currentColor" 
                stroke-width="2" 
                stroke-linecap="round" 
                stroke-linejoin="round"
                :class="{ 'rotate-180': activeIndex === index }"
                >
                <polyline points="6 9 12 15 18 9"></polyline>
                </svg>
            </div>
            <transition name="slide-down">
                <div v-if="activeIndex === index" class="faq-answer">
                <p>{{ item.answer }}</p>
                </div>
            </transition>
            </div>
        </div>

        <div class="contact-form">
            <h2>Не нашли ответ?</h2>
            <p>Отправьте нам свой вопрос</p>
            
            <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="name">Ваше имя</label>
                <input 
                type="text" 
                id="name" 
                v-model="form.name" 
                placeholder="Иван Иванов" 
                required
                >
            </div>
            
            <div class="form-group">
                <label for="email">Email</label>
                <input 
                type="email" 
                id="email" 
                v-model="form.email" 
                placeholder="example@mail.com" 
                required
                >
            </div>
            
            <div class="form-group">
                <label for="question">Ваш вопрос</label>
                <textarea 
                id="question" 
                v-model="form.question" 
                placeholder="Опишите ваш вопрос подробнее..." 
                rows="4" 
                required
                ></textarea>
            </div>
            
            <button type="submit" class="submit-btn">
                Отправить вопрос
            </button>
            </form>
        </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'

const activeIndex = ref(null)
const form = ref({
  name: '',
  email: '',
  question: ''
})

const faqItems = ref([
  {
    question: 'Как создать мероприятие?',
    answer: 'Для создания мероприятия перейдите в раздел "Мои мероприятия" и нажмите кнопку "Создать". Заполните все необходимые поля и сохраните.'
  },
  {
    question: 'Как пригласить друзей?',
    answer: 'В карточке мероприятия есть кнопка "Пригласить". Вы можете отправить приглашение через соцсети или скопировать ссылку.'
  },
  {
    question: 'Можно ли отменить регистрацию?',
    answer: 'Да, вы можете отменить регистрацию в любое время в разделе "Мои мероприятия".'
  },
  {
    question: 'Как связаться с организатором?',
    answer: 'На странице мероприятия есть контакты организатора. Также вы можете написать сообщение через наш сервис.'
  },
  {
    question: 'Есть ли мобильное приложение?',
    answer: 'Да, наше приложение доступно в App Store и Google Play. Скачайте его для более удобного использования сервиса.'
  }
])

const toggleFaq = (index) => {
  activeIndex.value = activeIndex.value === index ? null : index
}

const submitForm = () => {
  // Here would be the form submission logic
  console.log('Form submitted:', form.value)
  alert('Ваш вопрос отправлен! Мы свяжемся с вами в ближайшее время.')
  form.value = { name: '', email: '', question: '' }
}
</script>

<style scoped>
.faq-container {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem 1rem;
}

.faq-header {
  text-align: center;
  margin-bottom: 2rem;
}

.faq-header h1 {
  font-size: 2rem;
  font-weight: 600;
  color: #E4E4F0;
  margin-bottom: 0.5rem;
}

.faq-header p {
  color: rgba(228, 228, 240, 0.7);
  font-size: 1rem;
}

.faq-content {
  display: grid;
  gap: 2rem;
}

.faq-list {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.faq-item {
  background: rgba(42, 42, 64, 0.8);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  padding: 1rem;
  border: 1px solid rgba(228, 228, 240, 0.1);
  transition: all 0.3s ease;
}

.faq-item.active {
  background: rgba(58, 58, 80, 0.8);
}

/* Общие стили */
.nav-bar {
  position: sticky;
  top: 0;
  z-index: 10;
  transition: all 0.3s ease;
}
.faq-question {
  display: flex;
  justify-content: space-between;
  align-items: center;
  cursor: pointer;
}

.faq-question h3 {
  font-size: 1rem;
  font-weight: 500;
  color: #E4E4F0;
}

.faq-question svg {
  transition: transform 0.3s ease;
  color: #6C63FF;
}

.faq-question svg.rotate-180 {
  transform: rotate(180deg);
}

.faq-answer {
  padding-top: 1rem;
  margin-top: 1rem;
  border-top: 1px solid rgba(228, 228, 240, 0.1);
}

.faq-answer p {
  color: rgba(228, 228, 240, 0.8);
  font-size: 0.9rem;
  line-height: 1.5;
}

.contact-form {
  background: rgba(42, 42, 64, 0.8);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  padding: 1.5rem;
  border: 1px solid rgba(228, 228, 240, 0.1);
}

.contact-form h2 {
  font-size: 1.25rem;
  font-weight: 600;
  color: #E4E4F0;
  margin-bottom: 0.5rem;
}

.contact-form p {
  color: rgba(228, 228, 240, 0.7);
  font-size: 0.9rem;
  margin-bottom: 1.5rem;
}

.form-group {
  margin-bottom: 1rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  color: #E4E4F0;
  font-size: 0.9rem;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 0.75rem;
  border-radius: 8px;
  border: 1px solid rgba(228, 228, 240, 0.2);
  background: rgba(58, 58, 80, 0.5);
  color: #E4E4F0;
  font-size: 0.9rem;
  transition: all 0.2s ease;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #6C63FF;
  box-shadow: 0 0 0 2px rgba(108, 99, 255, 0.2);
}

.submit-btn {
  width: 100%;
  padding: 0.75rem;
  border-radius: 8px;
  background: #6C63FF;
  color: white;
  font-weight: 500;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease;
  margin-top: 1rem;
}

.submit-btn:hover {
  background: #5a52e0;
  transform: translateY(-1px);
}

.submit-btn:active {
  transform: translateY(0);
}

.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.3s ease;
  max-height: 200px;
  overflow: hidden;
}

.slide-down-enter-from,
.slide-down-leave-to {
  max-height: 0;
  opacity: 0;
}

/* Light theme styles */
.light-theme .faq-item,
.light-theme .contact-form {
  background: rgba(255, 255, 255, 0.9);
  border-color: rgba(26, 26, 46, 0.1);
}

.light-theme .faq-item.active {
  background: rgba(240, 240, 245, 0.9);
}

.light-theme .faq-header h1,
.light-theme .faq-question h3,
.light-theme .contact-form h2,
.light-theme .form-group label {
  color: #1A1A2E;
}

.light-theme .faq-header p,
.light-theme .contact-form p,
.light-theme .faq-answer p {
  color: rgba(26, 26, 46, 0.7);
}

.light-theme .form-group input,
.light-theme .form-group textarea {
  background: rgba(240, 240, 245, 0.8);
  border-color: rgba(26, 26, 46, 0.2);
  color: #1A1A2E;
}

.light-theme .faq-answer {
  border-top-color: rgba(26, 26, 46, 0.1);
}
</style>
