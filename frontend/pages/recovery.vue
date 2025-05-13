<template>
    <div
        class="min-h-screen flex items-center justify-center relative overflow-hidden"
        :style="{ backgroundColor: '#1A1A2E' }"
    >
        <!-- Анимированный фон -->
        <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
            <div
                v-for="i in 40"
                :key="i"
                class="floating-circle"
                :style="randomStyle()"
            />
        </div>

        <!-- Основной блок -->
        <div
            class="relative z-10 bg-[#2A2A40]/90 backdrop-blur-xl border border-[#6C63FF] rounded-3xl p-10 w-[480px] max-w-[95vw] shadow-[0_0_30px_#6C63FF33] animate-fade-up"
        >
            <h2
                class="text-3xl font-bold text-center text-[#E4E4F0] font-orbitron mb-8 tracking-wide animate-pulse-slow"
            >
                Восстановление пароля
            </h2>
            
            <div v-if="step === 'email'" class="space-y-6">
                <p class="text-[#E4E4F0]/80 text-center mb-4">
                    Введите ваш email для получения ссылки на восстановление пароля
                </p>
                
                <form @submit.prevent="sendResetLink" class="space-y-5">
                    <input
                        v-model="email"
                        type="email"
                        placeholder="Email"
                        class="w-full p-4 rounded-lg bg-[#3A3A50] placeholder-[#E4E4F0]/70 text-[#E4E4F0] focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50 transition-all"
                    />
                    <button
                        type="submit"
                        class="w-full p-4 rounded-lg bg-[#6C63FF] text-[#1A1A2E] font-semibold hover:bg-opacity-90 transition-all ease-in-out shadow-xl hover:shadow-[0_0_25px_rgba(108,99,255,0.5)]"
                    >
                        Отправить ссылку
                    </button>
                    
                    <NuxtLink
                        to="/login"
                        class="block text-center text-[#6C63FF] hover:underline mt-4 transition-all hover:text-[#E4E4F0]"
                    >
                        Вернуться к входу
                    </NuxtLink>
                </form>
            </div>
            
            <div v-else-if="step === 'code'" class="space-y-6">
                <p class="text-[#E4E4F0]/80 text-center mb-4">
                    Введите код подтверждения, отправленный на ваш email
                </p>
                
                <form @submit.prevent="verifyCode" class="space-y-5">
                    <div class="flex justify-center space-x-3">
                        <input
                            v-for="(_, index) in 6"
                            :key="index"
                            v-model="verificationCode[index]"
                            type="text"
                            maxlength="1"
                            class="w-12 h-14 text-center text-xl bg-[#3A3A50] text-[#E4E4F0] rounded-lg focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50"
                            @input="handleCodeInput($event, index)"
                            @keydown="handleKeyDown($event, index)"
                            :ref="el => { if (el) codeInputRefs[index] = el }"
                        />
                    </div>
                    <button
                        type="submit"
                        class="w-full p-4 rounded-lg bg-[#6C63FF] text-[#1A1A2E] font-semibold hover:bg-opacity-90 transition-all ease-in-out shadow-xl hover:shadow-[0_0_25px_rgba(108,99,255,0.5)]"
                    >
                        Подтвердить код
                    </button>
                    
                    <div class="text-center">
                        <button
                            @click.prevent="sendResetLink"
                            class="text-[#6C63FF] hover:underline transition-all hover:text-[#E4E4F0]"
                        >
                            Отправить код повторно
                        </button>
                    </div>
                </form>
            </div>
            
            <div v-else-if="step === 'reset'" class="space-y-6">
                <p class="text-[#E4E4F0]/80 text-center mb-4">
                    Создайте новый пароль
                </p>
                
                <form @submit.prevent="resetPassword" class="space-y-5">
                    <input
                        v-model="newPassword"
                        type="password"
                        placeholder="Новый пароль"
                        class="w-full p-4 rounded-lg bg-[#3A3A50] placeholder-[#E4E4F0]/70 text-[#E4E4F0] focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50 transition-all"
                    />
                    <input
                        v-model="confirmPassword"
                        type="password"
                        placeholder="Подтвердите пароль"
                        class="w-full p-4 rounded-lg bg-[#3A3A50] placeholder-[#E4E4F0]/70 text-[#E4E4F0] focus:outline-none focus:ring-2 focus:ring-[#6C63FF]/50 transition-all"
                    />
                    <button
                        type="submit"
                        class="w-full p-4 rounded-lg bg-[#6C63FF] text-[#1A1A2E] font-semibold hover:bg-opacity-90 transition-all ease-in-out shadow-xl hover:shadow-[0_0_25px_rgba(108,99,255,0.5)]"
                    >
                        Сохранить новый пароль
                    </button>
                </form>
            </div>
        </div>

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
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

