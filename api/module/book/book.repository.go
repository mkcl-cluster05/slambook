package book

import (
	"context"
	"slambook/api/middlewere"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository interface {
	create(context.Context, middlewere.User, BookDTO) (Book, error)
	readAll(context.Context, middlewere.User) ([]Book, error)
	read(context.Context, middlewere.User, string) (Book, error)
	update(context.Context, middlewere.User, string) (Book, error)
	delete(context.Context, middlewere.User, string) error
}

type bookRepository struct {
	Mongo *mongo.Client
}

func NewBookRepository(mongo *mongo.Client) BookRepository {
	return &bookRepository{
		Mongo: mongo,
	}
}

func (repo *bookRepository) create(ctx context.Context, user middlewere.User, bookDTO BookDTO) (Book, error) {
	return Book{}, nil
}
func (repo *bookRepository) readAll(ctx context.Context, user middlewere.User) ([]Book, error) {
	return []Book{}, nil
}
func (repo *bookRepository) read(ctx context.Context, user middlewere.User, bookId string) (Book, error) {
	return Book{}, nil
}
func (repo *bookRepository) update(ctx context.Context, user middlewere.User, bookId string) (Book, error) {
	return Book{}, nil
}
func (repo *bookRepository) delete(ctx context.Context, user middlewere.User, bookId string) error {
	return nil

}
