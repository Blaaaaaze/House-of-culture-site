<template>
  <div class="page">
    <Header />
    <div class="page__content">
      <ContentSingle type="news" plural-label="новостей" />
    </div>
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import ContentSingle from '@/components/ContentSingle.vue'
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
