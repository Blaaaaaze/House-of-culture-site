<template>
  <div class="page">
    <Header />
    <div class="page__content">
      <NewsSingle />
    </div>
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import NewsSingle from '@/components/NewsSingle.vue'
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
