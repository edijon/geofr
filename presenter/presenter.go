package presenter

import (
	"reflect"
	"strings"
)

func Header(aStruct interface{}) []string {
	var aStructMetaData reflect.Type = reflect.TypeOf(aStruct)
	var length int = aStructMetaData.NumField()
	var result []string = make([]string, length)
	for i := 0; i < length; i++ {
		field := strings.ToUpper(
			aStructMetaData.Field(i).Name)
		result[i] = field
	}
	return result
}
