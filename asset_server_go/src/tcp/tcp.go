package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
	"strings"

	"capstone.com/module/asset"
)

var Receive_Id int = 2

// 고루틴 처리
func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)
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

			// 클라이언트한테 반환 못 함(왜?? tcp연결에 따른 문제 같음)

			// if Receive_Id > 0 {

			// 	str_Receive_Id := string(Receive_Id)
			// 	id_data := []byte(str_Receive_Id)
			// 	_, err = conn.Write(id_data)

			// } else {
			_, err = conn.Write(data[:n])

			//}
			// err 처리
			if err != nil {
				log.Println("err:", err)
				return
			}

		}
	}
}

// func ReceiveId(asset_Id int) int {

// 	fmt.Println(asset_Id)
// 	Receive_Id := asset_Id
// 	return Receive_Id

// }
