<template>
  <router-link :to="`/event/${event.id}`" class="event-card">
    <img
      class="event-card__image"
      :src="getImageUrl(event.preview_image)"
      alt="Постер события"
    />
    <div class="event-card__overlay">
      {{ event.name }}
    </div>
  </router-link>
</template>

<script setup>
const props = defineProps({
  event: {
    type: Object,
    required: true
  }
});

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg';
}
</script>

<style lang="sass" scoped>
@import '@/assets/style.sass'
</style>
