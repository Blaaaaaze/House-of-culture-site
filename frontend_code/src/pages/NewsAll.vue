<template>
  <div>
    <Header />
    <ContentMore type="news" plural-label="новостей" />
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import ContentMore from '@/components/ContentMore.vue'
</script>
