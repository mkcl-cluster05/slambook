package book

import (
	"context"
	"slambook/api/middlewere"
	"time"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	var books []Book = []Book{}

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

	query := bson.M{
		"authId":    user.AuthId,
		"bookId":    bookId,
		"isDeleted": false,
	}

	var book Book = Book{}
	err := bookCollection.FindOne(ctx, query).Decode(&book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}
func (repo *bookRepository) update(ctx context.Context, user middlewere.User, bookId string, bookDTO BookDTO) (Book, error) {

	bookCollection := repo.Mongo.Database(DBName).Collection(ColName)

	query := bson.M{
		"authId":    user.AuthId,
		"bookId":    bookId,
		"isDeleted": false,
	}

	updateQuery := bson.M{
		"$set": bson.M{
			"name":        bookDTO.Name,
			"description": bookDTO.Description,
			"updatedAt":   time.Now().Unix(),
		},
	}

	var book Book
	err := bookCollection.FindOneAndUpdate(ctx, query, updateQuery, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&book)

	if err != nil {
		return Book{}, err
	}

	return book, nil
}
func (repo *bookRepository) delete(ctx context.Context, user middlewere.User, bookId string) error {

	bookCollection := repo.Mongo.Database(DBName).Collection(ColName)

	query := bson.M{
		"authId":    user.AuthId,
		"bookId":    bookId,
		"isDeleted": false,
	}

	deleteQuery := bson.M{
		"$set": bson.M{
			"isDeleted": true,
			"updatedAt": time.Now().Unix(),
		},
	}

	_ = bookCollection.FindOneAndUpdate(ctx, query, deleteQuery)

	return nil

}
