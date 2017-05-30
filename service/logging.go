package service

import (
	"github.com/sebastienmusso/infradatamgmt/config"
	"fmt"
	"strings"
	"net/http"
)

// Data structure for logging service.
type ServiceLogging struct {}

const loggingDirPipeline string = "./rooter/configuration/elasticsearch/pipeline/"

// Get action from the rooter in order to send to services logging.
func (ServiceLogging) GetAction(action string, data map[string]string) error {
	println("AGENT LOGGING : " + action)
	switch action {
		case "delete":
		case "create":
			deployLogging(data["application_type"])
	}

	return nil
}

func deployLogging(service string) {
	cfg := loadConfigFile("./service/logging.yml")

	// if needed deploy fileBeat
	app := cfg.GetStringMapStringSlice(service)
	if len(app) > 0 {
		// deploy pipeline elasticsearch for filebeat
		for _, grock := range app["grock"] {
			sendPipeline(grock, loggingDirPipeline + grock)
		}
	}
}

func sendPipeline(name string, cfg string) {
	auth := config.Config.Elasticsearch
	if config.Config.ElasticAuth {
		auth = config.Config.ElasticUser+":"+config.Config.ElasticPassword+"@"+config.Config.Elasticsearch
	}

	body := strings.NewReader(cfg)
	req, err := http.NewRequest("PUT", auth+"/_ingest/pipeline/"+name+"?pretty", body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}