<template>
  <div
    class="min-h-screen flex items-center justify-center relative overflow-hidden"
    :style="{ backgroundColor: '#1A1A2E' }"
  >
    <!-- Анимированный фон -->
    <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
      <div
        v-for="i in 30"
        :key="i"
        class="floating-circle"
        :style="randomStyle()"
      />
    </div>

    <!-- Основной контейнер -->
    <div
      class="relative z-10 bg-[#2A2A40]/90 backdrop-blur-xl border border-[#6C63FF] rounded-3xl p-8 w-[520px] max-w-[95vw] shadow-[0_0_30px_#6C63FF33] animate-fade-up"
    >
      <!-- Табы форм -->
      <div class="flex mb-6 border-b border-[#3A3A50] pb-3">
        <button 
          @click="activeForm = 'login'" 
          class="flex-1 py-2 text-lg font-semibold transition-all"
          :class="activeForm === 'login' ? 'text-[#6C63FF]' : 'text-[#E4E4F0]/60 hover:text-[#E4E4F0]'"
        >
          Вход
        </button>
        <button 
          @click="activeForm = 'register'" 
          class="flex-1 py-2 text-lg font-semibold transition-all"
          :class="activeForm === 'register' ? 'text-[#6C63FF]' : 'text-[#E4E4F0]/60 hover:text-[#E4E4F0]'"
        >
          Регистрация
        </button>
        <button 
          @click="activeForm = 'recovery'" 
          class="flex-1 py-2 text-lg font-semibold transition-all"
          :class="activeForm === 'recovery' ? 'text-[#6C63FF]' : 'text-[#E4E4F0]/60 hover:text-[#E4E4F0]'"
        >
          Восстановление
        </button>
      </div>

      <!-- Заголовок, изменяющийся в зависимости от активной формы -->
      <h2
        class="text-3xl font-bold text-center text-[#E4E4F0] font-orbitron mb-8 tracking-wide animate-pulse-slow"
      >
        {{ formTitle }}
      </h2>

      <!-- Формы с анимацией переключения -->
      <transition name="fade" mode="out-in">
        <!-- Форма входа -->
        <form v-if="activeForm === 'login'" @submit.prevent="onLogin" class="space-y-5">
          <input
            v-model="loginForm.email"
            type="email"
            placeholder="Email"
            class="input"
          />
          <input
            v-model="loginForm.password"
            type="password"
            placeholder="Пароль"
            class="input"
          />

          <button type="submit" class="btn-glow w-full">
            Войти
          </button>

          <div class="mt-4 flex justify-center">
            <button
              @click="handleGoogleLogin"
              class="flex items-center bg-white text-gray-800 px-4 py-2 rounded-lg font-medium hover:bg-gray-100 transition-all"
            >
              <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24">
                <path
                  fill="#4285F4"
                  d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                />
                <path
                  fill="#34A853"
                  d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                />
                <path
                  fill="#FBBC05"
                  d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                />
                <path
                  fill="#EA4335"
                  d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                />
              </svg>
              Войти через Google
            </button>
          </div>

          <div class="text-center text-[#E4E4F0]/70 mt-4">
            <a
              @click="activeForm = 'recovery'"
              class="cursor-pointer hover:text-[#E4E4F0] transition-all"
            >
              Забыли пароль?
            </a>
          </div>
        </form>

        <!-- Форма регистрации -->
        <form v-else-if="activeForm === 'register'" @submit.prevent="onRegister" class="space-y-5">
          <input
            v-model="registerForm.name"
            type="text"
            placeholder="Имя"
            class="input"
          />
          <input
            v-model="registerForm.email"
            type="email"
            placeholder="Email"
            class="input"
          />
          <input
            v-model="registerForm.password"
            type="password"
            placeholder="Пароль"
            class="input"
          />
          <input
            v-model="registerForm.confirmPassword"
            type="password"
            placeholder="Подтвердите пароль"
            class="input"
          />

          <div class="text-[#E4E4F0] space-y-2">
            <p class="font-semibold text-sm">Выберите роль:</p>
            <div class="flex flex-wrap gap-3">
              <label class="radio-option">
                <input
                  type="radio"
                  name="role"
                  value="participant"
                  v-model="registerForm.role"
                />
                Участник
              </label>
              <label class="radio-option">
                <input
                  type="radio"
                  name="role"
                  value="organizer"
                  v-model="registerForm.role"
                />
                Организатор
              </label>
            </div>
          </div>

          <button type="submit" class="btn-glow w-full">
            Зарегистрироваться
          </button>

          <div class="mt-4 flex justify-center">
            <button
              @click="handleGoogleLogin" 
              class="flex items-center bg-white text-gray-800 px-4 py-2 rounded-lg font-medium hover:bg-gray-100 transition-all"
            >
              <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24">
                <path
                  fill="#4285F4"
                  d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
                />
                <path
                  fill="#34A853"
                  d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
                />
                <path
                  fill="#FBBC05"
                  d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
                />
                <path
                  fill="#EA4335"
                  d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
                />
              </svg>
              Регистрация через Google
            </button>
          </div>
        </form>

        <!-- Форма восстановления пароля -->
        <form v-else-if="activeForm === 'recovery'" @submit.prevent="onRecovery" class="space-y-5">
          <p class="text-[#E4E4F0]/80 text-center mb-4">
            Введите ваш email для восстановления пароля
          </p>
          <input
            v-model="recoveryForm.email"
            type="email"
            placeholder="Email"
            class="input"
          />
          <button type="submit" class="btn-glow w-full">
            Отправить ссылку
          </button>
        </form>

        <!-- Форма 2FA -->
        <form v-else-if="activeForm === '2fa'" @submit.prevent="onVerify2FA" class="space-y-5">
          <p class="text-[#E4E4F0]/80 text-center mb-4">
            Введите код из приложения аутентификации или SMS
          </p>
          <div class="flex justify-center space-x-3">
            <input
              v-for="(_, index) in 6"
              :key="index"
              v-model="twoFAForm.code[index]"
              type="text"
              maxlength="1"
              class="w-12 h-14 text-center text-xl bg-[#3A3A50] text-[#E4E4F0] rounded-lg focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50"
              @input="handleCodeInput($event, index)"
              @keydown="handleKeyDown($event, index)"
              :ref="el => { if (el) codeInputRefs[index] = el }"
            />
          </div>
          <button type="submit" class="btn-glow w-full">
            Подтвердить
          </button>
        </form>
      </transition>

      <!-- Уведомления об успехе -->
      <div
        v-if="showSuccess"
        class="fixed top-6 right-6 bg-[#6C63FF] text-[#1A1A2E] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-50"
      >
        ✅ {{ successMessage }}
      </div>

      <!-- Уведомления об ошибке -->
      <div
        v-if="showError"
        class="fixed top-6 right-6 bg-[#FF6584] text-[#1A1A2E] px-6 py-4 rounded-xl shadow-lg animate-fade-slide font-semibold z-50"
      >
        ❌ {{ errorMessage }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";

// Настройка роутера
const router = useRouter();

// Состояние переключения форм
const activeForm = ref("login");

// Заголовок формы в зависимости от активной формы
const formTitle = computed(() => {
  switch (activeForm.value) {
    case "login":
      return "Вход в систему";
    case "register":
      return "Регистрация";
    case "recovery":
      return "Восстановление пароля";
    case "2fa":
      return "Двухфакторная аутентификация";
    default:
      return "";
  }
});

// Состояние уведомлений
const showSuccess = ref(false);
const showError = ref(false);
const successMessage = ref("");
const errorMessage = ref("");

// Форма входа
const loginForm = ref({
  email: "",
  password: "",
});

// Форма регистрации
const registerForm = ref({
  name: "",
  email: "",
  password: "",
  confirmPassword: "",
  role: "participant",
});

// Форма восстановления пароля
const recoveryForm = ref({
  email: "",
});

// Форма 2FA
const twoFAForm = ref({
  code: Array(6).fill(""),
});

// Ссылки на поля ввода кода 2FA
const codeInputRefs = ref([]);

// Мок-пользователь для демонстрации
const mockUser = {
  email: "user@example.com",
  password: "password123",
  name: "Тестовый Пользователь",
};

// Обработка отправки формы входа
const onLogin = () => {
  // Простая валидация
  if (!loginForm.value.email || !loginForm.value.password) {
    errorMessage.value = "Заполните все поля";
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
    return;
  }

  // Для демо-целей, проверяем совпадение с мок-данными
  if (loginForm.value.email === mockUser.email && loginForm.value.password === mockUser.password) {
    // Имитация успеха и переход к 2FA
    successMessage.value = "Проверка данных успешна";
    showSuccess.value = true;
    setTimeout(() => {
      showSuccess.value = false;
      // Переход к 2FA
      activeForm.value = "2fa";
    }, 1500);
  } else {
    // Отображение ошибки
    errorMessage.value = "Неверный email или пароль";
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
  }
};

// Обработка отправки формы регистрации
const onRegister = () => {
  // Валидация
  if (
    !registerForm.value.name ||
    !registerForm.value.email ||
    !registerForm.value.password ||
    !registerForm.value.confirmPassword
  ) {
    errorMessage.value = "Заполните все поля";
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
    return;
  }

  if (registerForm.value.password !== registerForm.value.confirmPassword) {
    errorMessage.value = "Пароли не совпадают";
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
    return;
  }

  // Имитация успешной регистрации
  successMessage.value = "Регистрация успешна";
  showSuccess.value = true;
  setTimeout(() => {
    showSuccess.value = false;
    // Переход к входу
    activeForm.value = "login";
  }, 1500);
};

// Обработка отправки формы восстановления пароля
const onRecovery = () => {
  if (!recoveryForm.value.email) {
    errorMessage.value = "Введите email";
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
    return;
  }

  // Имитация успеха
  successMessage.value = "Ссылка для восстановления отправлена на ваш email";
  showSuccess.value = true;
  setTimeout(() => {
    showSuccess.value = false;
    // Переход к входу
    activeForm.value = "login";
  }, 3000);
};

// Обработка 2FA проверки
const onVerify2FA = () => {
  const code = twoFAForm.value.code.join("");
  
  if (code.length !== 6) {
    errorMessage.value = "Введите 6-значный код";
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
    return;
  }

  // Для демо-целей принимаем "123456" как валидный код
  if (code === "123456") {
    successMessage.value = "Аутентификация успешна";
    showSuccess.value = true;
    setTimeout(() => {
      showSuccess.value = false;
      // Переход на страницу мероприятий
      router.push('/events');
    }, 1500);
  } else {
    errorMessage.value = "Неверный код";
    showError.value = true;
    setTimeout(() => {
      showError.value = false;
    }, 3000);
  }
};

// Обработка входа через Google
const handleGoogleLogin = (e) => {
  e.preventDefault();
  
  // Имитация входа через Google
  successMessage.value = `Вход выполнен через Google как ${mockUser.name}`;
  showSuccess.value = true;
  setTimeout(() => {
    showSuccess.value = false;
    // Переход на страницу мероприятий
    router.push('/events');
  }, 2000);
};

// Обработка ввода кода для 2FA
const handleCodeInput = (event, index) => {
  // Убедимся, что вводятся только цифры
  const value = event.target.value;
  twoFAForm.value.code[index] = value.replace(/[^0-9]/g, "");
  
  // Автоматический переход к следующему полю
  if (value && index < 5) {
    codeInputRefs.value[index + 1].focus();
  }
};

// Обработка нажатий клавиш при вводе кода 2FA
const handleKeyDown = (event, index) => {
  // Обработка backspace для перехода к предыдущему полю
  if (event.key === "Backspace" && !twoFAForm.value.code[index] && index > 0) {
    codeInputRefs.value[index - 1].focus();
  }
};

// Генерация случайного стиля для анимированных кругов фона
const randomStyle = () => {
  const size = Math.floor(Math.random() * 60) + 20;
  const left = Math.random() * 100;
  const delay = Math.random() * 10;
  const duration = 10 + Math.random() * 10;
  return {
    width: `${size}px`,
    height: `${size}px`,
    left: `${left}%`,
    animationDelay: `${delay}s`,
    animationDuration: `${duration}s`,
  };
};
</script>

<style scoped>
/* Анимация плавающих кругов фона */
@keyframes floatUp {
  0% {
    transform: translateY(0) scale(1);
    opacity: 0.2;
  }
  50% {
    transform: translateY(-50vh) scale(1.2);
    opacity: 0.5;
  }
  100% {
    transform: translateY(-100vh) scale(1);
    opacity: 0;
  }
}

.floating-circle {
  position: absolute;
  background-color: rgba(108, 99, 255, 0.06);
  border-radius: 9999px;
  animation: floatUp linear infinite;
  filter: blur(1px);
}

/* Анимация форм при переключении */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateY(10px);
}

