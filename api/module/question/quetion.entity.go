package question

type QuetionType int8

const (
	MCQ QuetionType = iota
	Descriptive
	Image
	Video
	Audio
	GIF
)

func (qt QuetionType) String() string {
	switch qt {
	case MCQ:
		return "MCQ"
	case Descriptive:
		return "Descriptive"
	case Image:
		return "Image"
	case Video:
		return "Video"
	case Audio:
		return "Audio"
	case GIF:
		return "GIF"
	}
	return "Unknown"
}

type Question struct {
	QuestionId  string      `json:"questionId" bson:"questionId"`
	BookId      string      `json:"bookId" bson:"bookId"`
	AuthId      string      `json:"authId" bson:"authId"`
	Title       string      `json:"title" bson:"title"`
	QuetionType QuetionType `json:"quetionType" bson:"quetionType"`
	CreatedAt   string      `json:"createdAt" bson:"createdAt"`
	UpdatedAt   string      `json:"updatedAt" bson:"updatedAt"`
	IsDeleted   string      `json:"isDeleted" bson:"isDeleted"`
}

type QuestionDTO struct {
	Title       string      `json:"title" binding:"required,min=2,max=750"`
	QuetionType QuetionType `json:"quetionType" binding:"required" validation:"quetionType"`
}
