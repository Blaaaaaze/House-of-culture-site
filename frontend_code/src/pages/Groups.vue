<template>
  <div class="page">
    <Header />
    <div class="page__content">
      <GroupList />
    </div>
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import GroupList from '@/components/GroupList.vue';
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>