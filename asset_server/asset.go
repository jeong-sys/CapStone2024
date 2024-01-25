package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive" 
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Asset struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `form:"name"`
	Uploader      string             `form:"uploader"`
	Price         int                `form:"price"`
	Category      string             `form:"category"`
	Subcategory   string             `form:"subcategory"`
	ThumbnailPath string             `json:"thumbnail_path"`
	FilePath      string             `json:"file_path"`
}

func main() {
	// MongoDB 연결 설정
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 데이터베이스와 컬렉션 선택
	collection := client.Database("asset_server").Collection("asset")

	// Gin 라우터 초기화
	r := gin.Default()

	r.LoadHTMLGlob("html/*")
	r.Static("/save_info", "./save_info")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "send.html", gin.H{"title": "Asset Upload"})
	})

	r.POST("/", func(c *gin.Context) {

		var asset Asset

		thumbnail, _ := c.FormFile("thumbnail")
		file, _ := c.FormFile("file")

		c.Bind(&asset)

		// thumbnailFile, file 경로로 저장
		thumbnailFilename := filepath.Base(thumbnail.Filename)
		fileFilename := filepath.Base(file.Filename)
		thumbnailPath := filepath.Join("./save_info/thumbnail_dir", thumbnailFilename)
		filePath := filepath.Join("./save_info/file_dir", fileFilename)

		c.SaveUploadedFile(thumbnail, thumbnailPath)
		c.SaveUploadedFile(file, filePath)

		asset.ThumbnailPath = "/save_info/thumbnail_dir/" + thumbnailFilename
		asset.FilePath = "/save_info/file_dir/" + fileFilename

		collection.InsertOne(context.TODO(), asset)
		c.Redirect(http.StatusMovedPermanently, "/result/"+asset.Name)
	})

	r.GET("/result/:name", func(c *gin.Context) {
		name := c.Param("name")

		// MongoDB에서 name 필드를 기준으로 데이터 조회
		var asset Asset
		err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&asset)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
			return
		}

		// 썸네일과 정보 표시
		c.HTML(http.StatusOK, "result.html", gin.H{
			"Thumbnail": asset.ThumbnailPath,
			"Asset":     asset,
		})
	})

	r.Run(":8080")
}
