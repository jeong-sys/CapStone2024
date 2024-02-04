// 사용안함.



use std::fmt;

#[derive(Debug)]
pub enum AppError {
    DbError(mysql_async::Error),
    NotFound(String),
    IoError(std::io::Error),
    ImageError(image::ImageError),
}

impl fmt::Display for AppError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            AppError::DbError(e) => write!(f, "Database error: {}", e),
            AppError::NotFound(e) => write!(f, "Not found: {}", e),
            AppError::IoError(e) => write!(f, "IO error: {}", e),
        }
    }
}

impl From<mysql_async::Error> for AppError {
    fn from(err: mysql_async::Error) -> AppError {
        AppError::DbError(err)
    }
}

impl From<std::io::Error> for AppError {
    fn from(err: std::io::Error) -> AppError {
        AppError::IoError(err)
    }
}

impl From<image::ImageError> for AppError {
    fn from(err: image::ImageError) -> Self {
        AppError::ImageError(err)
    }
}

// 애플리케이션에서 사용되는 `Result` 타입을 정의합니다.
pub type Result<T> = std::result::Result<T, AppError>;
