pub struct Setting {
    pub server_address: String,
    pub port: u16,
    pub database_url: String,
}

impl Setting {
    pub fn new() -> Config {
        Config {
            server_address: "192.168.50.88".to_string(),
            port: 8000,
            database_url: "local_url".to_string(), // 실제 데이터베이스 URL로 교체필요.
        }
    }
}

