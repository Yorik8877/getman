import axios from 'axios';
import {  getHeadersAndParams } from '../utils/common-utils';

export const getData = async (formData, jsonText, paramData, headerData) => {

    const apiType = formData.type.toLowerCase();            // Тип запроса
    const apiUrl = formData.url;                            // URL тестируемого ресурса
    const apiHeaders = getHeadersAndParams(headerData);     // Заголовки для тестируемого ресурса
    const apiParams = getHeadersAndParams(paramData);       // Параметры (?) для тестируемого ресурса
    const sendUrl = 'http://localhost:8000/data';           // Внешний сервер

    try {
        return await axios({
            method: 'POST',
            url: sendUrl,
            data: {
                method: apiType,
                apiUrl: apiUrl,
                body: jsonText,
                headers: apiHeaders,
                params: apiParams
            }
        })
    } catch (error) {
        console.log('Error while getting the response ', error);
        return 'Error: no response';
    }
}
