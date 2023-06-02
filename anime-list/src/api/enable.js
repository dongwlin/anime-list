import request from "@/utils/axiosInstance.js";

export const enable = (id) => {
    return request.get(
        '/enable',
        {
            params: {
                id
            }
        }
    );
}
