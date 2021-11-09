package region

import (
	"github.com/edijon/geofr/format"
	"github.com/edijon/geofr/presenter"
)

func GetRegion(repository RegionRepository, output presenter.Output, code string) {
	reg, err := repository.Create(code)
	output.Write(format.Header(Region{}))
	if err == nil {
		output.Write(format.Row(reg))
	}
}

func GetRegions(repository RegionRepository, output presenter.Output) {
	regs, err := repository.CreateAll()
	output.Write(format.Header(Region{}))
	if err == nil {
		writeRegionRows(regs, output)
	}
}

func writeRegionRows(regs []Region, output presenter.Output) {
	for _, reg := range regs {
		output.Write(format.Row(reg))
	}
}
