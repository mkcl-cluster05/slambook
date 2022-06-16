package book

type Book struct {
	BookId      string `json:"bookId" bson:"bookId"`
	AuthId      string `json:"authId" bson:"authId"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	CreatedAt   int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt" bson:"updatedAt"`
	IsDeleted   bool   `json:"isDeleted"`
}

type BookDTO struct {
	Name        string `json:"name" binding:"required,min=2,max=50"`
	Description string `json:"description,omitempty"`
}
