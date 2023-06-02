import request from "@/utils/axiosInstance.js";

export const read = (id) => {
    return request.get(
        '/read',
        {
            params: {
                id
            }
        }
    );
}
