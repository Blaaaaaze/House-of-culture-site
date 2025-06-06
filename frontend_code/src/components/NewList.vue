<template>
  <div class="main-section__flex-row">
    <div class="left-wrapper">
      <NewsCard
        v-for="news in news"
        :key="news.id"
        :news="news"
      />
    </div>

    <div class="right-wrapper">
      <div class="recent-box">
        <h3 class="recent-title">Недавнее</h3>
        <ul class="recent-list">
          <li v-for="news in news.slice(0, 6)" :key="news.id">
            <router-link :to="`/new/${news.id}`">{{ news.name }}</router-link>
          </li>
        </ul>
      </div>
      <router-link class="news-button" to="/new">Больше новостей →</router-link>
    </div>

  </div>
</template>


<script setup>
import { ref, onMounted } from 'vue'
import NewsCard from './NewsCard.vue'

const news = ref([])

onMounted(async () => {
  try {
    const res = await fetch('/api/news')
    news.value = await res.json()
  } catch (e) {
    console.error('Ошибка загрузки новостей:', e)
  }
})
</script>

<style lang="sass" scoped>
@import '@/assets/style.sass'
</style>