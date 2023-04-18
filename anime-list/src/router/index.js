import { createRouter, createWebHashHistory } from "vue-router"
import AnimeList from "@/components/AnimeList.vue"
import Setting from "@/components/Setting.vue"

const routes = [
    {
        path: '/',
        component: AnimeList
    },
    {
        path: '/setting',
        component: Setting
    }

];

export default createRouter({
    history: createWebHashHistory(),
    routes
});

