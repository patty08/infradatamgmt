package cmd

import (
	"github.com/sebastienmusso/infradatamgmt/rooter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
	"os"
)

var Config = struct {
	elasticsearch string
	kibana string
	agent []string
	client string
}{}

var RootCmd = &cobra.Command{
	Use:   "surikator",
	Short: "short description",
	Long: `A verry long description, verry long`,
	Run: func(cmd *cobra.Command, args []string) {
		rooter.Start()
	},
}

func Start() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().StringVarP(&Config.elasticsearch, "elasticsearch", "e", "", "set elastisearch localisation")
	RootCmd.PersistentFlags().StringVarP(&Config.kibana, "kibana", "k", "", "set kibana localisation")
	RootCmd.PersistentFlags().StringSliceVarP(&Config.agent, "agent", "a", nil, "set starting agent")
	RootCmd.PersistentFlags().StringVarP(&Config.client, "client", "c", "", "set default client")

	loadConfigFile()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadConfigFile() {
	viper.SetConfigFile("./config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		_, e := os.Create("./config.yml")
		if e != nil {
			panic(e)
		}
	}

	viper.BindPFlag("elasticsearch", RootCmd.PersistentFlags().Lookup("elasticsearch"))
	viper.BindPFlag("kibana", RootCmd.PersistentFlags().Lookup("kibana"))
	viper.BindPFlag("agent", RootCmd.PersistentFlags().Lookup("agent"))
	viper.BindPFlag("client", RootCmd.PersistentFlags().Lookup("client"))

	viper.SetDefault("elasticsearch", "127.0.0.1:9200")
	viper.SetDefault("kibana", "127.0.0.1:5601")
	viper.SetDefault("agent", []string{"docker"})
	viper.SetDefault("client", "docker")

	Config.elasticsearch = viper.GetString("elasticsearch")
	Config.kibana = viper.GetString("kibana")
	Config.agent = viper.GetStringSlice("agent")
	Config.client = viper.GetString("client")

}