import request from "@/utils/axiosInstance.js";

export const stop = () => {
    return request.get(
        '/stop',
        {
            params: {

            }
        }
    );
}
