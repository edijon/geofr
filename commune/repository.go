package commune

import (
	"net/http"

	"github.com/edijon/geofr/util"
)

type CommuneRepository interface {
	Create(code string) (Commune, error)
	CreateAll() ([]Commune, error)
}

type CommuneRepositoryREST struct {
	Url string
}

func (api CommuneRepositoryREST) Create(code string) (Commune, error) {
	// REST implementation of the CommuneRepository interface.
	url := queryUrl(api, code)
	commune := Commune{}
	defer func() { recover() }()
	resp, _ := http.Get(url)
	body, _ := util.ReadAllFromHttpResponse(resp)
	err := util.UnmarshalJsonFromBytes(body, &commune)
	return commune, err
}

func queryUrl(api CommuneRepositoryREST, code string) string {
	// Builds a query url for target REST API depending of departement code.
	return api.Url + "/" + code + "/"
}

func (api CommuneRepositoryREST) CreateAll() ([]Commune, error) {
	// REST implementation of the CommuneRepository interface.
	url := queryAllUrl(api)
	communes := make([]Commune, 0)
	defer func() { recover() }()
	resp, _ := http.Get(url)
	body, _ := util.ReadAllFromHttpResponse(resp)
	err := util.UnmarshalJsonFromBytes(body, &communes)
	return communes, err
}

func queryAllUrl(api CommuneRepositoryREST) string {
	// Builds a query url for target REST API depending of departement code.
	return api.Url + "/"
}
