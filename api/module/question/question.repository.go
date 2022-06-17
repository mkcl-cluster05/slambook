package question

import (
	"context"
	"slambook/api/middlewere"

	"go.mongodb.org/mongo-driver/mongo"
)

type QuestionRepository interface {
	create(context.Context, middlewere.User) (Question, error)
	read(context.Context, middlewere.User) (Question, error)
	update(context.Context, middlewere.User) (Question, error)
	delete(context.Context, middlewere.User) error
}

type questionRepository struct {
	Mongo *mongo.Client
}

func NewQuestionRepository(mongo *mongo.Client) QuestionRepository {
	return &questionRepository{
		Mongo: mongo,
	}
}

func (repo *questionRepository) create(ctx context.Context, user middlewere.User) (Question, error) {
	return Question{}, nil
}
func (repo *questionRepository) read(ctx context.Context, user middlewere.User) (Question, error) {
	return Question{}, nil
}
func (repo *questionRepository) update(ctx context.Context, user middlewere.User) (Question, error) {
	return Question{}, nil
}
func (repo *questionRepository) delete(ctx context.Context, user middlewere.User) error {
	return nil
}
