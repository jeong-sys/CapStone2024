mod config;
mod asset;

use tokio::net::TcpListener;
use config::Settings;
use asset::handle_client;

#[tokio::main] // 비동기 메인 함수
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let settings = Settings::new();
    let address = format!("{}:{}", settings.address, settings.port);
    let listener = TcpListener::bind(&address).await?;
    
    println!("Server listening on {}", address);

    loop {
        let (stream, _) = listener.accept().await?;
        tokio::spawn(async move {
            if let Err(e) = handle_client(stream).await {
                eprintln!("Error occurred while handling connection: {}", e);
            }
        });
    }
}
