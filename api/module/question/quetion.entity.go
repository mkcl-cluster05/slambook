package question

type Question struct {
	QuestionId   string `json:"questionId" bson:"questionId"`
	BookId       string `json:"bookId" bson:"bookId"`
	AuthId       string `json:"authId" bson:"authId"`
	Title        string `json:"title" bson:"title"`
	QuestionType string `json:"questionType" bson:"questionType"`
	CreatedAt    int64  `json:"createdAt" bson:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt" bson:"updatedAt"`
	IsDeleted    bool   `json:"isDeleted" bson:"isDeleted"`
}

type QuestionDTO struct {
	Title        string `json:"title" binding:"required,min=2,max=750"`
	QuestionType string `json:"questionType" binding:"required,questionType"`
}

type QuestionParam struct {
	BookId     string `uri:"bookId" binding:"required"`
	QuestionId string `uri:"questionId,omitempty"`
}
