use mysql_async::{Pool, Row, params};
use mysql_async::prelude::*;
use std::time::SystemTime;
use image::{DynamicImage, ImageOutputFormat::Png, imageops::FilterType};


async fn create_thumbnail(image_data: Vec<u8>) -> Vec<u8> {
    let thumbnail_data = {
        let img = image::load_from_memory(&image_data).unwrap_or_else(|_| DynamicImage::new_rgb8(128, 128));
        let resized_img = img.resize(128, 128, FilterType::Nearest);
        let mut thumbnail = Vec::new();
        resized_img.write_to(&mut thumbnail, image::ImageOutputFormat::Png).unwrap_or_default();
        thumbnail
    };

    thumbnail_data
}

// 에셋 업로드
pub async fn asset_upload(pool: &Pool, name: &str, category_id: i32, file: Vec<u8>, image: Vec<u8>, price: f64, is_disable: bool){
    let mut conn = pool.get_conn().await?;
    
    // 이미지와 파일을 Files와 Images 테이블에 저장
    let image_id: i32 = conn.exec_map(
        "INSERT INTO Images (image) VALUES (:image)", params! {"image" => &image}, |id| id,
    ).await?;
    
    let file_id: i32 = conn.exec_map(
        "INSERT INTO Files (file) VALUES (:file)", params! {"file" => &file}, |id| id,
    ).await?;

    // 썸네일 생성
    let thumbnail_data = create_thumbnail(image.clone()).await?;

    // 에셋 정보와 썸네일을 Assets 테이블에 저장
    conn.exec_drop(
        "INSERT INTO Assets (name, category_id, image_id, file_id, thumbnail, upload_date, price, is_disable) VALUES (:name, :category_id, :image_id, :file_id, :thumbnail, :upload_date, :price, :is_disable)",
        params! {
            "name" => name,
            "category_id" => category_id,
            "image_id" => image_id,
            "file_id" => file_id,
            "thumbnail" => thumbnail_data,
            "upload_date" => SystemTime::now().duration_since(SystemTime::UNIX_EPOCH)?.as_secs(),
            "price" => price,
            "is_disable" => is_disable,
        },
    ).await?;

    Ok(())
}


/*
// 카테고리 ID와 검색어를 사용한 에셋 정보와 썸네일 조회
pub async fn get_assets_by_category_and_search(
    pool: &Pool, 
    category_id: Option<i32>, // 카테고리 선택은 선택적으로 만듭니다.
    search_query: Option<String>, // 검색어도 선택적입니다.
) -> Result<Vec<(i32, String, Option<Vec<u8>>)>> {
    let mut conn = pool.get_conn().await.map_err(AppError::DbError)?;

    // 동적 쿼리 구성
    let mut query = "SELECT id, name, thumbnail FROM Assets".to_string();
    let mut conditions = vec![];
    let mut params = params! {};

    if let Some(cid) = category_id {
        conditions.push("category_id = :category_id");
        params.insert("category_id", cid);
    }

    if let Some(search) = search_query {
        conditions.push("name LIKE :search");
        params.insert("search", format!("%{}%", search));
    }

    if !conditions.is_empty() {
        query += " WHERE ";
        query += &conditions.join(" AND ");
    }

    let assets_with_thumbnails = conn.query_map(
        query,
        params,
        |(id, name, thumbnail): (i32, String, Option<Vec<u8>>)| (id, name, thumbnail),
    ).await.map_err(AppError::DbError)?;

    Ok(assets_with_thumbnails)
}
*/




/*
// 특정 에셋 정보 조회
pub async fn get_asset_details(pool: &Pool, asset_id: i32) -> Result<Row, mysql_async::Error> {
    let mut conn = pool.get_conn().await?;
    let asset = conn.exec_first(
        "SELECT a.id, a.name, c.name as category_name, i.image, f.file, a.upload_date, a.price, a.is_disable FROM Assets a JOIN Categories c ON a.category_id = c.id JOIN Images i ON a.image_id = i.id JOIN Files f ON a.file_id = f.id WHERE a.id = :asset_id",
        params! {"asset_id" => asset_id},
    ).await?.ok_or("Asset not found")?;
    Ok(asset)
}
*/