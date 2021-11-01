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

type DepartementRepositoryFake struct{}

func (repository DepartementRepositoryFake) Create(code string) Departement {
	return Departement{Code: "code", Nom: "nom", CodeRegion: "coderegion"}
}

func TestGivenFakeCodeAndRepositoryWhenCreateThenGetDepartement(t *testing.T) {
	repository := DepartementRepositoryFake{}
	code := ""
	departement := repository.Create(code)
	if departement.Code == "" || departement.Nom == "" || departement.CodeRegion == "" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode66WhenCreateThenGetDepartement(t *testing.T) {
	code := "66"
	departement := Departement{Code: "00", Nom: "", CodeRegion: ""}
	repository := DepartementRepositoryREST{Url: api}
	departement = DepartementCreate(&repository, code)
	if departement.Code != "66" || departement.Nom != "Pyrénées-Orientales" || departement.CodeRegion != "76" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode75WhenCreateThenGetDepartement(t *testing.T) {
	code := "75"
	departement := Departement{Code: "00", Nom: "", CodeRegion: ""}
	repository := DepartementRepositoryREST{Url: api}
	departement = DepartementCreate(&repository, code)
	if departement.Code != "75" || departement.Nom != "Paris" || departement.CodeRegion != "11" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode00WhenCreateThenErr(t *testing.T) {
	code := "00"
	repository := DepartementRepositoryREST{Url: api}
	defer func() { recover() }()
	DepartementCreate(&repository, code)
	t.Errorf("Did not panic.")
}
