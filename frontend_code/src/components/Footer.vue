<template>
  <section class="footer-section">
    <div class="container footer-section__container">
      <div class="footer-section__column footer-section__logo-column">
        <img src="@/assets/img/logo.png" alt="Логотип" class="footer-section__logo" />
        <p>153002, Россия, г. Ломоносов, Проспект Зорина, 8</p>
        <p>Касса Пн-Пт: 12.00–19.00 Сб: 12.00–17.00<br />Вс: Выходной</p>
        <p>Касса: 7(921) 20-34-71</p>
        <p>Справочная: 7(921) 23-67-17</p>
        <p>Email: <a href="mailto:gdk-super@yandex.ru">gdk-super@yandex.ru</a></p>
      </div>

    <div class="footer-section__column">
      <h4>Основное</h4>
      <ul>
        <li><a href="#" @click.prevent="handleMainClick">Главная</a></li>
        <li class="nav-menu__list-item">
          <a
            href="https://spb.kassir.ru/?..."
            target="_blank"
            rel="noopener"
          >Билеты</a>
        </li>
        <li><router-link to="/groups">Творческие группы</router-link></li>
        <li><router-link to="/festival">Конкурсы и фестивали</router-link></li>
        <li><router-link to="/contacts">Контакты</router-link></li>
      </ul>
    </div>

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
          <li>
            <a href="https://forms.yandex.ru/u/6842cc8a068ff0a44bb6f22f/" target="_blank" rel="noopener">
              Предложения по улучшению ГДК
            </a>
          </li>
          <li><router-link to="/links">Полезные ссылки</router-link></li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()

function handleMainClick() {
  if (route.path === '/') {
    window.scrollTo({ top: 0, behavior: 'smooth' })
  } else {
    router.push('/')
  }
}
</script>


<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
