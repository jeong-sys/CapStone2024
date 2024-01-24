package main

import (
	// "context"
	// "context"
	"fmt"
	// "log"

	// "io/ioutil"
	"net/http"

	// "github.com/mholt/binding"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Asset struct {
	ID          bson.ObjectId `bson:"_id, omitempty"`
	Name        string        `bson:"name" json:"name"`
	Uploader    string        `bson:"uploader" json:"uploader"`
	Price       int           `bson:"price" json:"price"`
	Category    string        `bson:"category" json:"category"`
	SubCategory string        `bson:"subcategory" json:"subcategory"`
}


func setRouter(router *gin.Engine) { // router

	//router 핸들러 정의
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "send.html", gin.H{
			"title": "업로드 페이지", // 제출시, 빈화면 뜸 (제출되었습니다 표시 뜨게 하기)  -- 1) 입력정보 terminal에 띄우기 (o)
			"url":   "/info",
		})
	})

	router.POST("/", func(c *gin.Context) {

		// 출력
		jsonData, err := c.GetRawData()
		if err != nil {
			fmt.Println("Error reading JSON data:", err)
		}

		fmt.Println("Received JSON data:", string(jsonData))


		// fmt.Println(c.PostForm("name")) // 사용자 업로드(DB저장 필요) ---> 2) 띄운 정보들 mongo 저장
		// fmt.Println(c.PostForm("uploader"))
		// fmt.Println(c.PostForm("price"))
		// fmt.Println(c.PostForm("category"))
		// fmt.Println(c.PostForm("subcategory"))

		// fmt.Println(c.PostForm("thumbnail")) --> 4) 파일 DB 업로드
		// fmt.Println(c.PostForm("file"))

	})

	router.GET("/info", func(c *gin.Context) {
		c.HTML(http.StatusOK, "info.html", gin.H{
			"content": "내용", // '/'에서 받아온 내용 표시(DB저장 받아오기 필요)
		})
	})
}

// func APISearch(c *gin.Context){

// 	c.Header("Content-Type", "application/json charset=utf-8")
// 	resp, err := http.Get("http://localhost:8080/")

// 	if err != nil{
// 		panic(err)
// 	}

// 	defer resp.Body.Close()

// 	respBody, err := ioutil.ReadAll(resp.Body)
// 	if err == nil{
// 		str := string(respBody)
// 		println(str)
// 		fmt.Printf("%S\n",str)
// 	}
// }

//func API
//json으로 변경된 것(js 코드 사용) 가져와서 db에 넣기

func main() {
	router := gin.Default()            // gin으로 router만들기
	router.LoadHTMLGlob("html/*.html") // 템플릿 불러옴
	setRouter(router)
	_ = router.Run(":8080")
}
