import axios from "axios";
import { TTSRequest } from "../../models/tts_request";
import { UploadPDFRequest } from "../../models/upload_pdf_request";
import instance from "../axios";

export const uploadPDfToText = async ({ fromPage, toPage, file }: UploadPDFRequest): Promise<string> => {
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
        return JSON.parse(response.data).text;

    } catch (error) {
        console.error("error uploading pdf", error)
        throw error;
    }
}

export const textToSpeach = async ({ text, api, voice }: TTSRequest): Promise<Blob> => {
    try {

        const response = await axios.post(
            "/tts/text-to-speach",
            { text: text, api: api, voice: voice },
            { responseType: "blob" }
        )
        return response.data

    } catch (error) {
        console.error("error in request texttospeach service")
        throw error
    }

}
