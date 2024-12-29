import { Button, CircularProgress } from "@mui/material";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { uploadPDfToText } from "../../api/services/TTS.service";

export const PdfToTextForm = ({ }) => {

	const { register, handleSubmit, reset } = useForm();
	const [loading, setLoading] = useState<boolean | null>(null)
	const [error, setError] = useState<string | null>(null)

	const onSubmit = async (data: any) => {
		if (!data.file) {
			setError("Please select a file before submitting")
			console.error("no file included")
			return;
		}

		setLoading(true);
		setError(null)

		try {
			const res = await uploadPDfToText({ file: data.file[0] })
			console.log("uploaded successfully", res);
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
				<input type="file" accept="application/pdf" {...register("file")} />
				<Button type="submit">
					{loading ? <CircularProgress size={24} /> : "upload PDF"}
				</Button>
			</form>
		</>
	)
}
