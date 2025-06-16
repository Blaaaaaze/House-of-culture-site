<template>
  <div class="news-card">
    <h3 class="news-card__title">{{ news.name }}</h3>
    <div class="news-card__card">
      <img
        :src="getImageUrl(news.preview_image)"
        alt="Изображение"
        class="news-slider__image active"
      />
      <p class="news-card__text">{{ news.short_description }}</p>
      <router-link class="news-card__link" :to="`/new/${news.id}`">Подробнее →</router-link>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  news: {
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
