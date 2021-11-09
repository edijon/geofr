package main

import (
	"github.com/edijon/geofr/commune"
	"github.com/edijon/geofr/departement"
	"github.com/edijon/geofr/presenter"
	"github.com/edijon/geofr/region"
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
	var cmdRegions = &cobra.Command{
		Use:   "regions [code région]",
		Short: "Affiche le découpage administratif des régions Françaises.",
		Long:  `Affiche le découpage administratif des régions Françaises.`,
		Args:  cobra.MinimumNArgs(0),
		Run:   regionsCmd,
	}
	rootCmd.AddCommand(cmdDepartements)
	rootCmd.AddCommand(cmdCommunes)
	rootCmd.AddCommand(cmdRegions)
	rootCmd.Execute()
}

func departementsCmd(cmd *cobra.Command, args []string) {
	repository := departement.DepartementRepositoryREST{
		Url: apiUrl + "/departements",
	}
	output := presenter.StandardOutput{ColumnSize: defaultColumnSize}
	if len(args) < 1 {
		departement.GetDepartements(repository, output)
	} else {
		departement.GetDepartement(repository, output, args[0])
	}
}

func communesCmd(cmd *cobra.Command, args []string) {
	repository := commune.CommuneRepositoryREST{
		Url: apiUrl + "/communes",
	}
	output := presenter.StandardOutput{ColumnSize: defaultColumnSize}
	if len(args) < 1 {
		commune.GetCommunes(repository, output)
	} else {
		commune.GetCommune(repository, output, args[0])
	}
}

func regionsCmd(cmd *cobra.Command, args []string) {
	repository := region.RegionRepositoryREST{
		Url: apiUrl + "/regions",
	}
	output := presenter.StandardOutput{ColumnSize: 40}
	if len(args) < 1 {
		region.GetRegions(repository, output)
	} else {
		region.GetRegion(repository, output, args[0])
	}
}
