package service

import (
	"github.com/sebastienmusso/infradatamgmt/config"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

// Data structure for logging service.
type ServiceLogging struct {}

const loggingDirPipeline string = "./rooter/configuration/elasticsearch/pipeline/"

// Get action from the rooter in order to send to services logging.
func (ServiceLogging) GetAction(action string, data map[string]string) error {
	println("AGENT LOGGING : " + action)
	switch action {
		case "stop":
			if len(data["application_type"]) == 0 {
				data["application_type"] = data["image"]
			}
			removeLogging(data["application_type"])
		case "start":
			if len(data["application_type"]) == 0 {
				data["application_type"] = data["image"]
			}
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
			sendPipeline(grock, loggingDirPipeline + grock + ".json")
		}

		// TODO : start filebeat
	}
}

func removeLogging(service string) {
	cfg := loadConfigFile("./service/logging.yml")

	// if had need deploy filebeat
	app := cfg.GetStringMapStringSlice(service)
	if len(app) > 0 {
		// TODO : Remove filebeat

	}
}

func sendPipeline(name string, cfg string) {
	auth := config.Config.Elasticsearch
	if config.Config.ElasticAuth {
		auth = config.Config.ElasticUser+":"+config.Config.ElasticPassword+"@"+config.Config.Elasticsearch
	}

	f, err := ioutil.ReadFile(cfg)
	if err != nil {
		fmt.Print(err)
	}

	body := strings.NewReader(string(f))
	fmt.Println(name, body)
	req, err := http.NewRequest("PUT", "http://"+auth+"/_ingest/pipeline/"+name+"?pretty", body)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}