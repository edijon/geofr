package region

import (
	"net/http"

	"github.com/edijon/geofr/util"
)

type RegionRepository interface {
	Create(code string) (Region, error)
	CreateAll() ([]Region, error)
}

type RegionRepositoryREST struct {
	Url string
}

func (api RegionRepositoryREST) Create(code string) (Region, error) {
	// REST implementation of the RegionRepository interface.
	url := queryUrl(api, code)
	region := Region{}
	defer func() { recover() }()
	resp, _ := http.Get(url)
	body, _ := util.ReadAllFromHttpResponse(resp)
	err := util.UnmarshalJsonFromBytes(body, &region)
	return region, err
}

func queryUrl(api RegionRepositoryREST, code string) string {
	// Builds a query url for target REST API depending of departement code.
	return api.Url + "/" + code + "/"
}

func (api RegionRepositoryREST) CreateAll() ([]Region, error) {
	// REST implementation of the RegionRepository interface.
	url := queryAllUrl(api)
	regions := make([]Region, 0)
	defer func() { recover() }()
	resp, _ := http.Get(url)
	body, _ := util.ReadAllFromHttpResponse(resp)
	err := util.UnmarshalJsonFromBytes(body, &regions)
	return regions, err
}

func queryAllUrl(api RegionRepositoryREST) string {
	// Builds a query url for target REST API depending of departement code.
	return api.Url + "/"
}
