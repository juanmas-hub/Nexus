use axum::{routing::get, Router};
use booking_infra::Database;
use clap::Parser;
use sqlx::PgPool;
use tokio::net::TcpListener;
use tokio_util::sync::CancellationToken;
use tracing::{info, Level};
use tracing_subscriber::{fmt, EnvFilter};

#[derive(Parser, Debug)]
#[command(name = "booking-service")]
struct Args {
    #[arg(short, long, env = "PORT", default_value = "3000")]
    port: u16,

    #[arg(long, env = "DATABASE_URL")]
    database_url: String,

    #[arg(short, long, action = clap::ArgAction::Count)]
    verbose: u8,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let _ = dotenvy::dotenv();

    let args = Args::parse();
    init_tracing(args.verbose);

    let cancel_token = CancellationToken::new();
    spawn_signal_handler(cancel_token.clone());

    info!("Conectando a la base de datos...");
    let pool = PgPool::connect(&args.database_url).await?;
    let db = Database::new(pool);

    // 5. Configurar el Router de Axum inyectando la DB
    let app = Router::new()
        .route("/health", get(health_check))
        // Acá agregaremos el POST /bookings luego
        .with_state(db);

    let addr = format!("0.0.0.0:{}", args.port);
    let listener = TcpListener::bind(&addr).await?;

    println!();
    println!("NEXUS BOOKING SERVICE");
    println!("  Listening on: http://{}", addr);
    println!("  Database:     Connected");
    println!("  Press Ctrl+C to stop");
    println!();

    axum::serve(listener, app)
        .with_graceful_shutdown(async move {
            cancel_token.cancelled().await;
            info!("Servidor web detenido.");
        })
        .await?;

    Ok(())
}

fn init_tracing(verbose: u8) {
    let level = match verbose {
        0 => Level::INFO,
        1 => Level::DEBUG,
        _ => Level::TRACE,
    };

    let filter = EnvFilter::from_default_env()
        .add_directive(format!("booking_api={level}").parse().unwrap())
        .add_directive(format!("booking_infra={level}").parse().unwrap())
        .add_directive(format!("booking_core={level}").parse().unwrap())
        .add_directive("tower_http=debug".parse().unwrap())
        .add_directive("sqlx=warn".parse().unwrap());

    fmt()
        .with_env_filter(filter)
        .with_target(false)
        .with_file(false)
        .with_line_number(false)
        .init();
}

fn spawn_signal_handler(cancel_token: CancellationToken) {
    tokio::spawn(async move {
        let _ = tokio::signal::ctrl_c().await;
        println!();
        info!("Apagando el sistema (graceful shutdown)...");
        cancel_token.cancel();
    });
}

// Un handler de prueba básico para asegurar que todo levanta bien
async fn health_check() -> &'static str {
    "Booking Service is UP"
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn cli_parses_defaults() {
        let args = Args::parse_from(["booking-service", "--database-url", "postgres://test"]);
        assert_eq!(args.port, 3000);
        assert_eq!(args.verbose, 0);
    }

    #[test]
    fn cli_parses_port_and_verbose() {
        let args = Args::parse_from([
            "booking-service",
            "-p",
            "8080",
            "-vv",
            "--database-url",
            "postgres://test"
        ]);
        assert_eq!(args.port, 8080);
        assert_eq!(args.verbose, 2);
    }
}