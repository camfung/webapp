import { useUser } from "../../hooks/useUser";
import { Box, Button, CircularProgress, Typography } from "@mui/material";

import { PdfToTextForm } from "../../components/forms/PdfToTextForm";
import { TtsForm } from "../../components/forms/TtsForm";

export const HomePage = () => {
    const user = useUser()


    return (
        <>
            {user ? (
                <>
                    <Box padding={2}>
                        <PdfToTextForm />
                        <TtsForm />
                    </Box>
                </>
            ) : (
                <>
                </>
            )}
        </>
    )
}
