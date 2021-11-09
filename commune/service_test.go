package commune

import "testing"

type fakeOutput struct{}

func (output fakeOutput) Write(row []string) {}

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
