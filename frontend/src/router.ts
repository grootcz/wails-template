import { createRouter, createWebHashHistory, RouteRecordRaw } from 'vue-router'

import Home from './components/Home.vue'
import HomeMenu1 from './components/HomeMenu1.vue'
import HomeMenu2 from './components/HomeMenu2.vue'
import HomeMenu3 from './components/HomeMenu3.vue'
import HomeMenu4 from './components/HomeMenu4.vue'
import Features from './components/Features.vue'

const routes: RouteRecordRaw[] = [
  {path: '/', component: Home},
  {
    path: '/home', component: Home,
    children: [
      {path: 'menu1', component: HomeMenu1},
      {path: 'menu2', component: HomeMenu2},
      {path: 'menu3', component: HomeMenu3},
      {path: 'menu4', component: HomeMenu4},
    ],
  },
  {path: '/features', component: Features},
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})

router.beforeEach((to, from, next) => {
  console.log("to=", to)
  //console.log("from=", from)

  next()
})

export default router
