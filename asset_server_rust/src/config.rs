pub struct Settings {
    pub address: String, // 서버 주소
    pub port: u16,       // 포트 번호
    pub database_url: String, // 데이터베이스 URL
}

impl Settings {
    // 새 설정을 생성하는 함수
    pub fn new() -> Self {
        // 환경변수에서 데이터베이스 URL을 가져오거나, 기본값을 사용
        let database_url = std::env::var("DATABASE_URL").unwrap_or_else(|_| "mysql://user:password@localhost/dbname".to_string());

        Settings {
            address: "192.168.50.88".to_string(),
            port: 8000,
            database_url, // 데이터베이스 URL 설정
        }
    }
}

