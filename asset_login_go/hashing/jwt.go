package hashing

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(Email string) (string, error){
	keySignUp := []byte(os.Getenv("SECRET_key"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	tk, err := token.SignedString(keySignUp)
	if err != nil{
		return "",err
	}
	return tk,nil
}