package auth

import (
	"context"
	"time"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var (
	DBName  = "slambook"
	ColName = "auth"
)

type AuthRepository interface {
	checkUser(context.Context, string) bool
	register(context.Context, RegisterDTO) (Auth, error)
	login(context.Context, LoginDTO) (Auth, error)
	changePassword(context.Context, ChangePasswordDTO)
	forgotPassword(context.Context, ForgotPasswordDTO)
}

type authRepository struct {
	Mongo *mongo.Client
}

func NewAuthRepository(mongo *mongo.Client) AuthRepository {
	return &authRepository{
		Mongo: mongo,
	}
}

func HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
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

func (repo *authRepository) checkUser(ctx context.Context, email string) bool {

	authCollection := repo.Mongo.Database(DBName).Collection(ColName)

	query := bson.M{
		"email": email,
	}
	var auth Auth

	_ = authCollection.FindOne(ctx, query).Decode(&auth)

	return auth.AuthId != ""
}

func (repo *authRepository) register(ctx context.Context, registerDTO RegisterDTO) (Auth, error) {

	authCollection := repo.Mongo.Database(DBName).Collection(ColName)

	hashPassword, err := HashPassword(registerDTO.Password)

	if err != nil {
		return Auth{}, err
	}

	auth := Auth{
		AuthId:    ksuid.New().String(),
		Username:  registerDTO.Username,
		Email:     registerDTO.Email,
		Password:  hashPassword,
		Role:      "user",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	_, err = authCollection.InsertOne(ctx, auth)

	if err != nil {
		return Auth{}, err
	}

	return auth, nil

}

func (repo *authRepository) login(ctx context.Context, loginDTO LoginDTO) (Auth, error) {

	authCollection := repo.Mongo.Database(DBName).Collection(ColName)

	var auth Auth
	query := bson.M{
		"email": loginDTO.Email,
	}
	err := authCollection.FindOne(ctx, query).Decode(&auth)

	if err != nil {
		return Auth{}, err
	}

	if err = ComparePassword(auth.Password, loginDTO.Password); err != nil {
		return Auth{}, err
	}

	return auth, nil
}
func (repo *authRepository) changePassword(ctx context.Context, changePasswordDTO ChangePasswordDTO) {

}
func (repo *authRepository) forgotPassword(ctx context.Context, forgotPasswordDTO ForgotPasswordDTO) {

}
