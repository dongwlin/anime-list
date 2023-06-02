import request from "@/utils/axiosInstance.js";

export const update = (formData) => {
    return request.post(
        '/update',
        formData,
        {
            headers: {
                'Content-Type': 'multipart/form-data;charset=utf-8'
            }
        }
    );
}
