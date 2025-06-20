<template>
  <div class="new-section">
    <div class="container new-section__container">
      <!-- Основной блок -->
      <div class="news-card news-card_single" v-if="item">
        <h3 class="news-card__title news-card__title_single">{{ item.name }}</h3>
        <div class="news-card__card">
          <img
            :src="getImageUrl(getImageMedia(item)[0].file_path)"
            alt="Изображение"
            class="news-slider__image news-slider__image_img"
          />
          <p class="news-card__text" v-html="item.full_description"></p>
          <!-- Только для festival -->
          <div v-if="props.type === 'festival' && getPdfMedia(item)" class="festival-pdf">
            <a
              :href="`http://localhost:8080/${getPdfMedia(item).file_path}`" 
              target="_blank"
              rel="noopener noreferrer"
            >
              📄 Открыть документ фестиваля
            </a>
          </div>

          <!-- Только для event -->
          <div v-else-if="props.type === 'event'" class="event-ticket-button">
            <a
              href="https://www.kassir.ru"
              target="_blank"
              rel="noopener noreferrer"
              class="news-card__link news-card__link_all"
            >
              Купить билет →
            </a>
          </div>

        </div>
      </div>

      <!-- Боковая колонка -->
      <div class="right-wrapper right-wrapper_single">
        <div class="recent-box">
          <h3 class="recent-title">Недавнее</h3>
          <ul class="recent-list">
            <li v-for="recent in recentList.slice(0, 6)" :key="recent.id">
              <router-link :to="getRoutePath(recent.id)">{{ recent.name }}</router-link>
            </li>
          </ul>
        </div>
        <router-link class="news-button" :to="type === 'news' ? '/new' : '/event'">
          Посмотреть больше {{ pluralLabel }} →
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watchEffect } from 'vue'
import { useRoute } from 'vue-router'

const props = defineProps({
  type: String,
  pluralLabel: String
})

const item = ref(null)
const recentList = ref([])
const route = useRoute()

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg'
}

function getRoutePath(id) {
  if (props.type === 'news') return `/new/${id}`
  if (props.type === 'event') return `/event/${id}`
  if (props.type === 'festival') return `/festival/${id}`
  return `/` // на всякий случай, если тип не распознан
}


async function loadData(id) {
  try {
    const res = await fetch(`/api/${props.type}`)
    const list = await res.json()
    item.value = list.find(el => el.id === Number(id))
    recentList.value = list.filter(el => el.id !== Number(id))

    item.value = {
      ...list.find(el => el.id === Number(id)), // старый формат
      media: [], // добавим поле, чтобы не было undefined
    }

    const detailedRes = await fetch(`/api/${props.type}/${id}`) // получаем полные данные
    const detailedItem = await detailedRes.json()

    item.value = {
      ...detailedItem.content,
      media: detailedItem.media,
    }
    console.log(item.value)
  } catch (e) {
    console.error(`Ошибка загрузки ${props.type}:`, e)
  }
}

watchEffect(() => {
  if (route.params.id) {
    loadData(route.params.id)
  }
})

function getPdfMedia(data) {
  return data.media?.find(item => item.type === 'pdf')
}

function getImageMedia(data) {
  return data.media?.filter(item => item.type === 'image')
}

</script>


<style scoped lang="sass">
@import '@/assets/style.sass'
</style>