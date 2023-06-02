import {defineStore} from "pinia";
import { list } from '@/api/list.js'

const useAnimeList = defineStore('list', {
    state: () => {
        return {
            data: []
        }
    },
    actions: {
        async handleList() {
            await list()
                .then(response => {
                    this.data = response.data;
                })
                .catch(error => {
                    console.log(error);
                });
        }
    }
});

export default useAnimeList;
