use serde::{Serialize, Deserialize};
use chrono::{DateTime, Utc};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Event {
    pub id: String,
    pub title: String,
    pub description: String,
    pub location: String,
    pub date: DateTime<Utc>,
    pub capacity: i32,
}