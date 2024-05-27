
import { Typography, Box, TextareaAutosize } from "@mui/material"
import axios from "axios";

const textareaStyle = { 
    width: '100%', 
    padding: 10,  
    background: `url(http://i.imgur.com/2cOaJ.png)`,
    backgroundAttachment: 'local',
    backgroundRepeat: 'no-repeat',
    paddingLeft: 35,
    paddingTop: 10,
    borderColor: '#ccc' 
}

const Response = ({ data }) => {
    if (!data || !data.body) {
        return ( 
            <Box>
            <Typography mt={2} mb={2}>Response</Typography>
            <TextareaAutosize 
                minRows={3}
                maxRows={20}
                style={textareaStyle}
                disabled="disabled"
                value='No answer'
            />
            </Box>
        );
    }

    let body;
    try {
        body = JSON.parse(data.body);
    } catch (error) {
        console.error('Error while trying to parse JSON:', error);
        return (
            <Box>
            <Typography mt={2} mb={2}>Response</Typography>
            <TextareaAutosize 
                minRows={3}
                maxRows={20}
                style={textareaStyle}
                disabled="disabled"
                value='Произошла ошибка при обработке данных'
            />
            </Box>
        );
    }

    let readableBody = JSON.stringify(body, null, 4); // преобразуем объект в читаемый формат

    return (
        <Box>
            <Typography mt={2} mb={2}>Response</Typography>
            <TextareaAutosize 
                minRows={3}
                maxRows={20}
                style={textareaStyle}
                disabled="disabled"
                value={readableBody}
            />
        </Box>
    );
}

export default Response;