<template>
    <div
        class="min-h-screen flex items-center justify-center relative overflow-hidden"
        :style="{ backgroundColor: '#1A1A2E' }"
    >
        <!-- 🌌 ЗВЕЗДНОЕ НЕБО -->
        <div class="absolute inset-0 z-0 pointer-events-none overflow-hidden">
            <div
                v-for="i in 30"
                :key="i"
                class="floating-circle"
                :style="randomStyle()"
            />
        </div>

        <!-- 🟢 ОСНОВНОЙ БЛОК -->
        <div
            class="relative z-10 bg-[#2A2A40]/90 backdrop-blur-xl border border-[#6C63FF] rounded-3xl p-12 w-[520px] shadow-[0_0_30px_#6C63FF33] animate-fade-up"
        >
            <h2
                class="text-4xl font-bold text-center text-[#E4E4F0] font-orbitron mb-10 tracking-wide animate-pulse-slow"
            >
                Регистрация
            </h2>

            <form @submit.prevent="onRegister" class="space-y-6">
                <transition name="fade-slide" mode="out-in">
                    <input
                        v-model="name"
                        type="text"
                        placeholder="Введите имя"
                        class="input"
                        key="name"
                    />
                </transition>

                <transition name="fade-slide" mode="out-in">
                    <input
                        v-model="surname"
                        type="text"
                        placeholder="Введите Фамилию"
                        class="input"
                        key="surname"
                    />
                </transition>

                <transition name="fade-slide" mode="out-in">
                    <input
                        v-model="email"
                        type="email"
                        placeholder="Email"
                        class="input"
                        key="email"
                    />
                </transition>

                <transition name="fade-slide" mode="out-in">
                    <input
                        v-model="password"
                        type="password"
                        placeholder="Пароль"
                        class="input"
                        key="password"
                    />
                </transition>

                <div class="text-[#E4E4F0] space-y-2">
                    <p class="font-semibold text-sm">Выберите роль:</p>
                    <div class="flex flex-wrap gap-3">
                        <label class="radio-option">
                            <input
                                type="radio"
                                name="role"
                                value="participant"
                                v-model="role"
                            />
                            Участник
                        </label>
                        <label class="radio-option">
                            <input
                                type="radio"
                                name="role"
                                value="organizer"
                                v-model="role"
                            />
                            Организатор
                        </label>
                        <label class="radio-option">
                            <input
                                type="radio"
                                name="role"
                                value="moderator"
                                v-model="role"
                            />
                            Модератор
                        </label>
                    </div>
                </div>

                <button type="submit" class="btn-glow">
                    Зарегистрироваться
                </button>

                <NuxtLink
                    to="/login"
                    class="block text-center text-[#6C63FF] hover:underline mt-4 transition-all hover:text-[#E4E4F0]"
                >
                    Уже есть аккаунт? Войти
                </NuxtLink>
            </form>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useRuntimeConfig } from "#app";

const router = useRouter();
const config = useRuntimeConfig();
const name = ref("");
const surname = ref("");
const email = ref("");
const password = ref("");
const role = ref("participant");
const csrfToken = ref("");
const errorMessage = ref("");

onMounted(async () => {
    console.log("Компонент register.vue монтируется");
    await fetchCSRFToken();
});

const fetchCSRFToken = async (attempts = 3, delay = 1000) => {
    for (let i = 0; i < attempts; i++) {
        try {
            const url = `${config.public.apiBase}/csrf-token`;
            console.log(`Запрос CSRF-токена (попытка ${i + 1}):`, url);
            const response = await $fetch(url, {
                method: "GET",
                credentials: "include",
            });
            console.log("Ответ от /csrf-token:", response);
            if (response.csrf_token) {
                csrfToken.value = response.csrf_token;
                console.log("CSRF-токен установлен:", csrfToken.value);
                errorMessage.value = "";
                return true;
            } else {
                console.error("Ответ не содержит csrf_token:", response);
                errorMessage.value = "CSRF-токен отсутствует в ответе";
            }
        } catch (error) {
            console.error(
                `Ошибка получения CSRF-токена (попытка ${i + 1}):`,
                error,
            );
            errorMessage.value = `Не удалось загрузить CSRF-токен: ${error.message}`;
            if (i < attempts - 1) {
                console.log(`Повтор через ${delay}мс...`);
                await new Promise((resolve) => setTimeout(resolve, delay));
            }
        }
    }
    console.error("Не удалось получить CSRF-токен");
    errorMessage.value = "Не удалось загрузить CSRF-токен";
    return false;
};

