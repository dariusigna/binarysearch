// src/apiService.ts
import axios from 'axios';
import { API_BASE_URL } from './config';

const apiClient = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});


export const findIndex = async (value: number): Promise<number> => {
    try {
        const response = await apiClient.get(`/index/${value}`);
        return response.data.index;
    } catch (error) {
        console.error('Error finding index:', error);
        throw error;
    }
};