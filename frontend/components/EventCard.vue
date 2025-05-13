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
            
            <div class="flex gap-2">
                <button @click.stop="showShareOptions = !showShareOptions" class="text-accent relative">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z" />
                    </svg>
                    
                    <!-- Share popover -->
                    <div v-if="showShareOptions" 
                         class="absolute right-0 bottom-full mb-2 p-2 rounded-lg shadow-lg z-10"
                         :class="isDarkMode ? 'bg-card-dark' : 'bg-card-light'">
                        <div class="flex gap-2">
                            <button @click.stop="shareEvent('telegram')" class="p-1 rounded hover:bg-purple hover:bg-opacity-20">
                                <svg class="w-5 h-5 text-[#0088cc]" viewBox="0 0 24 24" fill="currentColor">
                                    <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm4.64 6.8c-.15 1.58-.8 5.42-1.13 7.19-.14.75-.42 1-.68 1.03-.58.05-1.02-.38-1.58-.75-.88-.58-1.38-.94-2.23-1.5-.99-.65-.35-1.01.22-1.59.15-.15 2.71-2.48 2.76-2.69.01-.03.01-.14-.07-.2-.08-.06-.19-.04-.27-.02-.12.03-1.89 1.2-5.31 3.53-.5.34-.96.51-1.37.5-.45-.02-1.32-.26-1.97-.47-.8-.26-1.43-.4-1.37-.85.03-.22.27-.45.74-.68 2.87-1.25 4.79-2.07 5.76-2.46 2.74-1.12 3.31-1.31 3.69-1.32.08 0 .27.02.39.12.1.08.13.19.14.27-.01.06.01.24 0 .38z"/>
                                </svg>
                            </button>
                            <button @click.stop="shareEvent('vk')" class="p-1 rounded hover:bg-purple hover:bg-opacity-20">
                                <svg class="w-5 h-5 text-[#4C75A3]" viewBox="0 0 24 24" fill="currentColor">
                                    <path d="M21.547 7h-3.29a.743.743 0 0 0-.655.392s-1.312 2.416-1.734 3.23c-1.43 2.85-2.053 3.377-2.35 3.233-.565-.237-.494-1.9-.494-2.826V7.183c0-.575-.213-.783-.778-.783h-3.643c-.421 0-.678.267-.678.584 0 .541.82.717.893 2.26v3.253c.038.967-.166 1.179-.557 1.179-1.012 0-3.572-3.468-5.08-7.46C3.113 6.494 2.95 6 2.357 6H.803C.136 6 0 6.416 0 6.761c0 .6.096 3.627 4.537 7.85C7.237 17.056 10.775 18 13.927 18c1.82 0 2.112-.167 2.112-.908v-2.058c0-.67.142-.804.614-.804.37 0 1.002.176 2.476 1.607C20.857 17.446 21.046 18 21.895 18h3.291c.667 0 .998-.271.808-.817-.207-.632-2.436-2.855-2.539-2.988-.347-.407-.247-.631 0-1.024.173-.291 1.16-1.72 1.798-2.733.515-.814.902-1.493 1.097-1.926.184-.402.164-.662-.45-.902"></path>
                                </svg>
                            </button>
                            <button @click.stop="shareEvent('whatsapp')" class="p-1 rounded hover:bg-purple hover:bg-opacity-20">
                                <svg class="w-5 h-5 text-[#25D366]" viewBox="0 0 24 24" fill="currentColor">
                                    <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
                                </svg>
                            </button>
                            <button @click.stop="copyEventLink" class="p-1 rounded hover:bg-purple hover:bg-opacity-20">
                                <svg class="w-5 h-5 text-gray-500" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                                    <path d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
                                </svg>
                            </button>
                        </div>
                    </div>
                </button>
                
                <button class="text-accent">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M5 10a.75.75 0 01.75-.75h8.5a.75.75 0 010 1.5h-8.5A.75.75 0 015 10z" clip-rule="evenodd" />
                    </svg>
                </button>
            </div>
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
const showShareOptions = ref(false);

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

function shareEvent(platform) {
    const eventUrl = `${window.location.origin}/event/${props.event.id}`;
    const title = props.event.title;
    const text = props.event.description || '';
    
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
    
    showShareOptions.value = false;
}

function copyEventLink() {
    const eventUrl = `${window.location.origin}/event/${props.event.id}`;
    navigator.clipboard.writeText(eventUrl).then(() => {
        showShareOptions.value = false;
        // Could add a toast notification here
    });
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