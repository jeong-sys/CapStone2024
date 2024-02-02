package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if nil != err {
		log.Println(err.Error())
	}
	defer listener.Close()
	log.Printf("프로그램 시작")
	
	// fmt.Println(asset.AssetData())

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

		// 고루틴 처리
		go func() {
			recvBuf := make([]byte, 65536)
			for {
				// err 처리
				n, err := conn.Read(recvBuf)
				if err != nil {
					if io.EOF == err {
						log.Println("err(eof):", err)
						log.Printf("연결 종료: %v", conn.RemoteAddr().String())
						//break
					}
					log.Println("err:", err)
					//break
					return
				}

				if 0 < n {
					// 받아온 길이만큼 슬라이스 잘라서 출력
					data := recvBuf[:n]
					log.Println(string(data))
					_, err = conn.Write(data[:n])

					if err != nil {
						log.Println("err:", err)
						break
					}
				}
			}
		}()
	}
}