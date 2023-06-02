import request from "@/utils/axiosInstance.js";

export const hi = () => {
    return request.get(
        '/hi',
        {
            params: {

            }
        }
    );
}
