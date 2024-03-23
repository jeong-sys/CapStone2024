package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"capstone.com/module/db"
	"capstone.com/module/hashing"
	"capstone.com/module/models"
	"github.com/labstack/echo"
)

func SignUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	db := db.GetConnector()
	fmt.Println("Connected DB")

	// 아이디 존재 여부 확인
	query_id := "SELECT * FROM users WHERE user_id = " + user.User_id
	result_id := db.QueryRow(query_id).Scan(&user.User_id)
	if result_id != sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existing id",
		})
	}

	// 닉네임 존재 여부 확인
	query_nick := "SELECT * FROM users WHERE nickname = " + user.Nickname
	result_nick := db.QueryRow(query_nick).Scan(&user.Nickname)
	if result_nick != sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existing id",
		})
	}

	// 비밀번호 bycrypt 라이브러리 해싱 처리
	hashpw, err := hashing.HashPassword(user.User_pw)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	user.User_pw = hashpw

	// 유저 생성
	query_r := "INSERT INTO users (user_id, user_pw, nickname, email) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query_r, user.User_id, user.User_pw, user.Nickname, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed SignUp",
		})
	}

	// Success
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})

}