// Состояние формы
const step = ref("email"); // Шаги: email -> code -> reset
const email = ref("");
const verificationCode = ref(Array(6).fill(""));
const newPassword = ref("");
const confirmPassword = ref("");

// Уведомления
const showSuccess = ref(false);
const showError = ref(false);
const successMessage = ref("");
const errorMessage = ref("");

// Ссылки на поля ввода кода
const codeInputRefs = ref([]);

// Отправка ссылки восстановления
const sendResetLink = () => {
    if (!email.value) {
        errorMessage.value = "Введите email";
        showError.value = true;
        setTimeout(() => {
            showError.value = false;
        }, 3000);
        return;
    }

    // Имитация отправки ссылки
    successMessage.value = "Код подтверждения отправлен на ваш email";
    showSuccess.value = true;
    setTimeout(() => {
        showSuccess.value = false;
        step.value = "code"; // Переход к вводу кода
    }, 2000);
};

// Проверка кода подтверждения
const verifyCode = () => {
    const code = verificationCode.value.join("");
    
    if (code.length !== 6) {
        errorMessage.value = "Введите 6-значный код";
        showError.value = true;
        setTimeout(() => {
            showError.value = false;
        }, 3000);
        return;
    }

    // Для демо любой код 123456 считается правильным
    if (code === "123456") {
        successMessage.value = "Код подтвержден";
        showSuccess.value = true;
        setTimeout(() => {
            showSuccess.value = false;
            step.value = "reset"; // Переход к сбросу пароля
        }, 1500);
    } else {
        errorMessage.value = "Неверный код";
        showError.value = true;
        setTimeout(() => {
            showError.value = false;
        }, 3000);
    }
};

// Сброс пароля
const resetPassword = () => {
    if (!newPassword.value || !confirmPassword.value) {
        errorMessage.value = "Заполните оба поля";
        showError.value = true;
        setTimeout(() => {
            showError.value = false;
        }, 3000);
        return;
    }

    if (newPassword.value !== confirmPassword.value) {
        errorMessage.value = "Пароли не совпадают";
        showError.value = true;
        setTimeout(() => {
            showError.value = false;
        }, 3000);
        return;
    }

    // Имитация сброса пароля
    successMessage.value = "Пароль успешно изменен";
    showSuccess.value = true;
    setTimeout(() => {
        showSuccess.value = false;
        router.push("/login");
    }, 2000);
};

// Обработка ввода кода
const handleCodeInput = (event, index) => {
    // Убедимся, что вводятся только цифры
    const value = event.target.value;
    verificationCode.value[index] = value.replace(/[^0-9]/g, "");
    
    // Автоматический переход к следующему полю
    if (value && index < 5) {
        codeInputRefs.value[index + 1].focus();
    }
};

// Обработка нажатий клавиш при вводе кода
const handleKeyDown = (event, index) => {
    // Обработка backspace для перехода к предыдущему полю
    if (event.key === "Backspace" && !verificationCode.value[index] && index > 0) {
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