// 📁 src/router.js
import { createRouter, createWebHistory } from 'vue-router'

import Home from './pages/Home.vue'
import Groups from './pages/Groups.vue'
import Tickets from './pages/Tickets.vue'
import Contests from './pages/Contests.vue'
import Contacts from './pages/Contacts.vue'
import New from './pages/New.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/groups', component: Groups },
  { path: '/tickets', component: Tickets },
  { path: '/contests', component: Contests },
  { path: '/contacts', component: Contacts },
  { path: '/new/:id', component: New, props: true },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
