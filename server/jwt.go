package server

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var AddKey = "2108A"

// GetJwtToken @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func GetJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetCheckJwtToken(tokenString string) (int, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		hmacSampleSecret := []byte(AddKey)
		return hmacSampleSecret, nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	f := claims["exp"].(float64)
	s := claims["payload"].(string)
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	log.Println(s)
	if time.Now().Unix() > int64(f) {
		return 0, errors.New("token失效")
	}

	return atoi, nil
}
