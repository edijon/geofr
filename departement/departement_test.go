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

func (repository departementRepositoryFake) Create(code string) Departement {
	return Departement{Code: "code", Nom: "nom", CodeRegion: "coderegion"}
}

func TestGivenFakeCodeAndRepositoryWhenCreateThenGetDepartement(t *testing.T) {
	repository := departementRepositoryFake{}
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
	departement = repository.Create(code)
	if departement.Code != "66" || departement.Nom != "Pyrénées-Orientales" || departement.CodeRegion != "76" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode75WhenCreateThenGetDepartement(t *testing.T) {
	code := "75"
	departement := Departement{Code: "00", Nom: "", CodeRegion: ""}
	repository := DepartementRepositoryREST{Url: api}
	departement = repository.Create(code)
	if departement.Code != "75" || departement.Nom != "Paris" || departement.CodeRegion != "11" {
		t.Errorf("got %q", departement)
	}
}

func TestGivenCode00WhenCreateThenErr(t *testing.T) {
	code := "00"
	repository := DepartementRepositoryREST{Url: api}
	defer func() { recover() }()
	repository.Create(code)
	t.Errorf("Did not panic.")
}

func TestGivendepartementRepositoryFakeWhenCreateAllThenSliceOfDepartements(t *testing.T) {
	repository := departementRepositoryFake{}
	var departements []Departement = repository.CreateAll()
	t.Logf("Got %q.", departements)
}

func (repository departementRepositoryFake) CreateAll() []Departement {
	return []Departement{
		{Code: "code", Nom: "nom", CodeRegion: "coderegion"},
		{Code: "codeB", Nom: "nomB", CodeRegion: "coderegionB"},
	}
}

func TestGivendepartementRepositoryRESTWhenCreateAllThenSliceOfDepartements(t *testing.T) {
	repository := DepartementRepositoryREST{Url: api}
	var departements []Departement = repository.CreateAll()
	t.Logf("Got %q.", departements)
}
