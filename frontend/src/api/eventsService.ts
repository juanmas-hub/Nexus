import { api } from './axios';
import type { EventItem } from '../types';

export const getEvents = async (): Promise<EventItem[]> => {
    const response = await api.get<EventItem[]>('/events');
    return response.data;
};