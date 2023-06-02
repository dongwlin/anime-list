import request from "@/utils/axiosInstance.js";

export const disable = (id) => {
    return request.get(
        '/disable',
        {
            params: {
                id
            }
        }
    );
}
