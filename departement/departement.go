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
	Create(code *string) Departement
}

func DepartementCreate(factory DepartementRepository, code *string) Departement {
	return factory.Create(code)
}

type DepartementRepositoryRest struct {
	Url string
}

func (api *DepartementRepositoryRest) Create(code *string) Departement {
	url := queryUrl(api, code)
	departement := Departement{}

	resp, err := http.Get(url)
	util.AssertErrIsNotNil(&err)

	body, err := util.ReadAllFromHttpResponse(resp)
	util.AssertErrIsNotNil(&err)

	err = setDepartementFromJson(&departement, &body)
	util.AssertErrIsNotNil(&err)

	return departement
}

func queryUrl(api *DepartementRepositoryRest, code *string) string {
	return api.Url + "/" + *code + "/"
}

func setDepartementFromJson(departement *Departement, json *[]byte) error {
	return util.UnmarshalJsonFromBytes(json, &departement)
}
