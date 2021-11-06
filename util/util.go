package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadAllFromHttpResponse(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func UnmarshalJsonFromBytes(bytes []byte, iface interface{}) error {
	// Set "iface" structure according to "bytes" JSON value.
	return json.Unmarshal(bytes, iface)
}
