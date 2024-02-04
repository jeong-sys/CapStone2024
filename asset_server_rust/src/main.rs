mod config;
mod asset;
//mod user;
//mod scene;

use tokio::net::{TcpListener, TcpStream};
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use mysql_async::Pool;

#[tokio::main]
async fn main() {
    // 서버 초기화
    let setting = config::Setting::new();
    let address = format!("{}:{}", setting.server_address, setting.port);
    let pool = Pool::new(setting.database_url);

    // 서버 바인딩
    let listener = TcpListener::bind(&address).await?;
    println!("Server listening on {}", address);

    // 클라이언트 연결 처리
    loop {
        let (stream, _) = listener.accept().await?;
        let pool_clone = pool.clone();

        tokio::spawn(async move {
            handle_client(stream, pool_clone).await; // 파라미터로 받은 "stream", pool_clone 사용하여 로직 수행
        });
    }
}


// 요청 처리 함수
async fn handle_client(mut stream: tokio::net::TcpStream, pool: Pool) {
    let mut buffer = [0; 65536];
    // 클라이언트로부터 데이터 읽기
    match stream.read(&mut buffer).await {
        Ok(n) => {
            let request = String::from_utf8_lossy(&buffer[..n]).to_string();
            let request_parts: Vec<&str> = request.split_whitespace().collect();
            if let Some(command) = request_parts.get(0) {
                match *command {
                    "asset" => asset::handle_asset_request(&mut stream, &request_parts[1..], pool.clone()).await, 
                    //"user" => user::handle_user_request(&mut stream, &request_parts[1..], pool.clone()).await,
                    //"scene" => scene::handle_scene_request(&mut stream, &request_parts[1..], pool.clone()).await,
                    _ => {
                        let _ = stream.write_all(b"Unknown command\n").await;
                    }
                }
            }
        },
        Err(e) => eprintln!("Failed to read from socket: {}", e),
    }
}