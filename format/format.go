package format

import (
	"reflect"
	"strings"
)

func Header(aStruct interface{}) []string {
	// Get column names from struct field names.
	var metadata reflect.Type = reflect.TypeOf(aStruct)
	var length int = metadata.NumField()
	var result []string = make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = strings.ToUpper(metadata.Field(i).Name)
	}
	return result
}

func Row(aStruct interface{}) []string {
	// Get column values from struct field values.
	var data reflect.Value = reflect.ValueOf(aStruct)
	var length int = data.NumField()
	var result []string = make([]string, length)
	for i := 0; i < length; i++ {
		result[i] = data.Field(i).String()
	}
	return result
}
