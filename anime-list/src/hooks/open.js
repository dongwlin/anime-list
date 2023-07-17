import { open } from "@/api/open.js";
import { ElMessage } from "element-plus";
import { config } from "@/views/setting/edit/config.js";

export const useOpenUrl = (url) => {
    window.open(url, '_blank');
}

export const useOpenDir = async (folderName) => {
    await open(folderName)
        .then(response => {
            if (response.code === 200) {
                ElMessage({
                    message: 'Open folder success.',
                    type: 'success',
                    grouping: true
                });
            } else if (response.code === 500) {
                ElMessage({
                    message: 'Open folder fail.',
                    type: 'error',
                    grouping: true
                });
            }
        })
        .catch(error => {
            console.log(error);
        });
}

export const useOpen = (type, url, dir) => {
    switch (type) {
        case config.type.none:
            break;
        case config.type.network:
            useOpenUrl(url);
            break;
        case config.type.local:
            useOpenDir(dir);
            break;
        default:
            break;
    }
}
