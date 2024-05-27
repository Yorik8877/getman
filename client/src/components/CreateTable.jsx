import { useState } from 'react';

import { Box, Typography, Table, TableHead, TableBody, TableCell, TableRow, Button } from '@mui/material';
import { makeStyles } from '@mui/styles';

import ChangeRow from './ChangeRow';

const useStyles = makeStyles({
    textfield: {
        width: '100%'
    },
    tablecell: {
        padding: ['7px 5px', '!important'],
        border: '1px solid rgba(224, 224, 224, 1)',
        borderCollapse: 'collapse'
    },
    checkbox: {
        padding: ['2px 5px', '!important'],
    },
    button: {
        margin: ['10px 0px', '!important']
    }
});

const CreateTable = ({ text, data, setData }) => {
    const classes = useStyles();
    const [rows, changeRows] = useState([0]);
    
    const handleAddRow = () => {
        changeRows(prevRows => [...prevRows, prevRows.length]);
    }

    return (
        <Box>
            <Typography mt={2} mb={2}>{text}</Typography>
            <Table sx={{ minWidth: '100%', border: '1px solid rgba(224, 224, 224, 1)' }} aria-label="simple table">
                <TableHead>
                    <TableRow>
                        <TableCell className={classes.tablecell}></TableCell>
                        <TableCell className={classes.tablecell}>KEY</TableCell>
                        <TableCell className={classes.tablecell}>VALUE</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {
                        rows.map((row, index) => {
                            return <ChangeRow 
                                changeRows={changeRows}
                                rowId={index} 
                                key={index}
                                setData={setData}
                                data={data} 
                            />
                        })
                    }
                </TableBody>
            </Table>
            <Button className={classes.button} variant="contained" onClick={handleAddRow}>Add Row</Button>
        </Box>
    )
}

export default CreateTable;