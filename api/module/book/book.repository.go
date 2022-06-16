package book

import (
	"context"
	"slambook/api/middlewere"
	"time"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBName  = "slambook"
	ColName = "books"
)

type BookRepository interface {
	create(context.Context, middlewere.User, BookDTO) (Book, error)
	readAll(context.Context, middlewere.User) ([]Book, error)
	read(context.Context, middlewere.User, string) (Book, error)
	update(context.Context, middlewere.User, string, BookDTO) (Book, error)
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

	bookCollection := repo.Mongo.Database(DBName).Collection(ColName)

	book := Book{
		BookId:      ksuid.New().String(),
		AuthId:      user.AuthId,
		Name:        bookDTO.Name,
		Description: bookDTO.Description,
		IsDeleted:   false,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
	}

	_, err := bookCollection.InsertOne(ctx, book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}
func (repo *bookRepository) readAll(ctx context.Context, user middlewere.User) ([]Book, error) {

	bookCollection := repo.Mongo.Database(DBName).Collection(ColName)

	var books []Book

	query := bson.M{
		"authId":    user.AuthId,
		"isDeleted": false,
	}

	cursor, err := bookCollection.Find(ctx, query)

	if err != nil {
		return []Book{}, err
	}

	if err = cursor.All(ctx, &books); err != nil {
		return []Book{}, err
	}

	return books, nil
}
func (repo *bookRepository) read(ctx context.Context, user middlewere.User, bookId string) (Book, error) {

	bookCollection := repo.Mongo.Database(DBName).Collection(ColName)

	var book Book

	query := bson.M{
		"authId":    user.AuthId,
		"bookId":    book.BookId,
		"isDeleted": false,
	}

	err := bookCollection.FindOne(ctx, query).Decode(&book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}
func (repo *bookRepository) update(ctx context.Context, user middlewere.User, bookId string, bookDTO BookDTO) (Book, error) {

	bookCollection := repo.Mongo.Database(DBName).Collection(ColName)

	var book Book

	query := bson.M{
		"authId":    user.AuthId,
		"bookId":    book.BookId,
		"isDeleted": false,
	}

	updateQuery := bson.M{
		"name":        bookDTO.Name,
		"description": bookDTO.Description,
		"updatedAt":   time.Now().Unix(),
	}

	err := bookCollection.FindOneAndUpdate(ctx, query, updateQuery).Decode(&book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}
func (repo *bookRepository) delete(ctx context.Context, user middlewere.User, bookId string) error {

	bookCollection := repo.Mongo.Database(DBName).Collection(ColName)

	var book Book

	query := bson.M{
		"authId":    user.AuthId,
		"bookId":    book.BookId,
		"isDeleted": false,
	}

	deleteQuery := bson.M{
		"isDeleted": true,
		"updatedAt": time.Now().Unix(),
	}

	err := bookCollection.FindOneAndUpdate(ctx, query, deleteQuery).Decode(&book)

	if err != nil {
		return err
	}

	return nil

}
