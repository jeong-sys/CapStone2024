package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func setRouter(router *gin.Engine) { // router
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "send.html", gin.H{
			"title": "업로드 페이지", // 제출시, 빈화면 뜸 (제출되었습니다 표시 뜨게 하기)
			"url":   "/info",
		})
	})

	router.POST("/", func(c *gin.Context) {
		fmt.Println(c.PostForm("input")) // 사용자 업로드(DB저장 필요)
	})

	router.GET("/info", func(c *gin.Context) {
		c.HTML(http.StatusOK, "info.html", gin.H{
			"content": "내용", // '/'에서 받아온 내용 표시(DB저장 받아오기 필요)
		})
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*.html")
	setRouter(router)
	_ = router.Run(":8080")
}
