pub mod asset_db;

use tokio::net::{TcpListener, TcpStream};
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use mysql_async::Pool;

// asset부분 요청 처리 부분
pub async fn handle_asset_request(stream: &mut tokio::net::TcpStream, request_parts: &[&str], pool: Pool) {
    match request_parts[0] {
        "upload" => handle_asset_upload(&mut stream, &request_parts[1..], pool.clone()).await,
        "search" => handle_asset_search(&mut stream, &request_parts[1..], pool.clone()).await,
        "inquiry" => handle_asset_inquiry(&mut stream, &request_parts[1..], pool.clone()).await,
        _ => {
            let _ = stream.write_all(b"Unknown asset command\n").await;
        }
    }
}






// 요청 파일이 JSON 형식이라고 가정 - 업로드 부분
#[derive(Debug, Serialize, Deserialize)]
struct UploadRequest {
    name: String,
    category_id: i32,
    image: Vec<u8>,
    file: Vec<u8>,
    price: f64,
    is_disable: bool,
}

pub async handle_asset_upload(mut stream: TcpStream, pool: Pool, request_body: String){
    // JSON 페이로드를 UploadRequest 구조체로 역직렬화
    let upload_request: UploadRequest = serde_json::from_str(&request_body)?;

    // 역직렬화된 데이터를 사용하여 데이터베이스에 에셋 정보 업로드
    asset_upload(&pool, 
        &upload_request.name, 
        upload_request.category_id, 
        upload_request.image, 
        upload_request.file, 
        upload_request.price, 
        upload_request.is_disable).await?;

    // 업로드 성공 메시지를 클라이언트에게 전송
    let response = "Asset uploaded successfully";
    stream.write_all(response.as_bytes()).await?;

    Ok(())
}





/*
// 요청 파일이 JSON 형식이라고 가정 - 검색 부분
#[derive(Debug, Deserialize)]
struct SearchParams {
    category_id: Option<i32>,
    search_query: Option<String>,
}

pub async handle_asset_search(
    mut stream: TcpStream, 
    pool: Pool, 
    category_id: Option<i32>, // 카테고리 ID를 선택적으로 받음
    search_query: Option<String>){ // 검색어도 선택적으로 받음, 현재 예시에서는 사용되지 않음
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


pub async handle_asset_inquiry(){

}
*/