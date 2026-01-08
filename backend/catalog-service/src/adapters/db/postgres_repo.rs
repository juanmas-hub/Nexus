use sqlx::PgPool;
use crate::domain::event::Event;
use crate::ports::EventRepository;
use async_trait::async_trait;

pub struct PostgresEventRepository {
    pool: PgPool,
}

impl PostgresEventRepository {
    pub fn new(pool: PgPool) -> Self {
        Self { pool }
    }
}

#[async_trait]
impl EventRepository for PostgresEventRepository {
    async fn find_all(&self) -> Result<Vec<Event>, String> {
        let records = sqlx::query_as!(
            Event,
            r#"SELECT id, title, description, location, event_date, capacity FROM events"#
        )
        .fetch_all(&self.pool)
        .await
        .map_err(|e: sqlx::Error| e.to_string())?;

        Ok(records)
    }
}