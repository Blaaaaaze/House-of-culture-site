<template>
  <section class="footer-section">
    <div class="container footer-section__container">
      <div class="footer-section__column footer-section__logo-column">
        <img src="@/assets/img/logo.png" alt="Логотип" class="footer-section__logo" />
        <p>153002, Россия, г. Ломоносов, Проспект Зорина, 8</p>
        <p>Касса Пн-Пт: 12.00–19.00 Сб: 12.00–17.00<br />Вс: Выходной</p>
        <p>Касса: 7(921) 20-34-71</p>
        <p>Справочная: 7(921) 23-67-16</p>
        <p>Email: <a href="mailto:gdk-super@yandex.ru">gdk-super@yandex.ru</a></p>
      </div>

      <div class="footer-section__column">
        <h4>Основное</h4>
        <ul>
          <li><a href="#" @click.prevent="scrollToTop">Главная</a></li>
          <li class="nav-menu__list-item"><a href="https://spb.kassir.ru/?utm_medium=cpc&utm_source=yandex&utm_campaign=nbo-pr14908-RF-reg-BRAND_KASSIR_SPB-handmade-search-rt%7C111328974&utm_term=kassir+официальный+сайт&utm_content=k50id%7C0100000051983206058_51983206058%7Ccid%7C111328974%7Cgid%7C5452312392%7Caid%7C16183795214%7Cadp%7Cno%7Cpos%7Cpremium2%7Csrc%7Csearch_none%7Cdvc%7Cdesktop&yclid=17582532117915762687" target="_blank" rel="noopener">Билеты</a></li>
          <li><router-link to="/teams">Творческие группы</router-link></li>
          <li><router-link to="/fests">Конкурсы и фестивали</router-link></li>
          <li><router-link to="/contacts">Контакты</router-link></li>
        </ul>
      </div>

      <div class="footer-section__column">
        <h4>О нас</h4>
        <ul>
          <li><router-link to="/about">Основная информация</router-link></li>
          <li><router-link to="/jobs">Вакансии</router-link></li>
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
function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}
</script>


<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
