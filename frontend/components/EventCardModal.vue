<template>
    <div class="bg-black bg-opacity-90 p-3 rounded-lg shadow-lg">
        <h4 class="font-orbitron">{{ event.title }}</h4>
        <p class="text-sm">{{ event.date }}</p>
        <div class="mt-1 flex gap-1">
        <span v-for="tag in event.tags" :key="tag" class="px-1 py-0.5 bg-purple bg-opacity-70 rounded-full text-xs">
            {{ tag }}
        </span>
        </div>
    </div>
</template>

<script setup>
defineProps({
    event: { type: Object, required: true }
});
</script>