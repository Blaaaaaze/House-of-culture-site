<template>
  <section class="events-section">

    <div class="events-grid">
      <EventCard
        v-for="event in events.slice(0, 3)"
        :key="event.id"
        :event="event"
      />
    </div>

    <router-link class="event-button" to="/event">Больше мероприятий →</router-link>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import EventCard from './EventCard.vue'

const events = ref([])

onMounted(async () => {
  try {
    const res = await fetch('/api/event')
    const data = await res.json()
    events.value = data.filter(e => e.type === 'event' && e.is_active)
  } catch (error) {
    console.error('Ошибка загрузки мероприятий:', error)
  }
})
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
