use actix_web::{web, HttpResponse, Responder};
use actix_multipart::Multipart;
use futures::{TryStreamExt};
use mongodb::{Database};
use bson::{Document, Bson};
use std::io::Write;
use sanitize_filename::sanitize;
use html_entities::decode_html_entities;
use percent_encoding::{percent_decode, utf8_percent_encode, NON_ALPHANUMERIC};

// HTML 폼 표시
pub async fn show_asset_upload_html() -> impl Responder {
    HttpResponse::Ok().content_type("text/html").body(include_str!("static/asset_upload.html"))
}

pub async fn show_asset_info_html() -> impl Responder {
    HttpResponse::Ok().content_type("text/html").body(include_str!("static/asset_info.html"))
}


pub async fn asset_upload(db: web::Data<Database>, mut payload: Multipart) -> impl Responder {
    let mut asset_data = Document::new();
    //payload.content_disposition().get_name();
    while let Ok(Some(mut field)) = payload.try_next().await {
        let content_disposition = field.content_disposition();
        let field_name = content_disposition.get_name().unwrap().to_string();
        let filename = content_disposition.get_filename().map(sanitize);
        
        println!("{}", asset_data);

        match field_name.as_str() {
            "thumbnail" | "file" => {
                if let Some(filename) = filename {
                    // 파일명 인코딩
                    let decoded_filename = percent_decode(filename.as_bytes())
                        .decode_utf8()
                        .unwrap()
                        .to_string();
                    let folder = if field_name == "thumbnail" { "uploads/thumbnails" } else { "uploads/files" };
                    let filepath = format!("{}/{}", folder, decoded_filename);

                    // 파일 객체를 여기서 생성
                    let mut file = std::fs::File::create(&filepath).unwrap();
                    while let Some(chunk) = field.try_next().await.unwrap() {
                        // 각 청크를 파일에 쓰기
                        file.write_all(&chunk).unwrap();
                    }

                    asset_data.insert(field_name, Bson::String(filepath));
                }
            },
            _ => {
                if let Ok(Some(data)) = field.try_next().await {
                    match String::from_utf8(data.to_vec()) {
                        Ok(text) => {
                            let decoded_text = decode_html_entities(&text).unwrap_or_default(); // 수정
                            asset_data.insert(field_name, Bson::String(decoded_text));
                        }
                        Err(e) => {
                            eprintln!("Invalid UTF-8 sequence: {}", e);
                            // 적절한 에러 처리
                        }
                    }
                }
            }
        }
    }

    db.collection("assets").insert_one(asset_data, None).await.unwrap();
    HttpResponse::Ok().body("Asset uploaded")
}

// 에셋 정보 조회
pub async fn asset_info(db: web::Data<Database>) -> impl Responder {
    let mut cursor: mongodb::Cursor<Document> = db.collection("assets").find(None, None).await.unwrap();
    let mut assets = Vec::new();
    
    while let Some(asset) = cursor.try_next().await.unwrap() {
        assets.push(asset);
    }

    let json_assets = serde_json::to_string(&assets).unwrap();
    HttpResponse::Ok().content_type("application/json").body(json_assets)
}

