package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/3t-dev/golayout"
)

var (
	projName string
)

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&projName, "proj_name", "n", "", "Path to the JSON files directory")
	genCmd.MarkFlagRequired("proj_name")
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate Golang project layout",
	Long:  `Generate Golang project layout`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start generating Golang project layout")
		golayout.Generate(projName)
	},
}
