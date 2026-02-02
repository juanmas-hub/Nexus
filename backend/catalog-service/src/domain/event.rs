use serde::{Deserialize, Serialize};
use sqlx::FromRow;
use uuid::Uuid;
use chrono::{DateTime, Utc};

#[derive(Debug, Serialize, Deserialize, FromRow)] // FromRow es clave para SQLx
pub struct Event {
    pub id: Uuid,
    pub title: String,
    pub description: Option<String>, // Option porque tiene '?' en Typescript
    pub image: String,               // Nuevo
    pub category: String,            // Nuevo
    pub event_date: DateTime<Utc>,
    pub location: String,
    pub price: f64,                  // Nuevo (number en TS)
    pub capacity: i32,
}