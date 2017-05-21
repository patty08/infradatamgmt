package cmd

import (
	//"github.com/sebastienmusso/infradatamgmt/rooter"
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var elasticsearch string
var kibana string
var agent []string
var client string

func Start() {
	var RootCmd = &cobra.Command{
		Use:   "surikator",
		Short: "short description",
		Long: `A verry long description, verry long`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(agent)
			//rooter.Start()
		},
	}

	RootCmd.PersistentFlags().StringVarP(&elasticsearch, "elasticsearch", "e", "127.0.0.1:9200", "set elastisearch localisation")
	RootCmd.PersistentFlags().StringVarP(&kibana, "kibana", "k", "127.0.0.1:5601", "set kibana localisation")
	RootCmd.PersistentFlags().StringSliceVarP(&agent, "agent", "a", []string{"docker"}, "set starting agent")
	RootCmd.PersistentFlags().StringVarP(&client, "client", "c", "docker", "set default client")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
