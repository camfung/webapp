
import { Button, CircularProgress } from "@mui/material";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { textToSpeach } from "../../api/services/TTS.service";

export const TtsForm = ({ }) => {

	const { register, handleSubmit, reset } = useForm();
	const [loading, setLoading] = useState<boolean | null>(null)
	const [error, setError] = useState<string | null>(null)


	const saveMp3File = (blob: Blob, fileName: string) => {
		const url = URL.createObjectURL(blob);
		const link = document.createElement("a");
		link.href = url;
		link.download = fileName;
		link.click();
		URL.revokeObjectURL(url)
	}

	const onSubmit = async (data: any) => {
		if (data.text == "") {
			console.error("must include text")
			return
		}
		setLoading(true);
		setError(null)

		try {
			const blob = await textToSpeach({ text: data.text, api: "openai", voice: "alloy" })
			saveMp3File(blob, "test.mp3")

			console.log(data)
			reset();
		} catch (err) {
			console.error("error uplaoding", error);
		} finally {
			setLoading(false);
		}
	}
	return (
		<>

			<form onSubmit={handleSubmit(onSubmit)}>
				<textarea name="" id="" {...register("text")}></textarea>
				<Button type="submit">
					{loading ? <CircularProgress size={24} /> : "upload PDF"}
				</Button>
			</form>
		</>
	)
}
