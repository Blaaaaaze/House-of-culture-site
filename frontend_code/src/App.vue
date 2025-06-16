<template>
  <div id="app">
    <transition name="fade-slide" mode="out-in">
      <router-view :key="$route.fullPath" />
    </transition>
  </div>
</template>

<script setup>
// оставляем пустым, если ничего не нужно
</script>

<style lang="sass">
.fade-slide-enter-active, .fade-slide-leave-active
  transition: opacity 0.4s ease, transform 0.4s ease

.fade-slide-enter-from
  opacity: 0
  transform: translateY(30px)

.fade-slide-enter-to
  opacity: 1
  transform: translateY(0)

.fade-slide-leave-from
  opacity: 1
  transform: translateY(0)

.fade-slide-leave-to
  opacity: 0
  transform: translateY(-30px)
</style>
