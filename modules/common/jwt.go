package common

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetJWTToken(data map[string]interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data":   data,
		"claims": time.Now().Add(5 * time.Minute).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("secret"))
	return tokenString
}

func ParseTokenFromHeader(r *http.Request) {
	authHeader := r.Header.Get("authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		return
	}
	tokenString := strings.Split(authHeader, " ")
	r.Header.Set("user", tokenString[1])

	// if (req.user) {
	//     setHashWithOverWrite(req.user._id, { loginTime: new Date().toISOString(), lastSeen: new Date().toISOString() })
	// }
}
