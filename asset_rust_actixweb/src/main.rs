use actix_web::{web, App, HttpServer};
use actix_files::Files;
use mongodb::{Client, options::ClientOptions};
use std::fs;

mod asset; // asset 모듈 포함

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let client_options = ClientOptions::parse("mongodb://localhost:27017").await.unwrap();
    let client = Client::with_options(client_options).unwrap();
    let db = client.database("asset_db");

    // 폴더 생성 로직
    create_dir_if_not_exists("uploads/files");
    create_dir_if_not_exists("uploads/thumbnails");

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(db.clone()))
            .service(Files::new("/uploads", "uploads/")) // 정적 파일 서비스
            .route("/asset/upload", web::get().to(asset::show_asset_upload_html))
            .route("/asset/upload", web::post().to(asset::asset_upload))
            .route("/asset/info", web::get().to(asset::show_asset_info_html)) // HTML 페이지 반환
            .route("/api/asset_info", web::get().to(asset::asset_info)) // JSON 데이터 반환
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}

fn create_dir_if_not_exists(dir: &str) {
    if !std::path::Path::new(dir).exists() {
        fs::create_dir_all(dir).expect("Failed to create directory");
    }
}






