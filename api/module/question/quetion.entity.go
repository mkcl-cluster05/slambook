package question

type Question struct {
	QuestionId   string `json:"questionId" bson:"questionId"`
	BookId       string `json:"bookId" bson:"bookId"`
	AuthId       string `json:"authId" bson:"authId"`
	Title        string `json:"title" bson:"title"`
	QuestionType string `json:"questionType" bson:"questionType"`
	CreatedAt    string `json:"createdAt" bson:"createdAt"`
	UpdatedAt    string `json:"updatedAt" bson:"updatedAt"`
	IsDeleted    string `json:"isDeleted" bson:"isDeleted"`
}

type QuestionDTO struct {
	Title        string `json:"title" binding:"required,min=2,max=750"`
	QuestionType string `json:"questionType" binding:"required,questionType"`
}
