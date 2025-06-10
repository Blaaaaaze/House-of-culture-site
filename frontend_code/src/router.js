// 📁 src/router.js
import { createRouter, createWebHistory } from 'vue-router'

import Home from './pages/Home.vue'
import Groups from './pages/Groups.vue'
import Contacts from './pages/Contacts.vue'
import New from '@/pages/New.vue'
import NewsAll from '@/pages/NewsAll.vue'
import Event from '@/pages/Event.vue'
import EventAll from '@/pages/EventAll.vue'
import Festival from '@/pages/Festival.vue'
import FestivalAll from '@/pages/FestivalAll.vue'
import faq from './pages/faq.vue'
import UsefulLinks from '@/pages/UsefulLinks.vue'
import Vacancy from '@/pages/Vacancy.vue'
import VacancyApply from '@/pages/VacancyApply.vue'
import About from '@/pages/About.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/groups', component: Groups },
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
  {
  path: '/festival',
  component: FestivalAll, // или сделай FestivalAll.vue при желании
  props: () => ({ type: 'festival', pluralLabel: 'фестивалей' })
  },
  {
    path: '/festival/:id',
    component: Festival, // использует ContentSingle внутри
    props: route => ({ id: route.params.id, type: 'festival', pluralLabel: 'фестивалей' })
  },

  { path: '/faq', component: faq},
  { path: '/links', component: UsefulLinks},
  {path: '/vacancy', component: Vacancy},
  {path: '/vacancy/:id/apply',component: VacancyApply},
  {path: '/about', component: About}
]






const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      
      return { top: 0 }
    }
  }
})

export default router


