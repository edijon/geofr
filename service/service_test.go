package service

import (
	"testing"

	"github.com/edijon/geofr/commune"
	"github.com/edijon/geofr/departement"
)

type fakeOutput struct{}

func (output fakeOutput) Write(row []string) {}

type departementRepositoryFake struct{}

func (repository departementRepositoryFake) Create(code string) (departement.Departement, error) {
	return departement.Departement{}, nil
}

func (repository departementRepositoryFake) CreateAll() ([]departement.Departement, error) {
	return []departement.Departement{{}}, nil
}

func TestGivenFakeOutputAndDepartementRepositoryFakeWhenGetDepartement(t *testing.T) {
	output := fakeOutput{}
	repository := departementRepositoryFake{}
	GetDepartement(repository, output, "66172")
}

func TestGivenFakeOutputAndDepartementRepositoryFakeWhenGetDepartements(t *testing.T) {
	output := fakeOutput{}
	repository := departementRepositoryFake{}
	GetDepartements(repository, output)
}

type communeRepositoryFake struct{}

func (repository communeRepositoryFake) Create(code string) (commune.Commune, error) {
	return commune.Commune{}, nil
}

func (repository communeRepositoryFake) CreateAll() ([]commune.Commune, error) {
	return []commune.Commune{{}}, nil
}

func TestGivenFakeOutputAndCommuneRepositoryFakeWhenGetCommune(t *testing.T) {
	output := fakeOutput{}
	repository := communeRepositoryFake{}
	GetCommune(repository, output, "66172")
}

func TestGivenFakeOutputAndCommuneRepositoryFakeWhenGetCommunes(t *testing.T) {
	output := fakeOutput{}
	repository := communeRepositoryFake{}
	GetCommunes(repository, output)
}
