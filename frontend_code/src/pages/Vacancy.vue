<template>
  <div class="page">
    <Header />
    <div class="page__content">
      <div class="vacancy-list">
        <h1 class="vacancy-list__title">Актуальные вакансии</h1>

        <div v-if="vacancies.length > 0">
          <VacancyCard
            v-for="vacancy in vacancies"
            :key="vacancy.id"
            :vacancy="vacancy"
          />
        </div>
        <p v-else>Вакансий пока нет.</p>
      </div>
      <Footer />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import VacancyCard from '@/components/VacancyCard.vue'

const vacancies = ref([])

onMounted(async () => {
  try {
    const res = await fetch('/api/vacancies')
    if (!res.ok) throw new Error('Ошибка запроса')
    vacancies.value = await res.json()
    console.log('Загруженные вакансии:', vacancies.value)
  } catch (e) {
    console.error('Ошибка загрузки вакансий:', e)
  }
})
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>

