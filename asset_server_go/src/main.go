package main

import (
	"fmt"
	"io"
	"log"
	"module/asset"
	"net"
	"reflect"
	"strings"
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
		go ConnHandler(conn)
	}
}

// 고루틴 처리
func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 65536)
	var conn_split_table string = ""
	var conn_split_db string = ""

	for {
		// err 처리
		n, err := conn.Read(recvBuf)
		if err != nil {
			if io.EOF == err {
				log.Println("err(eof):", err)
				log.Printf("연결 종료: %v", conn.RemoteAddr().String())
			}
			log.Println("err:", err)
			return
		}

		if 0 < n {
			// 받아온 길이만큼 슬라이스 잘라서 출력
			data := recvBuf[:n]

			str_data := string(data)
			log.Println(str_data)
			_, err = conn.Write(data[:n])

			// 함수 연결
			conn_split := strings.Split(str_data, " ")

			conn_split_db = conn_split[0]    // "받아온 값 0번째" - DB 구별
			conn_split_table = conn_split[1] // "받아온 값 1번째" - table 구별

			// asset 함수 연결(asset DB 연결)
			if conn_split_db == "asset" {

				fmt.Println("data type : ", reflect.TypeOf(data))

				fmt.Println("conn_split_db :" + conn_split_db)
				fmt.Println("conn_split_table :" + conn_split_table)

				asset.AssetData(conn_split_table, data) // 다른 파일 함수 연결
				// fmt.Println(message2) // test message

			}

			// err 처리
			if err != nil {
				log.Println("err:", err)
				return
			}
		}
	}
}
