package service

import (
	"github.com/edijon/geofr/commune"
	"github.com/edijon/geofr/departement"
	"github.com/edijon/geofr/format"
	"github.com/edijon/geofr/presenter"
)

func GetDepartement(repository departement.DepartementRepository, output presenter.Output, code string) {
	dept, err := repository.Create(code)
	output.Write(format.Header(departement.Departement{}))
	if err == nil {
		output.Write(format.Row(dept))
	}
}

func GetDepartements(repository departement.DepartementRepository, output presenter.Output) {
	depts, err := repository.CreateAll()
	output.Write(format.Header(departement.Departement{}))
	if err == nil {
		writeDepartementRows(depts, output)
	}
}

func writeDepartementRows(depts []departement.Departement, output presenter.Output) {
	for _, dept := range depts {
		output.Write(format.Row(dept))
	}
}

func GetCommune(repository commune.CommuneRepository, output presenter.Output, code string) {
	comm, err := repository.Create(code)
	output.Write(format.Header(commune.Commune{}))
	if err == nil {
		output.Write(format.Row(comm))
	}
}

func GetCommunes(repository commune.CommuneRepository, output presenter.Output) {
	comms, err := repository.CreateAll()
	output.Write(format.Header(commune.Commune{}))
	if err == nil {
		writeCommuneRows(comms, output)
	}
}

func writeCommuneRows(comms []commune.Commune, output presenter.Output) {
	for _, comm := range comms {
		output.Write(format.Row(comm))
	}
}
