mod config;
mod asset; 
mod error;

use config::Settings;
use asset::{handle_upload, handle_category_and_search};
use tokio::net::TcpListener;
use tokio::io::AsyncReadExt;
use mysql_async::Pool;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let settings = Settings::new();
    let address = format!("{}:{}", settings.address, settings.port);
    let pool = Pool::new(settings.database_url);

    let listener = TcpListener::bind(&address).await?;
    println!("Server listening on {}", address);

    loop {
        let (stream, _) = listener.accept().await?;
        let pool_clone = pool.clone();

        tokio::spawn(async move {
            handle_client(stream, pool_clone).await;
        });
    }
}

async fn handle_client(mut stream: tokio::net::TcpStream, pool: Pool) {
    let mut buffer = [0; 65536];
    match stream.read(&mut buffer).await {
        Ok(n) => {
            let request = String::from_utf8_lossy(&buffer[..n]).trim().to_string();
            match asset::handle_request(stream, pool, request).await {
                Ok(()) => {}, // 성공적인 요청 처리
                Err(e) => {
                    // 에러 응답을 클라이언트에게 전송
                    let error_message = format!("Error: {}\n", e);
                    if let Err(e) = stream.write_all(error_message.as_bytes()).await {
                        eprintln!("Failed to send error response to client: {}", e);
                    }
                }
            }
        },
        Err(e) => {
            let error_message = format!("Error reading from socket: {}\n", e);
            if let Err(e) = stream.write_all(error_message.as_bytes()).await {
                eprintln!("Failed to send error response to client: {}", e);
            }
        },
    }
}

