<template>
  <div>
    <Header />
    <ContentMore type="festival" plural-label="фестивалей" />
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import ContentMore from '@/components/ContentMore.vue'
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>