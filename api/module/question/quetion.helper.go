package question

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var QuestionType = map[string]string{
	"MCQ":         "MCQ",
	"Descriptive": "Descriptive",
	"Image":       "Image",
	"Video":       "Video",
	"Audio":       "Audio",
	"GIF":         "GIF",
}

var QuestionTypeValidation validator.Func = func(fl validator.FieldLevel) bool {
	qType := fmt.Sprintf("%s", fl.Field())
	_, ok := QuestionType[qType]
	return ok
}
