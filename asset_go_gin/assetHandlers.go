package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func assetUpload(c *gin.Context, db *mongo.Database) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File
	assetData := bson.M{}

	for key, fileHeaders := range files {
		for _, fileHeader := range fileHeaders {
			filename := filepath.Base(fileHeader.Filename) // sanitize filename if necessary
			filepath := "uploads/files/" + filename
			if key == "thumbnail" {
				filepath = "uploads/thumbnails/" + filename
			}

			// 파일 저장
			if err := c.SaveUploadedFile(fileHeader, filepath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			assetData[key] = filepath
		}
	}

	// 텍스트 필드 처리
	for key, values := range form.Value {
		assetData[key] = values[0]
	}

	// MongoDB에 데이터 저장
	collection := db.Collection("assets")
	if _, err := collection.InsertOne(context.Background(), assetData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Asset uploaded")
}

func assetInfo(c *gin.Context, db *mongo.Database) {
	var assets []bson.M
	collection := db.Collection("assets")

	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var asset bson.M
		if err := cursor.Decode(&asset); err != nil {
			log.Fatal(err)
		}
		assets = append(assets, asset)
	}

	c.JSON(http.StatusOK, assets)
}
