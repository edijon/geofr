package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func fakeNilError() error {
	return nil
}

func TestGivenResponseWithBodyWhenReadAllFromHttpResponseThenOK(t *testing.T) {
	response := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("test")),
	}
	ReadAllFromHttpResponse(&response)
}

func TestGivenResponseWithNoBodyWhenReadAllFromHttpResponseThenPanic(t *testing.T) {
	defer func() { recover() }()
	response := http.Response{}
	ReadAllFromHttpResponse(&response)
	t.Errorf("Did not panic.")
}

func TestGivenJsonBytesWhenUnmarshalJsonFromBytesThenGetStructure(t *testing.T) {
	jsonBytes := []byte(`{"message": "test"}`)
	message := &jsonStruct{}
	UnmarshalJsonFromBytes(jsonBytes, &message)
	if message.Message != "test" {
		t.Errorf("Didn't parse json, got : %q instead of %q", message, jsonBytes)
	}
}

type jsonStruct struct {
	Message string `json:"message"`
}
