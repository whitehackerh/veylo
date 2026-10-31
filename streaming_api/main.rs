use axum::{
    routing::get,
    Json,
    Router,
};
use serde_json::{json, Value};
use tokio::net::TcpListener;

async fn health() -> Json<Value> {
    Json(json!({
        "status": "ok"
    }))
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/api/health", get(health));

    let listener = TcpListener::bind("0.0.0.0:8081")
        .await
        .expect("failed to bind address");

    println!("Streaming API server listening on :8081");

    axum::serve(listener, app)
        .await
        .expect("failed to start server");
}