package main

import (
	"github.com/edijon/geofr/commune"
	"github.com/edijon/geofr/departement"
	"github.com/edijon/geofr/presenter"
	"github.com/edijon/geofr/service"
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
	var cmdCommunes = &cobra.Command{
		Use:   "communes [code commune]",
		Short: "Affiche le découpage administratif des communes Françaises.",
		Long:  `Affiche le découpage administratif des communes Françaises.`,
		Args:  cobra.MinimumNArgs(0),
		Run:   communesCmd,
	}
	rootCmd.AddCommand(cmdDepartements)
	rootCmd.AddCommand(cmdCommunes)
	rootCmd.Execute()
}

func departementsCmd(cmd *cobra.Command, args []string) {
	repository := departement.DepartementRepositoryREST{
		Url: apiUrl + "/departements",
	}
	output := presenter.StandardOutput{ColumnSize: defaultColumnSize}
	if len(args) < 1 {
		service.GetDepartements(repository, output)
	} else {
		service.GetDepartement(repository, output, args[0])
	}
}

func communesCmd(cmd *cobra.Command, args []string) {
	repository := commune.CommuneRepositoryREST{
		Url: apiUrl + "/communes",
	}
	output := presenter.StandardOutput{ColumnSize: defaultColumnSize}
	if len(args) < 1 {
		service.GetCommunes(repository, output)
	} else {
		service.GetCommune(repository, output, args[0])
	}
}
