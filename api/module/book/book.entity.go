package book

type Book struct {
	BookId      string `json:"bookId" bson:"bookId"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	CreatedAt   string `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string `json:"updatedAt" bson:"updatedAt"`
}

type BookDTO struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Description string `json:"description" `
}
