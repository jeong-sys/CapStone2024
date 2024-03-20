package main

import (
	"log"
	"net/http"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
)

// MySQL
func GetConnector(DB_name string) *sql.DB{
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "12345678",
		Net:                  "http",
		Addr:                 "localhost:3306",
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               DB_name,
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
}


func main() {

	// MySQL 연결 설정

	r.POST("/signup", func(c *gin.Context){
		// 데이터 저장
	})

	r.GET("/login", func(c *gin.Context){
		// 데이터 조회
		
	})
}