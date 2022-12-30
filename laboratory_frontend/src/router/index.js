import Vue from 'vue'
import Router from 'vue-router'
import mainMenu from '@/components/mainMenu'
import Login from '@/components/Login'

Vue.use(Router)

export default new Router({
    routes: [{
        path: '/',
        name: 'mainMenu',
        component: mainMenu
    }, {
        path: '/',
        name: 'Login',
        component: Login
    }]
})