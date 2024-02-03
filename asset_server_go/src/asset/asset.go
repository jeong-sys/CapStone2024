package asset

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	// _ "github.com/go-sql-driver/mysql"
)

// assetDB
func AssetData(name string, data []byte) {

	// 변수 선언
	var trim_name string = ""
	var connect string = ""
	var split_data_db []string

	// table명 받아옴(공백 제거)
	trim_name = strings.TrimSpace(name)
	// trim_name = strings.Trim(name, " ") // 이렇게 하는 경우 오류
	// 	fmt.Println(trim_name + "1connect")  // 이거는 호출만 됨(1connect만 출력됨) --> 이후 if문도 출력 안되는 문제
	// 	fmt.Println("1connect " + trim_name) // 1connect upload 잘 호출 됨
	fmt.Printf("trim_name : %s\n", trim_name)

	str_data_db := string(data)
	split_data_db = strings.Split(str_data_db, " ")


	

	//MySQL 데이터베이스 연결 정보
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/assets")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database!")




	// table명에 따라 data 받아오는 값 달라짐
	// data를 db에 저장
	switch trim_name {

	// upload table 연결
	case "upload":
		connect = fmt.Sprintf("Connected to %s table", trim_name)
		fmt.Printf(connect + "\n")

		// 내용 저장(test, string데이터만 전달하고 반환, 알아보기)
		// 입력 data : asset upload id[num] name[string] category[num] thum[images] (date:datetime) (count:num) price[num] is[T]
		id, _ := strconv.Atoi(split_data_db[2])
		fmt.Printf("id: %d, type: %s\n", id, reflect.TypeOf(id))

		name := split_data_db[3]
		fmt.Printf("name: %s\n", name)

		category, _ := strconv.Atoi(split_data_db[4])
		fmt.Printf("id: %d, type: %s\n", category, reflect.TypeOf(id))

		thum := split_data_db[5]
		fmt.Printf("thum_test: %s\n", thum)

		price, _ := strconv.Atoi(split_data_db[6])
		fmt.Printf("price: %d, type: %s\n", price, reflect.TypeOf(id))

		isbool := split_data_db[7]
		fmt.Printf("thum_test: %s\n", isbool)

	// file table 연결
	case "file":
		connect = fmt.Sprintf("Connected to %s table", trim_name)
		fmt.Printf(connect)

	// 그 외
	default:
		fmt.Println("생성되지 않은 table 명")
		return
	}

	// fmt.Println(connect)
	//test message
	// message := fmt.Sprintf("%s success", connect)
	// return message
}
