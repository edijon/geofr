package main

import (
	"github.com/edijon/geofr/departement"
	"github.com/edijon/geofr/format"
	"github.com/edijon/geofr/presenter"
	"github.com/spf13/cobra"
)

const ApiUrl string = "https://geo.api.gouv.fr"

func main() {
	var rootCmd = &cobra.Command{Use: "geofr"}
	var cmdDepartement = &cobra.Command{
		Use:   "departement [code département]",
		Short: "Affiche le découpage administratif d'un département Français.",
		Long:  `Affiche le découpage administratif d'un département Français.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   departementCmd,
	}
	rootCmd.AddCommand(cmdDepartement)
	rootCmd.Execute()
}

func departementCmd(cmd *cobra.Command, args []string) {
	repository := departement.DepartementRepositoryREST{
		Url: ApiUrl + "/departements",
	}
	departement := repository.Create(args[0])
	output := presenter.StandardOutput{ColumnSize: 30}
	output.Write(format.Header(departement))
	output.Write(format.Row(departement))
}
