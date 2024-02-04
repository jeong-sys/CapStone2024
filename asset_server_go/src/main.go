package main

import (
	"log"
	"net"

	"capstone.com/module/tcp"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if nil != err {
		log.Println(err.Error())
	}
	defer listener.Close()
	log.Printf("프로그램 시작")

	for {
		// 클라이언트 요청 대기
		conn, err := listener.Accept()

		// 예외처리
		if err != nil {
			log.Println(err.Error())
			continue
		} else {
			log.Printf("클라이언트 연결 : %v", conn.RemoteAddr())
		}
		defer conn.Close() // 메인 프로세스 종료시 소켓 종료

		tcp.ConnHandler(conn)
	}
}
