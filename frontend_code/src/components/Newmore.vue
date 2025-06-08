<template>
    <div class="new-section">
        <div class="container new-section__container new-section__container_all">
            <div class="news-card news-card_all" v-for="news in news" :key="news.id">
                <h3 class="news-card__title news-card__title_all">{{ news.name }}</h3>
                <div class="news-card__card">
                    <img
                        :src="getImageUrl(news.preview_image)"
                        alt="Изображение"
                        class="news-slider__image active"
                    />
                    <p class="news-card__text">{{ news.short_description }}</p>
                    <router-link class="news-card__link news-card__link_all" :to="`/new/${news.id}`">Подробнее →</router-link>
                </div>
            </div>
        </div>
    </div>
</template>



<script setup>
import { ref, onMounted } from 'vue'

const news = ref([])

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg'
}

onMounted(async () => {
  try {
    const res = await fetch('/api/news')
    news.value = await res.json()
  } catch (e) {
    console.error('Ошибка загрузки новостей:', e)
  }
})
</script>


<style scoped lang="sass">
@import '@/assets/style.sass'
</style>