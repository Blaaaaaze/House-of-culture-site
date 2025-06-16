<template>
  <div class="vacancy-card">
    <h3 class="vacancy-card__title">{{ vacancy.name }}</h3>
    <p class="vacancy-card__description">{{ vacancy.long_description }}</p>
    <router-link :to="`/vacancy/${vacancy.id}/apply`" class="news-card__link news-card__link_vac">
      Откликнуться →
    </router-link>
  </div>
</template>

<script setup>
const props = defineProps({
  vacancy: {
    type: Object,
    required: true
  }
})
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
