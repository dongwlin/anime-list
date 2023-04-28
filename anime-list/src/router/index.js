import { createRouter, createWebHashHistory } from "vue-router"
import AnimeList from "@/components/AnimeList.vue"
import Setting from "@/components/Setting.vue"
import SettingServer from '@/components/SettingServer.vue'
import SettingEdit from "@/components/SettingEdit.vue"

const routes = [
    {
        path: '/',
        name: 'home',
        component: AnimeList
    },
    {
        path: '/setting',
        name: 'setting',
        component: Setting,
        redirect: '/setting/server',
        children: [
            {
                path: '/setting/server',
                name: 'server',
                component: SettingServer
            },
            {
                path: '/setting/edit',
                name: 'edit',
                component: SettingEdit
            }
        ]
    }

];

export default createRouter({
    history: createWebHashHistory(),
    routes
});

