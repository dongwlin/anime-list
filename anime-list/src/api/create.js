import request from "@/utils/axiosInstance.js";

export const create = (formData) => {
    return request.post(
        '/create',
        formData,
        {
            headers: {
                'Content-Type': 'multipart/form-data;charset=utf-8'
            }
        }
    );
}
