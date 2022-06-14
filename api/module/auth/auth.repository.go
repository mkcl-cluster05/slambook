package auth

import "go.mongodb.org/mongo-driver/mongo"

type AuthRepository interface {
	register()
	login()
	changePassword()
	forgotPassword()
}

type authRepository struct {
	Mongo *mongo.Client
}

func NewAuthRepository(mongo *mongo.Client) AuthRepository {
	return &authRepository{
		Mongo: mongo,
	}
}

func (repo *authRepository) register() {
}

func (repo *authRepository) login() {

}
func (repo *authRepository) changePassword() {

}
func (repo *authRepository) forgotPassword() {

}
