<template>
  <section class="event-page">
    <h1>Новость</h1>
    <p v-if="event">{{ event.title }}</p>
    <p v-else>Загрузка...</p>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const event = ref(null)

onMounted(async () => {
  try {
    const res = await fetch(`/api/event/${route.params.id}`)
    event.value = await res.json()
  } catch (e) {
    console.error('Ошибка загрузки события:', e)
  }
})
</script>

<style scoped>
</style>
