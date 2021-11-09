package commune

import (
	"github.com/edijon/geofr/format"
	"github.com/edijon/geofr/presenter"
)

func GetCommune(repository CommuneRepository, output presenter.Output, code string) {
	comm, err := repository.Create(code)
	output.Write(format.Header(Commune{}))
	if err == nil {
		output.Write(format.Row(comm))
	}
}

func GetCommunes(repository CommuneRepository, output presenter.Output) {
	comms, err := repository.CreateAll()
	output.Write(format.Header(Commune{}))
	if err == nil {
		writeCommuneRows(comms, output)
	}
}

func writeCommuneRows(comms []Commune, output presenter.Output) {
	for _, comm := range comms {
		output.Write(format.Row(comm))
	}
}
