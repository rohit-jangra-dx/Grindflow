import {createMemoryHistory, createRouter, RouteRecordRaw} from 'vue-router'


import DashBoard from '../pages/DashBoard.vue'
import Schedule from '../pages/Schedule.vue'
import Plan from '../pages/Plan.vue'
import Boards from '../pages/Boards.vue'

const routes: RouteRecordRaw[] = [
    { path: '/', component: DashBoard},
    { path: '/plan', component: Plan},
    { path: '/schedule', component: Schedule},
    { path: '/boards', component: Boards},

]

export const router = createRouter({
    history: createMemoryHistory(),
    routes
})
