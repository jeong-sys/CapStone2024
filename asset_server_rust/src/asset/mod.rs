use super::asset_db::upload_asset;
use mysql_async::Pool;
use tokio::net::TcpStream;
use tokio::io::{self, AsyncReadExt, AsyncWriteExt};
use futures::future::join_all;


// 요청 처리 함수
pub async fn handle_request(mut stream: tokio::net::TcpStream, pool: mysql_async::Pool, request: String) {
    if request.starts_with("UPLOAD") {
        handle_upload(stream, pool, request).await;
    } else if request.starts_with("SEARCH") {
        handle_category_and_search(stream, pool, request).await;
    } else {
        eprintln!("Unsupported request: {}", request);
    }
}




// 파일 업로드 함수
pub async fn handle_upload(mut stream: TcpStream, pool: Pool) -> Result<(), Box<dyn std::error::Error>> {
    
    /*
    클라이언트 서버에서 오는 요청을 아래의 변수명으로 처리할 로직 필요
    메타버스 팀과 상의 후 진행
    */

    // 데이터베이스에 에셋 정보 업로드
    upload_asset(&pool, name, category_id, image, file, price, is_disable).await?;

    // 업로드 성공 메시지를 클라이언트에게 전송
    let response = "Asset uploaded successfully";
    stream.write_all(response.as_bytes()).await?;

    Ok(())
}




// 카테고리 조회랑 검색 함수
async fn handle_category_and_search(
    mut stream: TcpStream, 
    pool: Pool, 
    category_id: Option<i32>, // 카테고리 ID를 선택적으로 받음
    search_query: Option<String> // 검색어도 선택적으로 받음, 현재 예시에서는 사용되지 않음
) -> Result<(), Box<dyn std::error::Error>> {

    /*
    클라이언트 서버에서 오는 요청을 아래의 변수명으로 처리할 로직 필요
    메타버스 팀과 상의 후 진행
    */

    // 에셋 정보와 썸네일 조회
    let assets_with_thumbnails = get_assets_by_category_and_search(&pool, category_id, search_query).await.map_err(|e| e.into())?;

    // 아이디와 이름 정보를 먼저 클라이언트에 전송
    for (id, name, _) in &assets_with_thumbnails {
        let info = format!("ID: {}, Name: {}\n", id, name);
        stream.write_all(info.as_bytes()).await?;
    }

    // 썸네일 데이터를 포함해 추가 정보 전송
    for (_, _, thumbnail) in assets_with_thumbnails {
        if let Some(thumbnail_data) = thumbnail {
            // 썸네일 데이터 전송, 실제 구현에서는 썸네일 데이터의 전송 방식을 정의해야 함
            stream.write_all(&thumbnail_data).await?;
        }
    }

    Ok(())
}



/*
// 에셋 정보 조회 함수
async fn handle_asset_info_request(pool: &Pool, asset_id: i32) -> Result<AssetDetails, Box<dyn std::error::Error>> {
    let asset_details = get_asset_details(&pool, asset_id).await?;
    Ok(asset_details)
}
*/