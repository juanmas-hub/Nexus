use axum::routing::get;
use axum::Router;
use std::sync::Arc;
use crate::ports::EventRepository;
use crate::adapters::api::handlers;

pub fn create_routes(repo: Arc<dyn EventRepository>) -> Router {
    Router::new()
        .route("/health", get(handlers::health_check))
        .route("/events", get(handlers::get_events))
        .with_state(repo)
}