package middlewere

import (
	"fmt"
	"net/http"
	r "slambook/utils/response"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const jwtSecret = "jwtSecret"

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
		return []byte(jwtSecret), nil
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
		ctx.Next()
	}
}
