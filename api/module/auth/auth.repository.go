package auth

import (
	"context"
	"fmt"
	"slambook/api/middlewere"
	"time"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBName  = "slambook"
	ColName = "auth"
)

type AuthRepository interface {
	checkUser(context.Context, string) bool
	register(context.Context, RegisterDTO) (Auth, error)
	login(context.Context, LoginDTO) (Auth, error)
	changePassword(context.Context, middlewere.User, ChangePasswordDTO) (Auth, error)
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
		return Auth{}, fmt.Errorf("invalid credentails")
	}

	return auth, nil
}
func (repo *authRepository) changePassword(ctx context.Context, user middlewere.User, changePasswordDTO ChangePasswordDTO) (Auth, error) {

	authCollection := repo.Mongo.Database(DBName).Collection(ColName)

	var auth Auth

	hashPassword, err := HashPassword(changePasswordDTO.NewPassword)

	if err != nil {
		return Auth{}, err
	}

	query := bson.M{
		"authId": user.AuthId,
	}
	update := bson.M{
		"$set": bson.M{
			"password":  hashPassword,
			"updatedAt": time.Now().Unix(),
		},
	}
	err = authCollection.FindOneAndUpdate(ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&auth)
	if err != nil {
		return Auth{}, nil
	}

	return auth, nil

}
func (repo *authRepository) forgotPassword(ctx context.Context, forgotPasswordDTO ForgotPasswordDTO) {

}
