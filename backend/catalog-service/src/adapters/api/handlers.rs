use axum::{extract::State, Json, response::IntoResponse, http::StatusCode};
use std::sync::Arc;
use crate::ports::EventRepository;

// El State usa un Arc (Atomic Reference Counter) para compartir el repo entre hilos
pub async fn get_events(
    State(repo): State<Arc<dyn EventRepository>>,
) -> impl IntoResponse {
    match repo.find_all().await {
        Ok(events) => (StatusCode::OK, Json(events)).into_response(),
        Err(err) => (StatusCode::INTERNAL_SERVER_ERROR, err).into_response(),
    }
}

pub async fn health_check() -> impl IntoResponse {
    (StatusCode::OK, "ok")
}