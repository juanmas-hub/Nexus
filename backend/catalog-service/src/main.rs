mod domain;
mod ports;
mod adapters;

use std::{env, sync::Arc};
use sqlx::postgres::PgPoolOptions;
use dotenvy::dotenv;

#[tokio::main]
async fn main() {
    dotenv().ok();
    
    let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&db_url)
        .await
        .expect("Failed to connect to Postgres");

    let repo = Arc::new(adapters::db::postgres_repo::PostgresEventRepository::new(pool))
        as Arc<dyn ports::EventRepository>;

    let app = adapters::api::create_routes(repo);

    println!("🚀 Catalog Service running on http://0.0.0.0:{}", port);
    
    let listener = tokio::net::TcpListener::bind(format!("0.0.0.0:{}", port)).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}