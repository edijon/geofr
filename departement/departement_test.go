package departement

import (
	"testing"
)

const api string = "https://geo.api.gouv.fr/departements"

func TestGivenUrlWhenDepartementRepositoryRESTThenSet(t *testing.T) {
	repository := DepartementRepositoryREST{Url: "test"}
	if repository.Url != "test" {
		t.Errorf("DepartementRepositoryREST not set")
	}
}

type departementRepositoryFake struct{}

func (repository departementRepositoryFake) Create(code string) (Departement, error) {
	return Departement{Code: "code", Nom: "nom", CodeRegion: "coderegion"}, nil
}

func TestGivenFakeCodeAndRepositoryWhenCreateThenGetDepartement(t *testing.T) {
	repository := departementRepositoryFake{}
	code := ""
	departement, _ := repository.Create(code)
	if departement.Code == "" || departement.Nom == "" || departement.CodeRegion == "" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode66WhenCreateThenGetDepartement(t *testing.T) {
	code := "66"
	departement := Departement{Code: "00", Nom: "", CodeRegion: ""}
	repository := DepartementRepositoryREST{Url: api}
	departement, _ = repository.Create(code)
	if departement.Code != "66" || departement.Nom != "Pyrénées-Orientales" || departement.CodeRegion != "76" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode75WhenCreateThenGetDepartement(t *testing.T) {
	code := "75"
	departement := Departement{Code: "00", Nom: "", CodeRegion: ""}
	repository := DepartementRepositoryREST{Url: api}
	departement, _ = repository.Create(code)
	if departement.Code != "75" || departement.Nom != "Paris" || departement.CodeRegion != "11" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode00WhenCreateThenErr(t *testing.T) {
	code := "00"
	repository := DepartementRepositoryREST{Url: api}
	_, err := repository.Create(code)
	if err == nil {
		t.Errorf("Did not raise an error.")
	}
}

func TestGivenCodeAbWhenCreateThenErr(t *testing.T) {
	code := "ab"
	repository := DepartementRepositoryREST{Url: api}
	_, err := repository.Create(code)
	if err == nil {
		t.Errorf("Did not raise an error.")
	}
}

func TestGivendepartementRepositoryFakeWhenCreateAllThenSliceOfDepartements(t *testing.T) {
	repository := departementRepositoryFake{}
	departements, _ := repository.CreateAll()
	t.Logf("Got %q.", departements)
}

func (repository departementRepositoryFake) CreateAll() ([]Departement, error) {
	return []Departement{
		{Code: "code", Nom: "nom", CodeRegion: "coderegion"},
		{Code: "codeB", Nom: "nomB", CodeRegion: "coderegionB"},
	}, nil
}

func TestGivendepartementRepositoryRESTWhenCreateAllThenSliceOfDepartements(t *testing.T) {
	repository := DepartementRepositoryREST{Url: api}
	departements, _ := repository.CreateAll()
	t.Logf("Got %q.", departements)
}
