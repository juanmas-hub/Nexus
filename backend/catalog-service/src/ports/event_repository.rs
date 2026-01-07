use crate::domain::event::Event;

pub trait EventRepository: Send + Sync {
    async fn find_all(&self) -> Result<Vec<Event>, String>;
}