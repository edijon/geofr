package format

import (
	"reflect"
	"strconv"
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
		result[i] = stringFromValue(data.Field(i))
	}
	return result
}

func stringFromValue(value reflect.Value) string {
	// Get string representation from value for known cases, else value.String()
	result := value.String()
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result = strconv.FormatInt(value.Int(), 10)
	case reflect.Slice:
		elems, _ := value.Interface().([]string)
		result = strings.Join(elems, ", ")
	}
	return result
}
