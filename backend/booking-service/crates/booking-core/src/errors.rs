use thiserror::Error;
use uuid::Uuid;

pub type BookingResult<T> = Result<T, BookingError>;

#[derive(Debug, Error)]
pub enum BookingError {
    #[error("validation failed: {field} {reason}")]
    ValidationFailed {
        field: &'static str,
        reason: &'static str,
    },

    #[error("entity not found: {entity} with id {id}")]
    NotFound {
        entity: &'static str,
        id: String,
    },

    #[error("invalid state transition: cannot change order from {current} to {target}")]
    InvalidState {
        current: &'static str,
        target: &'static str,
    },
    
    #[error("repository operation failed for {entity}: {reason}")]
    RepositoryFailed { 
        entity: &'static str, 
        reason: String 
    },
}


#[derive(Debug, Error)]
pub enum ReservationError {
    #[error("event {event_id} is completely sold out")]
    EventSoldOut { event_id: Uuid },

    #[error("not enough capacity for ticket type {ticket_type_id} (requested: {requested}, available: {available})")]
    InsufficientCapacity {
        ticket_type_id: Uuid,
        requested: i32,
        available: i32,
    },

    #[error("reservation lock failed for event {event_id}: {reason}")]
    LockFailed {
        event_id: Uuid,
        reason: String,
    },
}