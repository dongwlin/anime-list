import { defineStore } from 'pinia';
import request from "@/utils/axiosInstance.js";

export const useAnimeOptions = defineStore('options', {
    state: () => {
        return {
            status: [
                {
                    label: 'true',
                    value: true
                },
                {
                    label: 'false',
                    value: false
                }
            ],
            type: [
                {
                    label: 'none',
                    value: -1
                },
                {
                    label: 'network',
                    value: 0
                },
                {
                    label: 'local',
                    value: 1
                }
            ],
            day: [
                {
                    label: 'none',
                    value: -1
                },
                {
                    label: 'Sunday',
                    value: 0
                },
                {
                    label: 'Monday',
                    value: 1
                },
                {
                    label: 'Tuesday',
                    value: 2
                },
                {
                    label: 'Wednesday',
                    value: 3
                },
                {
                    label: 'Thursday',
                    value: 4
                },
                {
                    label: 'Friday',
                    value: 5
                },
                {
                    label: 'Saturday',
                    value: 6
                }
            ]
        }
    },
    actions: {

    }
})

export const useAnimeList = defineStore('list', {
    state: () => {
        return {
            data: []
        }
    },
    actions: {
        async getAnimeList() {
            let tmp = [];
            await request.get(
                '/animeList',
                {
                    params: {

                    }
            }).then(function (response) {
                tmp.push(...response.data);
            }).catch(error => {
                console.log(error);
            });
            this.data = tmp;
        }
    }
})
