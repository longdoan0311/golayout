package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/3t-dev/golayout"
)

var (
	projName string
	modName  string
)

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&projName, "proj_name", "n", "", "Project name")
	genCmd.Flags().StringVarP(&modName, "module_name", "m", "", "Module name (in case of go module), default is proj_name")
	genCmd.MarkFlagRequired("proj_name")
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Golang project layout",
	Long:  `Generate Golang project layout`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start generating Golang project layout")
		normParam()
		golayout.Generate(golayout.ProjectOverall{
			ProjName: projName,
			ModName:  modName,
		})
	},
}

func normParam() {
	if modName == "" {
		modName = projName
	}
}
