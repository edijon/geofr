package departement

import (
	"net/http"

	"github.com/edijon/geofr/util"
)

type Departement struct {
	Nom        string `json:"nom"`
	Code       string `json:"code"`
	CodeRegion string `json:"codeRegion"`
}

type DepartementRepository interface {
	// This interface should be implemented for Departement operations.
	Create(code string) (Departement, error)
	CreateAll() ([]Departement, error)
}

type DepartementRepositoryREST struct {
	Url string
}

func (api DepartementRepositoryREST) Create(code string) (Departement, error) {
	// REST implementation of the DepartementRepository interface.
	url := queryUrl(api, code)
	departement := Departement{}
	defer func() { recover() }()
	resp, _ := http.Get(url)
	body, _ := util.ReadAllFromHttpResponse(resp)
	err := util.UnmarshalJsonFromBytes(body, &departement)
	return departement, err
}

func queryUrl(api DepartementRepositoryREST, code string) string {
	// Builds a query url for target REST API depending of departement code.
	return api.Url + "/" + code + "/"
}

func (api DepartementRepositoryREST) CreateAll() ([]Departement, error) {
	// REST implementation of the DepartementRepository interface.
	url := queryAllUrl(api)
	departements := make([]Departement, 0)
	defer func() { recover() }()
	resp, _ := http.Get(url)
	body, _ := util.ReadAllFromHttpResponse(resp)
	err := util.UnmarshalJsonFromBytes(body, &departements)
	return departements, err
}

func queryAllUrl(api DepartementRepositoryREST) string {
	// Builds a query url for target REST API depending of departement code.
	return api.Url + "/"
}
