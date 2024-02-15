use super::asset_db::upload_asset;
use mysql_async::Pool;
use tokio::net::TcpStream;
use tokio::io::{self, AsyncReadExt, AsyncWriteExt};
use futures::future::join_all;


// 파일 업로드 함수
pub async fn handle_upload(mut stream: TcpStream, pool: Pool) -> Result<(), Box<dyn std::error::Error>> {
 
    // 데이터베이스에 에셋 정보 업로드
    upload_asset(&pool, name, category_id, file, image, price, is_disable).await?;

    // 업로드 성공 메시지를 클라이언트에게 전송
    let response = "Asset uploaded successfully";
    stream.write_all(response.as_bytes()).await?;

    Ok(())
}




// 카테고리 조회 함수
async fn handle_category_selection(pool: &Pool, category_id: i32) -> Result<Vec<(i32, String, Option<Vec<u8>>)>, AppError> {
    let assets = get_assets_by_category(&pool, category_id).await?;
    let thumbnail_futures = assets.iter().map(|(id, _)| get_thumbnail_by_asset_id(&pool, *id));
    
    // 모든 썸네일을 병렬로 조회합니다.
    let thumbnails = join_all(thumbnail_futures).await.into_iter().collect::<Result<Vec<_>, _>>()?;
    
    let assets_with_thumbnails = assets.into_iter().zip(thumbnails).map(|((id, name), thumbnail)| (id, name, Some(thumbnail))).collect();
    
    Ok(assets_with_thumbnails)
}




// 에셋 정보 조회 함수
async fn handle_asset_info_request(pool: &Pool, asset_id: i32) -> Result<AssetDetails, Box<dyn std::error::Error>> {
    let asset_details = get_asset_details(&pool, asset_id).await?;
    Ok(asset_details)
}