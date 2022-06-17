package question

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func QuetionTypeValidation(v *validator.Validate,
	topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value,
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	params string) bool {

	fmt.Printf("topStruct %v  currentStructOrField %v  field %v  fieldType %v  fieldKind %v  params %v  ",
		topStruct,
		currentStructOrField,
		field,
		fieldType,
		fieldKind,
		params,
	)

	return true

}
