package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"capstone.com/module/db"
	"capstone.com/module/hashing"
	"capstone.com/module/models"
	"github.com/labstack/echo"
)

func LogIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	inputpw := user.User_pw

	db := db.GetConnector()
	fmt.Println("Connected DB")

	result_email := fmt.Sprintf("SELECT * FROM users WHERE email='%s';", user.Email)
	fmt.Println(result_email)

	// --- 아이디, 비번으로 변경 필요 --- 오류

	// 아이디 존재 여부 확인(회원인지 확인하기 위함)
	query_id := fmt.Sprintf("SELECT * FROM users WHERE user_id ='%s';", user.User_id)
	fmt.Println(query_id)
	result_id := db.QueryRow(query_id).Scan(&user.User_id)
	if result_id == sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Not existing nickname",
		})
	}

	res := hashing.CheckHashPassword(user.User_pw, inputpw)

	// 비밀번호 검증 실패
	if !res {
		return echo.ErrUnauthorized
	}

	// 토큰 발행
	accessToken, err := hashing.CreateJWT(user.Email)
	if err != nil {
		return echo.ErrInternalServerError
	}

	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(time.Hour * 24)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login Success",
	})
}

func MockData() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Mock Data를 생성한다.
		list := map[string]string{
			"1": "고양이",
			"2": "사자",
			"3": "호랑이",
		}
		return c.JSON(http.StatusOK, list)
	}
}
