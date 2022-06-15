package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBName  = "slambook"
	ColName = "auth"
)

type AuthRepository interface {
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

func (repo *authRepository) register(ctx context.Context, registerDTO RegisterDTO) (Auth, error) {

	authCollection := repo.Mongo.Database(DBName).Collection(ColName)

	_, err := authCollection.InsertOne(ctx, registerDTO)

	if err != nil {
		return Auth{}, err
	}

	auth := Auth{
		Username: registerDTO.Username,
		Email:    registerDTO.Email,
		AuthId:   "authId",
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
	return auth, nil
}
func (repo *authRepository) changePassword(ctx context.Context, changePasswordDTO ChangePasswordDTO) {

}
func (repo *authRepository) forgotPassword(ctx context.Context, forgotPasswordDTO ForgotPasswordDTO) {

}
