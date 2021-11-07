package commune

import "testing"

const api string = "https://geo.api.gouv.fr/communes"

func TestGivenUrlWhenCommuneRepositoryRESTThenSet(t *testing.T) {
	repository := CommuneRepositoryREST{Url: "test"}
	if repository.Url != "test" {
		t.Errorf("CommuneRepositoryREST not set")
	}
}

type communeRepositoryFake struct{}

func (repository communeRepositoryFake) Create(code string) (Commune, error) {
	return Commune{
		Nom:             "Saint-Estève",
		Code:            "66172",
		CodeDepartement: "66",
		CodeRegion:      "76",
		CodesPostaux:    []string{"66240"},
		Population:      11841,
	}, nil
}

func TestGivenFakeCodeAndRepositoryWhenCreateThenGetCommune(t *testing.T) {
	repository := communeRepositoryFake{}
	code := ""
	commune, _ := repository.Create(code)
	if commune.Code == "" || commune.Nom == "" || commune.CodeRegion == "" {
		t.Errorf("got %q", commune)
	}
}

func TestGivenCode66172WhenCreateThenGetCommune(t *testing.T) {
	code := "66172"
	commune := Commune{Code: "00", Nom: "", CodeRegion: ""}
	repository := CommuneRepositoryREST{Url: api}
	commune, _ = repository.Create(code)
	if commune.Code != "66172" || commune.Nom != "Saint-Estève" || commune.CodeRegion != "76" {
		t.Errorf("got %q", commune)
	}
}
