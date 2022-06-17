package question

import (
	"context"
	"slambook/api/middlewere"
	"time"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBName  = "slambook"
	ColName = "question"
)

type QuestionRepository interface {
	create(context.Context, middlewere.User, QuestionParam, QuestionDTO) (Question, error)
	read(context.Context, middlewere.User) (Question, error)
	update(context.Context, middlewere.User, QuestionParam, QuestionDTO) (Question, error)
	delete(context.Context, middlewere.User, QuestionParam, QuestionDTO) error
}

type questionRepository struct {
	Mongo *mongo.Client
}

func NewQuestionRepository(mongo *mongo.Client) QuestionRepository {
	return &questionRepository{
		Mongo: mongo,
	}
}

func (repo *questionRepository) create(ctx context.Context, user middlewere.User, questionParam QuestionParam, questionDTO QuestionDTO) (Question, error) {

	questionCollections := repo.Mongo.Database(DBName).Collection(ColName)

	question := Question{
		QuestionId:   ksuid.New().String(),
		BookId:       questionParam.BookId,
		AuthId:       user.AuthId,
		Title:        questionDTO.Title,
		QuestionType: questionDTO.QuestionType,
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    time.Now().Unix(),
		IsDeleted:    false,
	}

	_, err := questionCollections.InsertOne(ctx, question)

	if err != nil {
		return Question{}, err
	}

	return question, nil
}
func (repo *questionRepository) read(ctx context.Context, user middlewere.User) (Question, error) {
	return Question{}, nil
}
func (repo *questionRepository) update(ctx context.Context, user middlewere.User, questionParam QuestionParam, questionDTO QuestionDTO) (Question, error) {
	return Question{}, nil
}
func (repo *questionRepository) delete(ctx context.Context, user middlewere.User, questionParam QuestionParam, questionDTO QuestionDTO) error {
	return nil
}
