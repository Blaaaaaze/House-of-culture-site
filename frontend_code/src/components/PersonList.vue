<template>
    <section class="person-section">
        <div class="person-grid">
            <PersonCard 
            v-for="person in persons" 
            :key="person.id"
            :person="person"
            />
        </div>
    </section>
</template>

<script setup>
import PersonCard from '@/components/PersonCard.vue'
import { ref, onMounted } from 'vue'


const persons = ref([])

onMounted(async () => {
  try {
    const res = await fetch('/api/contacts')
    const data = await res.json()
    persons.value = data
  } catch (error) {
    console.error('Ошибка загрузки мероприятий:', error)
  }
})
</script>

<style lang="sass" scoped>
@import '@/assets/style.sass'
</style>