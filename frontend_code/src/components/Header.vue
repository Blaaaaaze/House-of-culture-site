<template>
  <header class="main-header">
    <div class="main-section__log-container">
      <header class="container main-section__header">
        <router-link to="/">
          <img class="main-section__logo" src="@/assets/img/logo.png" alt="Логотип">
        </router-link>
        <img
          id="specialButton"
          class="main-section__eye"
          style="cursor:pointer"
          src="https://lidrekon.ru/images/special.png"
          alt="Версия для слабовидящих"
          title="Версия для слабовидящих"
        />
      </header>
    </div>

    <div class="main-header__bottom">
      <div class="container main-section__nav-container">
        <nav class="main-section__nav nav-menu">
          <div class="hamburger">
            <span></span>
            <span></span>
            <span></span>
          </div>
          <ul class="nav-menu__list">
            <li class="nav-menu__list-item active"><router-link to="/">Главная</router-link></li>
            <li class="nav-menu__list-item"><router-link to="/tickets">Билеты</router-link></li>
            <li class="nav-menu__list-item"><router-link to="/groups">Творческие группы</router-link></li>
            <li class="nav-menu__list-item"><router-link to="/contests">Конкурсы и фестивали</router-link></li>
            <li class="nav-menu__list-item"><router-link to="/contacts">Контакты</router-link></li>
          </ul>
        </nav>
      </div>
    </div>
  </header>
</template>

<script setup>
</script>

<style lang="sass" scoped>
@import '@/assets/style.sass'
</style>