package departement

import "testing"

type fakeOutput struct{}

func (output fakeOutput) Write(row []string) {}

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
