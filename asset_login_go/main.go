package main


import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"capstone.com/module/db"
)

//json 파싱해서 받아온 값 db에 저장하기
type Info struct{
	User_ID string
	User_PW string
	nickname string
	email string
}


func main() {

	

	r := gin.Default()	

	r.POST("/signup", func(c *gin.Context){


	})

	router.POST("/signup", signUp)
	router.GET("/login", login)

}



