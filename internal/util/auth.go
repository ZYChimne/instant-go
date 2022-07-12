package util

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var hmac = []byte("zychimne")

type CustomClaims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

func GenerateJwt(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		userID,
		jwt.RegisteredClaims{
			Issuer:    "instant-go",
			Subject:   "authentication",
			Audience:  jwt.ClaimStrings{"instant-vue"},
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})
	tokenString, err := token.SignedString(hmac)
	if err != nil {
		log.Fatal("jwt token generate error ", err.Error())
	}
	return tokenString
}

func VerifyJwt(tokenString string) (string, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmac, nil
		},
	)
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims.UserID, nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
