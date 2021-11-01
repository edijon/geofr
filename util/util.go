package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AssertErrIsNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadAllFromHttpResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func UnmarshalJsonFromBytes(bytes []byte, iface interface{}) error {
	return json.Unmarshal(bytes, iface)
}
