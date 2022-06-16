package middlewere

import (
	"fmt"
	"net/http"
	"slambook/utils/config"
	r "slambook/utils/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type User struct {
	AuthId   string `json:"authId,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.AppConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if claim := token.Claims; claim.Valid() != nil {
		return err
	}
	return nil
}

func GetUser(r *http.Request) (string, error) {
	token, err := VerifyToken(r)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		user := User{
			AuthId:   fmt.Sprintf("%s", claims["authId"]),
			Username: fmt.Sprintf("%s", claims["username"]),
			Email:    fmt.Sprintf("%s", claims["email"]),
			Role:     fmt.Sprintf("%s", claims["role"]),
		}

		reqUser, _ := json.Marshal(&user)

		return string(reqUser), nil
	}
	return "", err
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := TokenValid(ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, r.ErrorResponse{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
				Error:   err.Error(),
			})
			ctx.Abort()
			return
		}

		user, _ := GetUser(ctx.Request)
		ctx.Request.Header.Add("user", user)

		ctx.Next()
	}
}
