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
	Create(code string) Departement
}

type DepartementRepositoryREST struct {
	Url string
}

func (api DepartementRepositoryREST) Create(code string) Departement {
	// REST implementation of the DepartementRepository interface.
	url := queryUrl(api, code)
	departement := Departement{}

	resp, err := http.Get(url)
	util.AssertErrIsNotNil(err)

	body, err := util.ReadAllFromHttpResponse(resp)
	util.AssertErrIsNotNil(err)

	err = util.UnmarshalJsonFromBytes(body, &departement)
	util.AssertErrIsNotNil(err)

	return departement
}

func queryUrl(api DepartementRepositoryREST, code string) string {
	// Builds a query url for target REST API depending of departement code.
	return api.Url + "/" + code + "/"
}
