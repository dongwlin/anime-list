import request from "@/utils/axiosInstance.js";

export const list = () => {
    return  request.get(
        '/list',
        {
            params: {}
        });
}
