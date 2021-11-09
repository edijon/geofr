package region

import (
	"github.com/edijon/geofr/format"
	"github.com/edijon/geofr/presenter"
)

func GetRegion(repository RegionRepository, output presenter.Output, code string) {
	comm, err := repository.Create(code)
	output.Write(format.Header(Region{}))
	if err == nil {
		output.Write(format.Row(comm))
	}
}

func GetRegions(repository RegionRepository, output presenter.Output) {
	comms, err := repository.CreateAll()
	output.Write(format.Header(Region{}))
	if err == nil {
		writeRegionRows(comms, output)
	}
}

func writeRegionRows(comms []Region, output presenter.Output) {
	for _, comm := range comms {
		output.Write(format.Row(comm))
	}
}
