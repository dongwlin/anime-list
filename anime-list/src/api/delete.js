import request from "@/utils/axiosInstance.js";

export const deleteAnime = (id) => {
    return request.get(
        '/delete',
        {
            params: {
                id
            }
        }
    );
}
