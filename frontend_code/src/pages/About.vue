<template>
    <div class="page">
    <Header />
        <div class="page__content">
            <section class="about-section">
                <div class="container about-section__container">
                    <h1 class="about-section__title">О доме культуры</h1>

                    <h2 class="about-section__subtitle">Полное название</h2>
                    <p>Муниципальное автономное учреждение культуры "Городской Дом культуры «Огуречик»"</p>

                    <h2 class="about-section__subtitle">Краткое название</h2>
                    <p>ГДК «Огуречик»</p>

                    <h2 class="about-section__subtitle">Режим работы</h2>
                    <p>Пн-Пт: 12.00–19.00 <br>Сб: 12.00–17.00 <br>Вс: Выходной</p>

                    <h2 class="about-section__subtitle">История</h2>
                    <p>Городской Дом культуры «Огуречик» был открыт в 1978 году как центр общественной и 
                        культурной жизни для жителей города. Сначала он располагался в небольшом здании, 
                        где проводились концерты, кинопоказы и танцевальные вечера. 
                        В 1990-е годы Дом культуры получил новое современное здание, расширил творческие направления и 
                        стал базой для десятков любительских коллективов. 
                        Сегодня «Огуречик» — это динамично развивающееся учреждение, сочетающее традиции и 
                        современные формы культурной деятельности.</p>

                    <h2 class="about-section__subtitle">Вместительность зала</h2>
                    <p>
                        <ul>
                            <li>Большой зрительный зал: 500 мест</li>
                            <li>Малый зал / конференц-зал: 100 мест</li>
                        </ul>
                    </p>

                    <h2 class="about-section__subtitle">Оборудование</h2>
                    <p>
                        <ul>
                            <li>Профессиональное световое и звуковое оборудование</li>
                            <li>Сцена с выдвижным подиумом и занавесом</li>
                            <li>Проекционная техника (Full HD)</li>
                            <li>Гримерные комнаты и костюмерные</li>
                            <li>Зеркальный хореографический зал</li>
                            <li>Мастерские декоративно-прикладного творчества</li>
                            <li>Мобильная выставочная система</li>
                        </ul>
                    </p>

                    <h2 class="about-section__subtitle">Культурно-досуговые услуги</h2>
                    <p>
                        <ul>
                            <li>Организация и проведение концертов, спектаклей, выставок и фестивалей</li>
                            <li>Проведение календарных и народных праздников, торжественных мероприятий</li>
                            <li>Тематические вечера, кинопоказы, поэтические встречи</li>
                            <li>Экскурсионные программы по зданию Дома культуры и истории района</li>
                        </ul>
                    </p>

                    <h2 class="about-section__subtitle">Образовательные и творческие услуги</h2>
                    <p>
                        <ul>
                            <li>Проведение кружков и студий (вокал, хореография, театр, ИЗО, декоративно-прикладное искусство)</li>
                            <li>Мастер-классы для детей и взрослых</li>
                            <li>Индивидуальные и групповые занятия по интересам</li>
                        </ul>
                    </p>

                    <h2 class="about-section__subtitle">Аренда и техническое сопровождение</h2>
                    <p>
                        <ul>
                            <li>Аренда зрительного зала для проведения корпоративов, конференций, съёмок</li>
                            <li>Аренда залов для проведения репетиций, собраний, семинаров</li>
                            <li>Предоставление сцены с техническим обслуживанием (звук, свет, экран)</li>
                            <li>Аренда выставочных площадей</li>
                        </ul>
                    </p>

                </div>
                <div class="container about-section__container">
                    <h1 class="about-section__title">Нормативные документы</h1>
                    <div class="" v-for="doc in docs" :key="docs.id">
                        <a :href="getFileUrl(doc.file_path)" target="_blank" rel="noopener">
                            📄 {{ doc.description }}
                        </a>
                    </div>
                </div>
            </section>
        </div>

    <Footer />
    </div>
</template>

<script setup>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import { ref, onMounted } from 'vue'

const docs = ref([])

onMounted(async () => {
  try {
    const res = await fetch('/api/media/pdf?folder=uploads/docs/info&deep=true')
    docs.value = await res.json()
  } catch (e) {
    console.error('Ошибка загрузки новостей:', e)
  }
})

const getFileUrl = (path) => {
  return `http://localhost:8080/${path}`
}

</script>

<style scoped>
@import '@/assets/style.sass'
</style>