package service

import (
	"github.com/sebastienmusso/infradatamgmt/config"
	"fmt"
	"strings"
	"net/http"
)

// Data structure for logging service.
type ServiceLogging struct {}


// Get action from the rooter in order to send to services logging.
func (ServiceLogging) GetAction(action string, data map[string]string) error {
	println("AGENT LOGGING sortie .....")
	fmt.Println(action, data)
	switch action {
		case "delete":
		case "create":
	}

	return nil
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

func removePipeline(name string) {
	auth := config.Config.Elasticsearch
	if config.Config.ElasticAuth {
		auth = config.Config.ElasticUser+":"+config.Config.ElasticPassword+"@"+config.Config.Elasticsearch
	}

	req, err := http.NewRequest("DELETE", auth+"/_ingest/pipeline/"+name+"?pretty", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}