package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/3t-dev/golayout"
)

var (
	filePath string
)

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Flags().StringVarP(&filePath, "file_path", "f", "", "Path to file need to view")
	testCmd.MarkFlagRequired("file_path")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Generate Golang project layout",
	Long:  `Generate Golang project layout`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start generating Golang project layout")

		s, err := golayout.TemplateBox.FindString(filePath)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	},
}
