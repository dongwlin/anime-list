import request from "@/utils/axiosInstance.js";

export const about = () => {
    return request.get(
        '/about',
        {
            params: {}
        }
    );
}
