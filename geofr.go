package main

import (
	"github.com/edijon/geofr/departement"
	"github.com/edijon/geofr/format"
	"github.com/edijon/geofr/presenter"
	"github.com/spf13/cobra"
)

const apiUrl string = "https://geo.api.gouv.fr"
const defaultColumnSize uint = 25

func main() {
	var rootCmd = &cobra.Command{Use: "geofr"}
	var cmdDepartements = &cobra.Command{
		Use:   "departements [code département]",
		Short: "Affiche le découpage administratif des départements Français.",
		Long:  `Affiche le découpage administratif des départements Français.`,
		Args:  cobra.MinimumNArgs(0),
		Run:   departementsCmd,
	}
	rootCmd.AddCommand(cmdDepartements)
	rootCmd.Execute()
}

func departementsCmd(cmd *cobra.Command, args []string) {
	repository := departement.DepartementRepositoryREST{
		Url: apiUrl + "/departements",
	}
	output := presenter.StandardOutput{ColumnSize: defaultColumnSize}
	if len(args) < 1 {
		getDepartements(repository, output)
	} else {
		getDepartement(repository, output, args[0])
	}
}

func getDepartement(repository departement.DepartementRepository, output presenter.StandardOutput, code string) {
	departement := repository.Create(code)
	output.Write(format.Header(departement))
	output.Write(format.Row(departement))
}

func getDepartements(repository departement.DepartementRepository, output presenter.StandardOutput) {
	departements := repository.CreateAll()
	output.Write(format.Header(departements[0]))
	for _, departement := range departements {
		output.Write(format.Row(departement))
	}
}
