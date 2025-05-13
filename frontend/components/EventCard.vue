<template>
    <div :id="`event-${event.id}`"
         :class="['p-4 rounded-lg shadow-lg hover:shadow-purple transition cursor-pointer', 
                 isDarkMode ? 'bg-card-dark bg-opacity-70' : 'bg-card-light bg-opacity-90']"
         @click="$emit('click', event)">
        <div class="flex justify-between items-start">
            <h3 class="text-lg font-orbitron font-semibold">{{ event.title }}</h3>
            <div class="text-xs px-2 py-1 rounded-full" 
                :class="categoryClass">
                {{ event.category || 'Без категории' }}
            </div>
        </div>
        
        <p class="text-sm mt-2 opacity-80">{{ formatDate(event.date) || 'Дата неизвестна' }}</p>
        
        <p v-if="event.description" class="text-sm mt-2 line-clamp-2">{{ event.description }}</p>
        
        <div class="mt-3 flex justify-between items-center">
            <div class="flex gap-2 flex-wrap">
                <span v-for="tag in event.tags" :key="tag" 
                    :class="['px-2 py-1 rounded-full text-xs', 
                            isDarkMode ? 'bg-purple bg-opacity-50' : 'bg-purple bg-opacity-30']">
                    {{ tag }}
                </span>
            </div>
            
            <button class="text-accent">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                    <path fill-rule="evenodd" d="M5 10a.75.75 0 01.75-.75h8.5a.75.75 0 010 1.5h-8.5A.75.75 0 015 10z" clip-rule="evenodd" />
                </svg>
            </button>
        </div>
    </div>
</template>

<script setup>
const props = defineProps({
    event: { type: Object, required: true }
});

defineEmits(['click']);

const theme = useState('theme');
const isDarkMode = computed(() => theme.value === 'dark');

const categoryClass = computed(() => {
    const categories = {
        'tech': 'bg-blue-500 bg-opacity-70 text-white',
        'music': 'bg-purple-500 bg-opacity-70 text-white',
        'art': 'bg-pink-500 bg-opacity-70 text-white',
        'default': isDarkMode.value ? 'bg-gray-700 text-gray-200' : 'bg-gray-300 text-gray-700'
    };
    
    return categories[props.event.category] || categories.default;
});

function formatDate(dateString) {
    if (!dateString) return null;
    
    try {
        const date = new Date(dateString);
        return date.toLocaleDateString('ru-RU', {
            day: 'numeric',
            month: 'long',
            year: 'numeric',
            hour: '2-digit', 
            minute: '2-digit'
        });
    } catch {
        return dateString;
    }
}
</script>

<style scoped>
.bg-card-dark {
    background-color: var(--card-dark);
}
.bg-card-light {
    background-color: var(--card-light);
}
</style>