package main

import (
	"log"

	"capstone.com/module/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo" // echo프레임워크 사용
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.POST("/signup", handler.SignUp)
	e.GET("/login", handler.LogIn)

	// 목데이터로 테스트
	// e.GET("/test", handler.MockData(), md.JWTWithConfig(md.JWTConfig{
	// 	SigningKey:  []byte(os.Getenv("SECRET_KEY")),
	// 	TokenLookup: "cookie:access-token",
	// }))

	e.Logger.Fatal(e.Start(":8000"))

}
