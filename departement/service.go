package departement

import (
	"github.com/edijon/geofr/format"
	"github.com/edijon/geofr/presenter"
)

func GetDepartement(repository DepartementRepository, output presenter.Output, code string) {
	dept, err := repository.Create(code)
	output.Write(format.Header(Departement{}))
	if err == nil {
		output.Write(format.Row(dept))
	}
}

func GetDepartements(repository DepartementRepository, output presenter.Output) {
	depts, err := repository.CreateAll()
	output.Write(format.Header(Departement{}))
	if err == nil {
		writeDepartementRows(depts, output)
	}
}

func writeDepartementRows(depts []Departement, output presenter.Output) {
	for _, dept := range depts {
		output.Write(format.Row(dept))
	}
}
