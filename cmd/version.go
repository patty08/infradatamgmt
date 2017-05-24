package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Surikator",
	Long:  `Print the version number of Surikator`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Surikator version : v0.3-alpha")
	},
}