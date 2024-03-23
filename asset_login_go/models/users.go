package models

type User struct{
	user_id string `json:user_id`
	user_pw string `json:user_pw`
	nickname string `json:nickname`
	email string `json:email`
}