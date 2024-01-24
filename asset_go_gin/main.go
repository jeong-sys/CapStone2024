package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := gin.Default()

	// MongoDB 설정
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("asset_db")

	// 폴더 생성
	createDirIfNotExists("uploads/files")
	createDirIfNotExists("uploads/thumbnails")

	// 정적 파일 서비스
	router.Static("/uploads", "./uploads")
	router.LoadHTMLGlob("templates/*")

	// 라우트 설정
	router.GET("/asset/upload", showAssetUploadHTML)
	router.POST("/asset/upload", func(c *gin.Context) {
		assetUpload(c, db)
	})
	router.GET("/asset/info", showAssetInfoHTML)
	router.GET("/api/asset_info", func(c *gin.Context) {
		assetInfo(c, db)
	})

	router.Run(":8080")
}

func createDirIfNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755) // or os.ModePerm
	}
}

// showAssetUploadHTML, showAssetInfoHTML 함수 정의
func showAssetUploadHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "asset_upload.html", nil)
}

func showAssetInfoHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "asset_info.html", nil)
}
