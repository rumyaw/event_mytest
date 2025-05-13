<template>
    <div>
        <slot />
    </div>
    </template>

    <script setup>
    useHead({
    link: [
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter&family=Orbitron&display=swap' }
    ]
    });
</script>