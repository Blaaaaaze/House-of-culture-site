<template>
  <div class="page">
    <div class="page__content">
      <section class="main-section">
        <Header />
        <div class="container main-section__container">
          <h1 class="main-section__title title">Новости</h1>
          <NewList />
        </div>
      </section>

      <section class="event-section">
        <div class="container event-section__container">
          <h1 class="main-section__title title"> Мероприятия</h1>
          <EventList />
        </div>
      </section>

      <section class="gos-section">
        <div class="container gos-section__container">
          <Gos />
        </div>
      </section>
    </div>
    <Footer />
  </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import NewList from '@/components/NewList.vue'
import EventList from '@/components/EventList.vue'
import Gos from '@/components/Gos.vue'
import Footer from '@/components/Footer.vue'
</script>

<style lang="sass" scoped>
@import '@/assets/style.sass'
</style>


