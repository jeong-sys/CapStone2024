package main

// 모든 HTTP 요청에 대해 "Hello Docker!!" 라는 응답을 보냄
// 포트 8080로 요청을 받는 서버 애플리케이션으로 동작함
// 클라이언트로부터 요청을 받으면 received request라는 메시지를 표준으로 출력함

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request")
		fmt.Fprintf(w, "Hello Docker!!")
	})

	log.Println("start server")
	server := &http.Server{
		Addr: ":8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
