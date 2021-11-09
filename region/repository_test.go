package region

import "testing"

const api string = "https://geo.api.gouv.fr/regions"

func TestGivenUrlWhenRegionRepositoryRESTThenSet(t *testing.T) {
	repository := RegionRepositoryREST{Url: "test"}
	if repository.Url != "test" {
		t.Errorf("RegionRepositoryREST not set")
	}
}

type regionRepositoryFake struct{}

func (repository regionRepositoryFake) Create(code string) (Region, error) {
	return Region{
		Nom:  "Occitanie",
		Code: "76",
	}, nil
}

func TestGivenFakeCodeAndRepositoryWhenCreateThenGetRegion(t *testing.T) {
	repository := regionRepositoryFake{}
	code := ""
	region, _ := repository.Create(code)
	if region.Code == "" || region.Nom == "" {
		t.Errorf("got %q", region)
	}
}

func TestGivenCode66172WhenCreateThenGetSaintEsteve(t *testing.T) {
	code := "76"
	region := Region{Code: "00", Nom: ""}
	repository := RegionRepositoryREST{Url: api}
	region, _ = repository.Create(code)
	if region.Code != "76" || region.Nom != "Occitanie" {
		t.Errorf("got %q", region)
	}
}

func TestGivenCode75056WhenCreateThenGetParis(t *testing.T) {
	code := "11"
	region := Region{Code: "00", Nom: ""}
	repository := RegionRepositoryREST{Url: api}
	region, _ = repository.Create(code)
	if region.Code != "11" || region.Nom != "Île-de-France" {
		t.Errorf("got %q", region)
	}
}

func TestGivenCode00WhenCreateThenErr(t *testing.T) {
	code := "00"
	repository := RegionRepositoryREST{Url: api}
	_, err := repository.Create(code)
	if err == nil {
		t.Errorf("Did not raise an error.")
	}
}

func TestGivenCodeAbWhenCreateThenErr(t *testing.T) {
	code := "ab"
	repository := RegionRepositoryREST{Url: api}
	_, err := repository.Create(code)
	if err == nil {
		t.Errorf("Did not raise an error.")
	}
}

func TestGivenRegionRepositoryFakeWhenCreateAllThenSliceOfRegions(t *testing.T) {
	repository := regionRepositoryFake{}
	departements, _ := repository.CreateAll()
	t.Logf("Got %q.", departements)
}

func (repository regionRepositoryFake) CreateAll() ([]Region, error) {
	return []Region{
		{
			Nom:  "Occitanie",
			Code: "76"},
		{
			Nom:  "Occitanie",
			Code: "76"},
	}, nil
}

func TestGivenRegionRepositoryRESTWhenCreateAllThenSliceOfRegions(t *testing.T) {
	repository := RegionRepositoryREST{Url: api}
	regs, _ := repository.CreateAll()
	t.Logf("Got %q.", regs)
}