const onRegister = async () => {
    errorMessage.value = "";
    console.log("Попытка регистрации, текущий CSRF-токен:", csrfToken.value);

    // Проверяем токен и запрашиваем снова, если он отсутствует
    if (!csrfToken.value) {
        const success = await fetchCSRFToken();
        if (!success) {
            errorMessage.value = "CSRF-токен не загружен";
            console.error("Регистрация прервана:", errorMessage.value);
            alert(errorMessage.value);
            return;
        }
    }

    try {
        const payload = {
            name: name.value,
            surname: surname.value,
            email: email.value,
            password: password.value,
            role: role.value,
        };
        console.log("Отправка данных на /register:", payload);
        const response = await $fetch(`${config.public.apiBase}/register`, {
            method: "POST",
            credentials: "include",
            headers: {
                "X-CSRF-Token": csrfToken.value,
                "Content-Type": "application/json",
            },
            body: payload,
        });
        console.log("Регистрация успешна:", response);
        alert("Регистрация успешна!");
        router.push("/login");
    } catch (error) {
        const errMsg = error.data?.error || "Неизвестная ошибка сервера";
        console.error("Ошибка регистрации:", error);
        errorMessage.value = errMsg;
        alert(`Ошибка: ${errMsg}`);
    }
};

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
/* 🌫️ АНИМАЦИИ */
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

@keyframes fadeUp {
    from {
        opacity: 0;
        transform: translateY(40px);
    }
    to {
        opacity: 1;
        transform: translateY(0);
    }
}
.animate-fade-up {
    animation: fadeUp 0.8s ease-out forwards;
}

@keyframes pulseSlow {
    0%,
    100% {
        opacity: 1;
    }
    50% {
        opacity: 0.7;
    }
}
.animate-pulse-slow {
    animation: pulseSlow 3s infinite;
}

.floating-circle {
    position: absolute;
    background-color: rgba(108, 99, 255, 0.06);
    border-radius: 9999px;
    animation: floatUp linear infinite;
    filter: blur(1px);
}

/* ✨ ПОЛЯ */
.input {
    width: 100%;
    padding: 14px 16px;
    background-color: #1a1a2e;
    border: 1px solid #2a2a40;
    border-radius: 12px;
    color: #e4e4f0;
    font-size: 16px;
    outline: none;
    transition: all 0.3s;
}
.input::placeholder {
    color: #e4e4f088;
}
.input:focus {
    border-color: #6c63ff;
    box-shadow: 0 0 8px #6c63ff66;
}

/* 🎚️ РАДИО */
.radio-option {
    display: flex;
    align-items: center;
    gap: 8px;
    background-color: #1a1a2e;
    padding: 6px 12px;
    border-radius: 10px;
    border: 1px solid #2a2a40;
    cursor: pointer;
    transition: all 0.3s;
}
.radio-option:hover {
    border-color: #6c63ff;
    background-color: #1a1a2e60;
}
input[type="radio"] {
    accent-color: #6c63ff;
}

/* 🌟 КНОПКА */
.btn-glow {
    width: 100%;
    padding: 14px;
    background-color: #6c63ff;
    color: #1a1a2e;
    font-weight: bold;
    font-size: 16px;
    border-radius: 12px;
    transition: all 0.3s;
    box-shadow: 0 0 10px #6c63ff55;
}
.btn-glow:hover {
    background-color: #8b79ff;
    box-shadow: 0 0 20px #6c63ff99;
}

/* КНОПКА "ВЫХОД" */
.btn-exit {
    background-color: #ff6584;
    box-shadow: 0 0 10px #ff658455;
    color: #1a1a2e;
}
.btn-exit:hover {
    background-color: #ff5271;
    box-shadow: 0 0 20px #ff658499;
}

/* ПЛАВНЫЙ ПЕРЕХОД */
.fade-slide-enter-active,
.fade-slide-leave-active {
    transition: all 0.4s ease;
}
.fade-slide-enter-from {
    opacity: 0;
    transform: translateY(20px);
}
.fade-slide-leave-to {
    opacity: 0;
    transform: translateY(-20px);
}
</style>
