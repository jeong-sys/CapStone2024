package handler

import (
	"fmt"
	"net/http"

	"capstone.com/module/db"
	"capstone.com/module/hashing"
	"capstone.com/module/models"

	// db "command-line-argumentsC:\\Users\\sysailab\\capstone\\CapStone2024\\asset_login_go\\db\\connect.go"
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

	// 이미 아이디 존재할 경우
	result_id := db.Find(&user, "user_id=?", user.user_id)
	if result_id.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existring id",
		})
	}

	// 이미 닉네임 존재할 경우
	result_nick := db.Find(&user, "nickname=?", user.nickname)
	if result_nick.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existing nickname",
		})

	}

	// 비밀번호 bycrypt 라이브러리 해싱 처리
	hashpw, err := handler.HashPassword(user.user_pw)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	user.User_PW = hashpw

	if err := db.GetConnector(&user); err.Eror != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed SignUp",
		})
	}

	// Success
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})

}
