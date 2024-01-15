package main

// import "net/http" // http패키지
import (
	"encoding/json"
	"net/http"
)

var users = map[string]*User{}

type User struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// 4
func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// 들어오는 요청의 Response Header에 Content-Type 추가
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		//전달받은 http.Handler호출
		next.ServeHTTP(rw, r)
	})
}

func main() {

	// 1
	mux := http.NewServeMux()

	// 2
	userHandler := http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) { // 라우팅 작업, 내가 원하는 경로와 함수 연결
		switch r.Method {
		case http.MethodGet: // 조회
			json.NewEncoder(wr).Encode(users) // 등록된 사용자 json 인코딩하여 응답 보냄
		case http.MethodPost: // 등록
			var user User
			json.NewDecoder(r.Body).Decode(&user) // User구조체로 디코딩
			users[user.Email] = &user
			json.NewEncoder(wr).Encode(user) // json 인코딩하여 응답 보냄

		}
		//wr.Write([]byte("hello"))
	})

	// 3 /users에 미들웨어와 핸들러 등록
	mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
	http.ListenAndServe(":8080", mux) // http패키지 이용하여 서버 띄우기(포트 실행)
}
