package presenter

import (
	"reflect"
	"strings"
)

func Header(aStruct interface{}) string {
	reflectionType := reflect.TypeOf(aStruct)
	var fieldCount int = reflectionType.NumField()
	var result []string = make([]string, fieldCount)
	for i := 0; i < fieldCount; i++ {
		field := strings.ToUpper(
			reflectionType.Field(i).Name)
		result[i] = field
	}
	return strings.Join(result, "; ")
}
