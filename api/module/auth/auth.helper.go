package auth

import (
	"slambook/utils/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

func generateJWT(auth Auth) (string, error) {

	var jwtSecret = []byte(config.AppConfig.JWTSecret)
	jwtToken := jwt.New(jwt.SigningMethodHS256)
	claims := jwtToken.Claims.(jwt.MapClaims)

	claims["authId"] = auth.AuthId
	claims["username"] = auth.Username
	claims["email"] = auth.Email
	claims["role"] = auth.Role
	claims["createdAt"] = auth.CreatedAt
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token, err := jwtToken.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return token, nil
}
