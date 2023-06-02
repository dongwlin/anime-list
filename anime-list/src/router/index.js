import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('@/layout/index/index.vue'),
        redirect: '/index',
        children: [
            {
                path: '/index',
                name: 'Index',
                component: () => import('@/views/index/index.vue')
            }
        ]
    },
    {
        path: '/setting',
        name: 'Setting',
        component: () => import('@/layout/setting/index.vue'),
        redirect: '/setting/server',
        children: [
            {
                path: '/setting/server',
                name: 'Server',
                component: () => import('@/views/setting/server/index.vue')
            },
            {
                path: '/setting/edit',
                name: 'Edit',
                component: () => import('@/views/setting/edit/index.vue')
            },
            {
                path: '/setting/about',
                name: 'About',
                component: () => import('@/views/setting/about/index.vue')
            }
        ]
    },
];

export default createRouter({
    history: createWebHashHistory(),
    routes
});

