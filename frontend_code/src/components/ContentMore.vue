<template>
  <div class="new-section">
    <div class="container new-section__container new-section__container_all">
      <div
        class="news-card news-card_all"
        v-for="item in items"
        :key="item.id"
      >
        <h3 class="news-card__title news-card__title_all">{{ item.name }}</h3>
        <div class="news-card__card">
          <img
            :src="getImageUrl(item.preview_image)"
            alt="Изображение"
            class="news-slider__image active"
          />
          <p class="news-card__text">{{ item.short_description }}</p>
          <router-link class="news-card__link news-card__link_all" :to="getRoutePath(item.id)">Подробнее →</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

const props = defineProps({
  type: {
    type: String,
    required: true // 'news' или 'event'
  },
  pluralLabel: {
    type: String,
    default: 'новостей'
  }
})

const items = ref([])

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg'
}

function getRoutePath(id) {
  return props.type === 'news' ? `/new/${id}` : `/event/${id}`
}

onMounted(async () => {
  try {
    const res = await fetch(`/api/${props.type}`)
    items.value = await res.json()
  } catch (e) {
    console.error(`Ошибка загрузки ${props.pluralLabel}:`, e)
  }
})
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
