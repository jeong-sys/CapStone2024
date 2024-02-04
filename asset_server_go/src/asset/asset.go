package asset

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"capstone.com/module/db"
)

// assetDB
func AssetData(name string, data []byte) {

	// 변수 선언
	var trim_name string = ""
	var connect string = ""
	var split_data_db []string

	// table명 받아옴(공백 제거)
	trim_name = strings.TrimSpace(name)
	fmt.Printf("trim_name : %s\n", trim_name)

	str_data_db := string(data)
	split_data_db = strings.Split(str_data_db, " ")

	// ===================  DB연결  ===================================================
	db_name := "cap_asset" // asset 일때 db_name = cap_asset

	db := db.GetConnector(db_name)
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database!")

	// ===================  DB업로드(table 연결)  =======================================

	// table명에 따라 data 받아오는 값 달라짐 // data를 알맞는 table에 저장
	switch trim_name {

	// upload table 연결
	case "upload":
		connect = fmt.Sprintf("Connected to %s table", trim_name)
		fmt.Printf(connect + "\n")

		// 내용 저장(test : string데이터만 전달하고 반환)
		// 입력 data : asset upload id[num] name[string] category[num] thum[images] (date:datetime) (count:num) price[num] is[T]

		//id, _ := strconv.Atoi(split_data_db[2]) // 우선 table저장은 자동 증가로 설정
		//fmt.Printf("id: %d, type: %s\n", id, reflect.TypeOf(id))

		name := split_data_db[2]
		fmt.Printf("name: %s\n", name)

		category_id, _ := strconv.Atoi(split_data_db[3])
		fmt.Printf("id: %d, type: %s\n", category_id, reflect.TypeOf(category_id))

		thumbnail := split_data_db[4]
		fmt.Printf("thum_test: %s\n", thumbnail)

		price, _ := strconv.Atoi(split_data_db[5])
		fmt.Printf("price: %d, type: %s\n", price, reflect.TypeOf(price))

		isbool := split_data_db[6]
		fmt.Printf("thum_test: %s\n", isbool)

		query := "INSERT INTO assets (name, category_id, thumbnail, upload_date, download_count, price, is_disable) VALUES (?, ?, ?, NOW(), 0, ?, ?)"
		_, err := db.Exec(query, name, category_id, thumbnail, price, isbool)
		if err != nil {
			log.Fatalf("Failed to insert data: %v", err)
		}

		// client에게 id와 성공 메시지 반환
		var asset_id int = 0
		err = db.QueryRow("SELECT LAST_INSERT_ID();").Scan(&asset_id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(asset_id)
		message := "Success upload table"
		fmt.Println(message)

		// tcp.Receive_Id(asset_id)

	// file table 연결
	case "file":
		connect = fmt.Sprintf("Connected to %s table", trim_name)
		fmt.Printf(connect)

	// 그 외
	default:
		fmt.Println("생성되지 않은 table 명")
		return
	}
}
