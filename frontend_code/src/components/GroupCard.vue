<template>
  <div class="group-card">
    <img
      :src="getImageUrl(group.preview_image)"
      alt="Изображение группы"
      class="group-card__image"
    />
    <h3 class="group-card__title">{{ group.name }}</h3>
    <p class="group-card__description">{{ group.short_description }}</p>
    <router-link :to="`/group/${group.id}`" class="news-card__link">
      Подробнее →
    </router-link>
  </div>
</template>

<script setup>
const props = defineProps({
  group: Object
})

function getImageUrl(path) {
  return path ? `http://localhost:8080/${path}` : '/icons/placeholder.jpg'
}
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
.group-card
  background: #fff
  border-radius: 12px
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08)
  padding: 16px
  display: flex
  flex-direction: column
  transition: transform 0.2s ease

  &:hover
    transform: translateY(-4px)

  &__image
    width: 100%
    height: 180px
    object-fit: cover
    border-radius: 8px
    margin-bottom: 12px

  &__title
    font-size: 20px
    font-weight: 600
    margin-bottom: 8px

  &__description
    flex-grow: 1
    font-size: 14px
    color: #555
    margin-bottom: 12px

  &__link
    align-self: flex-start
    font-weight: 500
    color: #2e7bcf
    text-decoration: none

    &:hover
      text-decoration: underline
</style>
