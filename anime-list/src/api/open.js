import request from "@/utils/axiosInstance.js";

export const open = (folderName) => {
    return request.get(
        '/open',
        {
            params: {
                folder: folderName
            }
        }
    );
}