/* Стили полей ввода */
.input {
  @apply w-full p-4 rounded-lg bg-[#3A3A50] placeholder-[#E4E4F0]/70 text-[#E4E4F0] focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50 transition-all;
}

/* Стили кнопок */
.btn-glow {
  @apply p-4 rounded-lg bg-[#6C63FF] text-[#1A1A2E] font-semibold hover:bg-opacity-90 transition-all ease-in-out shadow-xl hover:shadow-[0_0_25px_rgba(108,99,255,0.5)];
}

/* Стили радио-кнопок */
.radio-option {
  @apply flex items-center bg-[#3A3A50] px-4 py-2 rounded-lg cursor-pointer hover:bg-[#4A4A60] transition-all;
}

.radio-option input {
  @apply mr-2 accent-[#6C63FF];
}

/* Анимация уведомлений */
@keyframes fadeSlide {
  0% {
    opacity: 0;
    transform: translateY(-20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-slide {
  animation: fadeSlide 0.3s ease forwards;
}

/* Анимация пульсации для заголовка */
@keyframes pulseSlow {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.8;
  }
}

.animate-pulse-slow {
  animation: pulseSlow 4s ease-in-out infinite;
}

/* Анимация появления для основного контейнера */
@keyframes fadeUp {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fade-up {
  animation: fadeUp 0.8s ease-out forwards;
}
</style> 