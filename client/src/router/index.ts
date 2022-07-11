import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Index from '../pages/Index.vue'
import Login from '../pages/Login.vue'
import Register from '../pages/Register.vue'
import Forgot from '../pages/Forgot.vue'
import Reset from '../pages/Reset.vue'
import AccountProfile from '../pages/AccountProfile.vue'
import AccountSetting from '../pages/AccountSetting.vue'
import TimeLine from '../pages/timeLine.vue'
import Chat from '../pages/Chat.vue'

const routes: Array<RouteRecordRaw> = [
    {
        path: '/',
        name: 'index',
        component: Index,
    },
    {
        path: '/login',
        name: 'login',
        component: Login,
    },
    {
        path: '/register',
        name: 'register',
        component: Register,
    },
    {
        path: '/forgot',
        name: 'forgot',
        component: Forgot,
    },
    {
        path: '/reset/:token',
        name: 'reset',
        component: Reset,
    },
    {
        path: '/account/profile',
        name: 'account_profile',
        component: AccountProfile,
    },
    {
        path: '/account/setting',
        name: 'account_setting',
        component: AccountSetting,
    },
    {
        path: '/home',
        name: 'home',
        component: TimeLine,
    },
    {
        path: '/chat',
        name: 'chat',
        component: Chat,
    }

]

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
})

export default router
