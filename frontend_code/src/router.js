// 📁 src/router.js
import { createRouter, createWebHistory } from 'vue-router'

import Home from './pages/Home.vue'
import Groups from './pages/Groups.vue'
import Contests from './pages/Contests.vue'
import Contacts from './pages/Contacts.vue'
import New from '@/pages/New.vue'
import NewsAll from '@/pages/Newsall.vue'
import faq from './pages/faq.vue'
import UsefulLinks from '@/pages/UsefulLinks.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/groups', component: Groups },
  { path: '/contests', component: Contests },
  { path: '/contacts', component: Contacts },
  {path: '/new/:id', component: New},
  {path: '/new', component: NewsAll},
  { path: '/faq', component: faq},
  { path: '/links', component: UsefulLinks },
]






const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router


