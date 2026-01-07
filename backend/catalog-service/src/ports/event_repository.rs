use async_trait::async_trait;
use crate::domain::event::Event;

#[async_trait]
pub trait EventRepository: Send + Sync {
    async fn find_all(&self) -> Result<Vec<Event>, String>;
}