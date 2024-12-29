import { useEffect, useState } from "react";
import { useUser } from "../../hooks/useUser";
import { Box, Button, CircularProgress, Typography } from "@mui/material";

import { uploadPDfToText } from "../../api/services/TTS.service";
import { useForm } from "react-hook-form";

export const HomePage = () => {
    const user = useUser()
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
        } catch (error) {
            console.error("error uplaoding");
        } finally {
            setLoading(false);
        }
    }

    return (
        <>
            {user ? (
                <>
                    <Box padding={2}>
                        <form onSubmit={handleSubmit(onSubmit)}>
                            <input type="file" accept="application/pdf" {...register("file")} />
                            <Button type="submit">
                                {loading ? <CircularProgress size={24} /> : "upload PDF"}
                            </Button>
                        </form>
                    </Box>
                </>
            ) : (
                <>
                </>
            )}
        </>
    )
}
