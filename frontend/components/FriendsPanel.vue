<template>
    <transition name="slide-down">
        <div v-if="show" class="friends-panel">
        <div class="panel-header">
            <h3>Друзья</h3>
            <button @click="$emit('close')" class="close-btn">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="18" y1="6" x2="6" y2="18"></line>
                <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
            </button>
        </div>
        <div class="friends-list">
            <div v-for="(friend, index) in friends" :key="index" class="friend-item">
            <div class="friend-avatar">
                <img :src="friend.avatar" :alt="friend.name" class="avatar-img">
                <span v-if="friend.online" class="online-dot"></span>
            </div>
            <div class="friend-info">
                <p class="friend-name">{{ friend.name }}</p>
                <p class="friend-status">{{ friend.status }}</p>
            </div>
            <button class="action-btn">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="1"></circle>
                <circle cx="12" cy="5" r="1"></circle>
                <circle cx="12" cy="19" r="1"></circle>
                </svg>
            </button>
            </div>
        </div>
        </div>
    </transition>
</template>

<script setup>
defineProps({
    show: Boolean,
    friends: {
        type: Array,
        default: () => [
        {
            name: 'Алексей Иванов',
            status: 'В сети',
            online: true,
            avatar: 'https://randomuser.me/api/portraits/men/1.jpg'
        },
        {
            name: 'Мария Петрова',
            status: 'Был(а) 5 мин назад',
            online: false,
            avatar: 'https://randomuser.me/api/portraits/women/1.jpg'
        },
        {
            name: 'Дмитрий Смирнов',
            status: 'В сети',
            online: true,
            avatar: 'https://randomuser.me/api/portraits/men/2.jpg'
        },
        {
            name: 'Елена Козлова',
            status: 'Был(а) 2 часа назад',
            online: false,
            avatar: 'https://randomuser.me/api/portraits/women/2.jpg'
        }
        ]
    }
})
</script>

<style scoped>
.friends-panel {
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

.friends-list {
  padding: 8px 0;
}

.friend-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  transition: background 0.2s;
}

.friend-item:hover {
  background: rgba(255, 255, 255, 0.05);
}

.friend-avatar {
  position: relative;
  margin-right: 12px;
}

.avatar-img {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.online-dot {
  position: absolute;
  bottom: 0;
  right: 0;
  width: 10px;
  height: 10px;
  background: #4CAF50;
  border-radius: 50%;
  border: 2px solid #2A2A40;
}

.friend-info {
  flex: 1;
}

.friend-name {
  font-weight: 500;
  color: #E4E4F0;
  margin-bottom: 2px;
  font-size: 14px;
}

.friend-status {
  color: rgba(228, 228, 240, 0.7);
  font-size: 12px;
}

.action-btn {
  background: transparent;
  border: none;
  color: rgba(228, 228, 240, 0.7);
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  transition: all 0.2s;
}

.action-btn:hover {
  color: #E4E4F0;
  background: rgba(255, 255, 255, 0.1);
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

.light-theme .friends-panel {
  background: rgba(255, 255, 255, 0.95);
  border: 1px solid rgba(0, 0, 0, 0.1);
}

.light-theme .panel-header h3 {
  color: #1A1A2E;
}

.light-theme .close-btn {
  color: #1A1A2E;
}

.light-theme .friend-name {
  color: #1A1A2E;
}

.light-theme .friend-status {
  color: rgba(26, 26, 46, 0.7);
}

.light-theme .friend-item:hover {
  background: rgba(0, 0, 0, 0.03);
}

.light-theme .action-btn {
  color: rgba(26, 26, 46, 0.7);
}

.light-theme .action-btn:hover {
  color: #1A1A2E;
}

.light-theme .online-dot {
  border-color: #FFFFFF;
}
</style>
