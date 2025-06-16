<template>
    <div class="page">
        <Header />
        <div class="page__content">
            <div class="links-page">
                <div class="container links-page__container">
                    <h2 class="links-page__title links-page__title_up">Ссылки</h2>

                    <section class="links-page__section">
                    <h2 class="links-page__title">Противодействие коррупции</h2>
                    <div class="links-page__grid">
                    <a
                        href="https://genproc.gov.ru/"
                        target="_blank"
                        rel="noopener"
                    >
                        <img src="@/assets/img/links/genproc.jpg" alt="Генеральная прокуратура" />
                    </a>
                    <a
                        href="https://prokuratura.ivanovo.ru/"
                        target="_blank"
                        rel="noopener"
                    >
                        <img src="@/assets/img/links/ivanovo-proc.jpg" alt="Прокуратура Ивановской области" />
                    </a>
                    </div>
                </section>

                <section class="links-page__section">
                    <h2 class="links-page__title">Полезные ссылки</h2>
                    <div class="links-page__grid">
                    <a href="https://www.gosuslugi.ru/" target="_blank" rel="noopener">
                        <img src="@/assets/img/links/gosuslugi.jpg" alt="Госуслуги" />
                    </a>
                    <a href="https://culture.ivanovoobl.ru/" target="_blank" rel="noopener">
                        <img src="@/assets/img/links/departament.jpg" alt="Департамент культуры" />
                    </a>
                    <a href="https://culture.gov.ru/" target="_blank" rel="noopener">
                        <img src="@/assets/img/links/mincult.jpeg" alt="Минкульт" />
                    </a>
                    <a href="https://ivgoradm.ru/" target="_blank" rel="noopener">
                        <img src="@/assets/img/links/adm-ivanovo.jpeg" alt="Администрация Иваново" />
                    </a>
                    <a href="https://www.culture.ru/pushkinskaya-karta" target="_blank" rel="noopener">
                        <img src="@/assets/img/links/pushkinskaya.jpg" alt="Пушкинская карта" />
                    </a>
                    <a href="https://www.culture.ru/school-diary" target="_blank" rel="noopener">
                        <img src="@/assets/img/links/school-diary.jpg" alt="Культурный дневник" />
                    </a>
                    </div>
                </section>
                </div>
            </div>
        </div>
        <Footer />
    </div>
</template>

<script setup>
    import Header from '@/components/Header.vue'
    import Footer from '@/components/Footer.vue'
</script>

<style scoped lang="sass">
@import '@/assets/style.sass'
</style>
