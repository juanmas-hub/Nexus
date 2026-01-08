export interface EventItem {
  id: string;
  title: string;
  description?: string;
  image: string;
  category: string; 
  event_date: string;  
  location: string;
  price: number;
  capacity: number;
}