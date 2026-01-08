mod domain;
mod ports;
mod adapters;

use std::{env, sync::Arc, time::Duration};
use sqlx::postgres::PgPoolOptions;
use dotenvy::dotenv;
use tokio::time::sleep;

#[tokio::main]
async fn main() {
    dotenv().ok();
    
    let db_url = env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let port = env::var("PORT").unwrap_or_else(|_| "8082".to_string());

    println!("Attempting to connect to database...");

    let mut pool = None;
    let mut attempts = 0;
    
    while attempts < 5 {
        match PgPoolOptions::new()
            .max_connections(2)
            .acquire_timeout(Duration::from_secs(5))
            .connect(&db_url)
            .await 
        {
            Ok(p) => {
                pool = Some(p);
                println!("Connected to Postgres successfully!");
                break;
            }
            Err(e) => {
                attempts += 1;
                println!("Connection attempt {} failed: {}. Retrying in 5s...", attempts, e);
                sleep(Duration::from_secs(5)).await;
            }
        }
    }

    let pool = pool.expect("Could not connect to database after several attempts");

    let repo = Arc::new(adapters::db::postgres_repo::PostgresEventRepository::new(pool))
        as Arc<dyn ports::EventRepository>;

    let app = adapters::api::create_routes(repo);

    let address = format!("0.0.0.0:{}", port);
    println!("Catalog Service running on http://{}", address);
    
    let listener = tokio::net::TcpListener::bind(&address).await.expect("Failed to bind port");
    axum::serve(listener, app).await.expect("Failed to start server");
}