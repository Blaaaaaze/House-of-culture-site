// 📁 src/router.js
import { createRouter, createWebHistory } from 'vue-router'

import Home from './pages/Home.vue'
import Groups from './pages/Groups.vue'
import Contests from './pages/Contests.vue'
import Contacts from './pages/Contacts.vue'
import New from '@/pages/New.vue'
import NewsAll from '@/pages/NewsAll.vue'
import Event from '@/pages/Event.vue'
import EventAll from '@/pages/EventAll.vue'
import faq from './pages/faq.vue'
import UsefulLinks from '@/pages/UsefulLinks.vue'
import Vacancy from '@/pages/Vacancy.vue'
import VacancyApply from '@/pages/VacancyApply.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/groups', component: Groups },
  { path: '/contests', component: Contests },
  { path: '/contacts', component: Contacts },
  {
    path: '/new',
    component: NewsAll,
    props: () => ({ type: 'new', pluralLabel: 'новостей' })
  },
  {
    path: '/event',
    component: EventAll,
    props: () => ({ type: 'event', pluralLabel: 'мероприятий' })
  },

  // Одиночные (обе страницы используют ContentSingle внутри)
  {
    path: '/new/:id',
    component: New,
    props: route => ({ id: route.params.id, type: 'new', pluralLabel: 'новостей' })
  },
  {
    path: '/event/:id',
    component: Event,
    props: route => ({ id: route.params.id, type: 'event', pluralLabel: 'событий' })
  },

  { path: '/faq', component: faq},
  { path: '/links', component: UsefulLinks },
  {path: '/vacancy', component: Vacancy},
  {path: '/vacancy/:id/apply',component: VacancyApply},
]






const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router


