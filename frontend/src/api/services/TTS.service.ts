import { UploadPDFRequest } from "../../models/upload_pdf_request";
import instance from "../axios";

export const uploadPDfToText = async ({ fromPage, toPage, file }: UploadPDFRequest): Promise<any> => {
    try {
        const formData = new FormData();
        formData.append('file', file);

        const response = await instance.post(
            '/tts/pdf-to-text',
            formData,
            {
                params: {
                    fromPage: fromPage,
                    toPage: toPage
                },

                headers: {
                    'Content-Type': "multipart/form-data"
                },
            },
        );
        return response.data;

    } catch (error) {
        console.error("error uploading pdf", error)
        throw error;
    }
}
