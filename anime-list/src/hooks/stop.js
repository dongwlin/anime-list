import {stop} from "@/api/stop.js";
import {ElMessage} from "element-plus";
import useServer from "@/store/modules/server.js";

export const useStop = async () => {
    await stop()
        .then(() => {
            useServer.status = false;
            ElMessage({
                message: 'The server is stopped.',
                type: 'success'
            });
        })
        .catch(error => {
            console.log(error);
            ElMessage({
                message: 'The server is failed to stop.',
                type: 'error'
            });
        })
}
