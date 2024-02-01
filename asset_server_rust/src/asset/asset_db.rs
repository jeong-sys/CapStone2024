use mysql_async::{Pool, Row, params};
use mysql_async::prelude::*;
use std::time::SystemTime;
use crate::error::{AppError, Result};

// 에셋 업로드
pub async fn upload_asset(pool: &Pool, name: &str, category_id: i32, file: Vec<u8>, image: Vec<u8>, price: f64, is_disable: bool) -> Result<()> {
    let mut conn = pool.get_conn().await.map_err(AppError::DbError)?;
    // 이미지와 파일을 먼저 Files와 Images 테이블에 저장
    let image_id: i32 = conn.exec_map(
        "INSERT INTO Images (image) VALUES (:image)",
        params! {"image" => image},
        |id| id,
    ).await?;
    let file_id: i32 = conn.exec_map(
        "INSERT INTO Files (file) VALUES (:file)",
        params! {"file" => file},
        |id| id,
    ).await?;
    // 에셋 정보를 Assets 테이블에 저장
    conn.exec_drop(
        "INSERT INTO Assets (name, category_id, image_id, file_id, upload_date, price, is_disable) VALUES (:name, :category_id, :image_id, :file_id, :upload_date, :price, :is_disable)",
        params! {
            "name" => name,
            "category_id" => category_id,
            "image_id" => image_id,
            "file_id" => file_id,
            "upload_date" => SystemTime::now(),
            "price" => price,
            "is_disable" => is_disable,
        },
    ).await?;
    Ok(())
}

// 카테고리별 에셋 조회
pub async fn get_assets_by_category(pool: &Pool, category_id: i32) -> Result<Vec<(i32, String)>, mysql_async::Error> {
    let mut conn = pool.get_conn().await?;
    let assets = conn.exec_map(
        "SELECT id, name FROM Assets WHERE category_id = :category_id",
        params! {"category_id" => category_id},
        |(id, name)| (id, name),
    ).await?;
    Ok(assets)
}

// 특정 에셋 정보 조회
pub async fn get_asset_details(pool: &Pool, asset_id: i32) -> Result<Row, mysql_async::Error> {
    let mut conn = pool.get_conn().await?;
    let asset = conn.exec_first(
        "SELECT a.id, a.name, c.name as category_name, i.image, f.file, a.upload_date, a.price, a.is_disable FROM Assets a JOIN Categories c ON a.category_id = c.id JOIN Images i ON a.image_id = i.id JOIN Files f ON a.file_id = f.id WHERE a.id = :asset_id",
        params! {"asset_id" => asset_id},
    ).await?.ok_or("Asset not found")?;
    Ok(asset)
}
