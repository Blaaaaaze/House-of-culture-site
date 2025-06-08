<template>
  <div class="new-section">
    <div class="container new-section__container">
      <!-- Одна новость -->
      <div class="news-card news-card_single" v-if="news">
        <h3 class="news-card__title news-card__title_single">{{ news.name }}</h3>
        <div class="news-card__card">
          <img
            :src="getImageUrl(news.preview_image)"
            alt="Изображение"
            class="news-slider__image news-slider__image_img"
          />
          <p class="news-card__text">{{ news.full_description }}</p>
        </div>
      </div>

      <!-- Правая колонка с недавними -->
      <div class="right-wrapper right-wrapper_single">
        <div class="recent-box">
          <h3 class="recent-title">Недавнее</h3>
          <ul class="recent-list">
            <li v-for="item in recentNews.slice(0, 6)" :key="item.id">
              <router-link :to="`/new/${item.id}`">{{ item.name }}</router-link>
            </li>
          </ul>
        </div>
        <router-link class="news-button" to="/new">Посмотреть больше новостей →</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'

const news = ref(null)
const recentNews = ref([])
const route = useRoute()

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg'
}

// Функция загрузки данных
async function loadNews(id) {
  try {
    const res = await fetch('/api/news')
    const list = await res.json()
    news.value = list.find(n => n.id === Number(id))
    recentNews.value = list.filter(n => n.id !== Number(id))
  } catch (e) {
    console.error('Ошибка загрузки новости:', e)
  }
}

// Загружаем при первом монтировании
onMounted(() => {
  loadNews(route.params.id)
})

// Перезагружаем при изменении маршрута
watch(
  () => route.params.id,
  (newId) => {
    loadNews(newId)
  }
)

</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
