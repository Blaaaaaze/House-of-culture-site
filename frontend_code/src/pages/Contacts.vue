<template>
  <div class="page">
    <Header />
    <div class="page__content">
      <h1 class="page__title">Контакты</h1>
      <PersonList />
    </div>
    <Footer />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import PersonList from '@/components/PersonList.vue'
</script>

<style scoped lang="sass">
.page
  &__title
    font-size: 2rem
    margin-bottom: 10px
    color: #1C5242
    font-weight: bold
</style>
