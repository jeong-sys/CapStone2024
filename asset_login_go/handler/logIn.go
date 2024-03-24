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

func LogIn(c echo.Context) error {
	user := new(models.User)
	inputpw := user.User_pw

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	db := db.GetConnector()
	fmt.Println("Connected DB")

	// 가입여부 확인
	err := db.QueryRow("SELECT user_id, user_pw FROM users WHERE user_id = ?", user.User_id).Scan(&user.User_id, &inputpw)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "User not found",
		})
	}

	// 비밀번호 검증
	res := hashing.CheckHashPassword(inputpw, user.User_pw)
	if !res {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login Success",
	})
}
