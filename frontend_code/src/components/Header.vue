<template>
  <header class="main-header">
    <!-- Верхняя часть: логотип и кнопка слабовидящего -->
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

    <!-- Нижняя часть: навигация -->
    <div class="main-header__bottom">
      <div class="container main-section__nav-container">
        <nav class="main-section__nav nav-menu">
          <div class="hamburger" @click="toggleExtraMenu">
            <span></span>
            <span></span>
            <span></span>
          </div>
          
          <ul
            class="nav-menu__list"
            :class="{ 'nav-menu__list--open': isMenuOpen }"
          >
            <!-- Главная -->
            <router-link to="/" custom v-slot="{ navigate, href, isActive }">
              <li
                :class="['nav-menu__list-item', { active: isActive }]"
                :href="href"
                @click="navigate"
              >
                Главная
              </li>
            </router-link>

            <!-- Билеты — внешняя ссылка -->
            <li class="nav-menu__list-item">
              <a href="https://spb.kassir.ru/?utm_medium=cpc&utm_source=yandex&utm_campaign=nbo-pr14908-RF-reg-BRAND_KASSIR_SPB-handmade-search-rt%7C111328974&utm_term=kassir+официальный+сайт&utm_content=k50id%7C0100000051983206058_51983206058%7Ccid%7C111328974%7Cgid%7C5452312392%7Caid%7C16183795214%7Cadp%7Cno%7Cpos%7Cpremium2%7Csrc%7Csearch_none%7Cdvc%7Cdesktop&yclid=17582532117915762687" target="_blank" rel="noopener">Билеты</a>
            </li>

            <!-- Творческие группы -->
            <router-link to="/groups" custom v-slot="{ navigate, href, isActive }">
              <li
                :class="['nav-menu__list-item', { active: isActive }]"
                :href="href"
                @click="navigate"
              >
                Творческие группы
              </li>
            </router-link>

            <!-- Конкурсы и фестивали -->
            <router-link to="/festival" custom v-slot="{ navigate, href, isActive }">
              <li
                :class="['nav-menu__list-item', { active: isActive }]"
                :href="href"
                @click="navigate"
              >
                Конкурсы и фестивали
              </li>
            </router-link>

            <!-- Контакты -->
            <router-link to="/contacts" custom v-slot="{ navigate, href, isActive }">
              <li
                :class="['nav-menu__list-item', { active: isActive }]"
                :href="href"
                @click="navigate"
              >
                Контакты
              </li>
            </router-link>
          </ul>
          <div class="extra-menu" v-show="isExtraMenuOpen">
            <div class="footer-section__column">
              <h4>О нас</h4>
              <ul>
                <li><router-link to="/about">Основная информация</router-link></li>
                <li><router-link to="/vacancy">Вакансии</router-link></li>
                <li><router-link to="/faq">Часто задаваемые вопросы</router-link></li>
              </ul>
            </div>

            <div class="footer-section__column">
              <h4>Другое</h4>
              <ul>
                <li><a href="https://forms.yandex.ru/" target="_blank">Предложения по улучшению ГДК</a></li>
                <li><router-link to="/links">Полезные ссылки</router-link></li>
              </ul>
            </div>
          </div>
        </nav>
      </div>
    </div>
  </header>
</template>

<script setup>
import { ref } from 'vue'

const isExtraMenuOpen = ref(false)

function toggleExtraMenu() {
  isExtraMenuOpen.value = !isExtraMenuOpen.value
}
</script>




<style lang="sass" scoped>
@import '@/assets/style.sass'
</style>
