package main

import (
    "github.com/sebastienmusso/infradatamgmt/config"
    "github.com/sebastienmusso/infradatamgmt/rooter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"fmt"
	"os"
)

var RootCmd = &cobra.Command{
    Use:   "surikator",
    Short: "short description",
    Long: `A verry long description, verry long`,
    Run: func(cmd *cobra.Command, args []string) {
        rooter.Start()
    },
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of Surikator",
    Long:  `Print the version number of Surikator`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Surikator version : v0.3-alpha")
    },
}

func main() {
    RootCmd.AddCommand(versionCmd)
    RootCmd.PersistentFlags().StringVarP(&config.Config.Elasticsearch, "Elasticsearch", "e", "", "set elastisearch localisation")
    RootCmd.PersistentFlags().StringVarP(&config.Config.Kibana, "Kibana", "k", "", "set Kibana localisation")
    RootCmd.PersistentFlags().StringSliceVarP(&config.Config.Agent, "Agent", "a", nil, "set starting Agent")
    RootCmd.PersistentFlags().StringVarP(&config.Config.Client, "Client", "c", "", "set default Client")

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

    viper.BindPFlag("Elasticsearch", RootCmd.PersistentFlags().Lookup("Elasticsearch"))
    viper.BindPFlag("Kibana", RootCmd.PersistentFlags().Lookup("Kibana"))
    viper.BindPFlag("Agent", RootCmd.PersistentFlags().Lookup("Agent"))
    viper.BindPFlag("Client", RootCmd.PersistentFlags().Lookup("Client"))

    viper.SetDefault("Elasticsearch", "127.0.0.1:9200")
    viper.SetDefault("ElasticAuth", true)
    viper.SetDefault("ElasticUser", "elastic")
    viper.SetDefault("ElasticPassword", "changeme")

    viper.SetDefault("Kibana", "127.0.0.1:5601")
    viper.SetDefault("Agent", []string{"docker"})
    viper.SetDefault("Client", "docker")

    config.Config.Elasticsearch = viper.GetString("Elasticsearch")
    config.Config.ElasticAuth = viper.GetBool("ElasticAuth")
    config.Config.ElasticUser = viper.GetString("ElasticUser")
	config.Config.ElasticPassword = viper.GetString("ElasticPassword")
    config.Config.Kibana = viper.GetString("Kibana")
    config.Config.Agent = viper.GetStringSlice("Agent")
    config.Config.Client = viper.GetString("Client")
}