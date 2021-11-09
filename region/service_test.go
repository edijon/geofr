package region

import "testing"

type fakeOutput struct{}

func (output fakeOutput) Write(row []string) {}

func TestGivenFakeOutputAndRegionRepositoryFakeWhenGetRegion(t *testing.T) {
	output := fakeOutput{}
	repository := regionRepositoryFake{}
	GetRegion(repository, output, "66172")
}

func TestGivenFakeOutputAndRegionRepositoryFakeWhenGetRegions(t *testing.T) {
	output := fakeOutput{}
	repository := regionRepositoryFake{}
	GetRegions(repository, output)
}
