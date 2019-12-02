import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import User from '@/components/User'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [
        {
            path: '/vue/helloworld/:id',
            name: 'HelloWorld',
            component: HelloWorld
        },
        {
            path: '/vue/user/:id',
            name: 'User',
            component: User
        }
    ]
})
