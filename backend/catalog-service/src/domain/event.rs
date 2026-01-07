use serde::{Serialize, Deserialize};
use chrono::{DateTime, Utc};
use uuid::Uuid;

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Event {
    pub id: Uuid,  
    pub title: String,      
    pub description: Option<String>, 
    pub location: String, 
    pub event_date: DateTime<Utc>,
    pub capacity: i32,
}