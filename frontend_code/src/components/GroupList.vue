<template>
    <div class="group-list">
        <h1 class="group-list__title">Творческие группы</h1>
        <div class="container group-list__container ">
            <GroupCard
                v-for="group in groups"
                :key="group.id"
                :group="group"
            />
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import GroupCard from './GroupCard.vue'

const groups = ref([])

onMounted(async () => {
  try {
    const res = await fetch('/api/groups')
    const data = await res.json()
    groups.value = data.filter(g => g.is_active)
  } catch (e) {
    console.error('Ошибка загрузки групп:', e)
  }
})
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
.group-list
    position: relative
    background: $color-white
    overflow: hidden
    &__container
        display: grid
        grid-template-columns: repeat(2, 1fr) // ← две колонки жёстко
        gap: 24px
        padding: 10px 20px 40px 20px
        max-width: 1200px
        margin: 0 auto
    &__title
        font-size: 2rem
        margin-bottom: 1.5rem
        color: #1C5242
        font-weight: bold
        text-align: center

</style>
