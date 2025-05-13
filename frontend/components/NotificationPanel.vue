<template>
    <transition name="slide-down">
        <div v-if="show" class="notification-panel">
        <div class="panel-header">
            <h3>Уведомления</h3>
            <button @click="$emit('close')" class="close-btn">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="18" y1="6" x2="6" y2="18"></line>
                <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
            </button>
        </div>
        <div class="notification-list">
            <div v-for="(notification, index) in notifications" :key="index" class="notification-item">
            <div class="notification-icon">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="12" y1="8" x2="12" y2="12"></line>
                <line x1="12" y1="16" x2="12.01" y2="16"></line>
                </svg>
            </div>
            <div class="notification-content">
                <p class="notification-title">{{ notification.title }}</p>
                <p class="notification-message">{{ notification.message }}</p>
                <span class="notification-time">{{ notification.time }}</span>
            </div>
            <div v-if="notification.unread" class="unread-dot"></div>
            </div>
        </div>
        </div>
    </transition>
</template>

<script setup>
defineProps({
  show: Boolean,
  notifications: {
    type: Array,
    default: () => [
      {
        title: 'Новое сообщение',
        message: 'У вас новое сообщение от Алексея',
        time: '2 мин назад',
        unread: true
      },
      {
        title: 'Напоминание',
        message: 'Завтра мероприятие в 14:00',
        time: '1 час назад',
        unread: false
      },
      {
        title: 'Обновление',
        message: 'Доступно новое обновление приложения',
        time: '3 часа назад',
        unread: true
      }
    ]
  }
})
</script>

<style scoped>
.notification-panel {
  position: absolute;
  top: 100%;
  right: 0;
  width: 360px;
  max-height: 500px;
  overflow-y: auto;
  background: rgba(42, 42, 64, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 14px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.1);
  margin-top: 10px;
  z-index: 1000;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.panel-header h3 {
  font-size: 18px;
  font-weight: 600;
  color: #E4E4F0;
}

.close-btn {
  background: transparent;
  border: none;
  color: #E4E4F0;
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  transition: background 0.2s;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.notification-list {
  padding: 8px 0;
}

.notification-item {
  display: flex;
  padding: 12px 16px;
  position: relative;
  transition: background 0.2s;
}

.notification-item:hover {
  background: rgba(255, 255, 255, 0.05);
}

.notification-icon {
  margin-right: 12px;
  color: #6C63FF;
}

.notification-content {
  flex: 1;
}

.notification-title {
  font-weight: 500;
  color: #E4E4F0;
  margin-bottom: 4px;
  font-size: 14px;
}

.notification-message {
  color: rgba(228, 228, 240, 0.7);
  font-size: 13px;
  margin-bottom: 4px;
}

.notification-time {
  color: rgba(228, 228, 240, 0.5);
  font-size: 12px;
}

.unread-dot {
  position: absolute;
  right: 16px;
  top: 16px;
  width: 8px;
  height: 8px;
  background: #6C63FF;
  border-radius: 50%;
}

.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.3s ease;
}

.slide-down-enter-from,
.slide-down-leave-to {
  transform: translateY(-10px);
  opacity: 0;
}

.light-theme .notification-panel {
  background: rgba(255, 255, 255, 0.95);
  border: 1px solid rgba(0, 0, 0, 0.1);
}

.light-theme .panel-header h3 {
  color: #1A1A2E;
}

.light-theme .close-btn {
  color: #1A1A2E;
}

.light-theme .notification-title {
  color: #1A1A2E;
}

.light-theme .notification-message {
  color: rgba(26, 26, 46, 0.7);
}

.light-theme .notification-time {
  color: rgba(26, 26, 46, 0.5);
}

.light-theme .notification-item:hover {
  background: rgba(0, 0, 0, 0.03);
}
</style>
