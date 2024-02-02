// 서버 설정을 위한 구조체
pub struct Settings {
    pub address: String, // 서버 주소
    pub port: u16,       // 포트 번호
}

impl Settings {
    // 새 설정을 생성하는 함수
    pub fn new() -> Self {
        Settings {
            address: "192.168.50.88".to_string(), // 로컬 주소
            port: 8000,                       // 사용할 포트 번호
        }
    }
}
