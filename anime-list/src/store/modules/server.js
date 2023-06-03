import {defineStore} from "pinia";
import { hi } from '@/api/hi.js';
import {about} from "@/api/about.js";

const useServer = defineStore('server', {
    state: () => {
        return {
            status: true,
            rootDir: '',
            versionCode: 0,
            versionName: 'NONE'
        }
    },
    actions: {
        async handleHi() {
            await hi()
                .then(() => {
                    this.status = true;
                })
                .catch(error => {
                    console.log(error);
                    this.status = false;
                })
        },
        async handleAbout() {
            await about()
                .then(res => {
                    this.rootDir = res.data.rootDir;
                    this.versionCode = res.data.versionCode;
                    this.versionName = res.data.versionName;
                })
                .catch(error => {
                    console.log(error);
                });
        }
    }
});

export default useServer;
