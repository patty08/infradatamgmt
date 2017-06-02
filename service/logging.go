package service

import (
	"github.com/sebastienmusso/infradatamgmt/config"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"log"
	"os"
)

// Data structure for logging service.
type ServiceLogging struct {}

const loggingDirPipeline string = "./rooter/configuration/elasticsearch/pipeline/"
const loggingDirCfgFilebeatOriginal string = "./rooter/configuration/filebeat/original/"
const loggingDirCfgFilebeat string = "./rooter/configuration/filebeat/custom/"

// Get action from the rooter in order to send to services logging.
func (ServiceLogging) GetAction(action string, data map[string]string, client chan *ClientIN) error {
	println("AGENT LOGGING : " + action)
	switch action {
		case "stop":
			if len(data["application_type"]) == 0 {
				data["application_type"] = data["image"]
			}
			removeLogging(data["application_type"], data, client)
		case "start":
			if len(data["application_type"]) == 0 {
				data["application_type"] = data["image"]
			}
			deployLogging(data["application_type"], data, client)
	}

	//close socket
	infos := &ClientIN{}
	infos.Action = "end"
	client <- infos

	return nil
}

func deployLogging(service string, data map[string]string, client chan *ClientIN) {
	cfg := loadConfigFile("./service/logging.yml")

	// if needed deploy fileBeat
	app := cfg.GetStringMapStringSlice(service)
	if len(app) > 0 {
		// deploy pipeline elasticsearch for filebeat
		for _, grock := range app["grock"] {
			sendPipeline(grock, loggingDirPipeline + grock + ".json")
		}

		infos := &ClientIN{}
		infos.Action = "start"
		infos.Data = map[string]string{}

		infos.Data["image"] = "docker.elastic.co/beats/filebeat:5.2.1"
		infos.Data["name"] = "filebeat_" + data["id"]
		infos.Data["user"] = "root"
		infos.Data["net-host"] = "true"

		infos.Data["who_id"] = data["id"]
		infos.Data["who_name"] = data["name"]
		infos.Data["volume_from"] = data["name"] + ":ro"
																																																			infos.Data["volume"] = "true"
		infos.Data["volume_src"], _ = filepath.Abs(filepath.Dir(loggingDirCfgFilebeat))
		infos.Data["volume_src"] += "/" + service + "_" + data["id"] + ".yml"
		fmt.Println(infos.Data["volume_src"])
		infos.Data["volume_container"] = "/usr/share/filebeat/filebeat.yml"

		dst := loggingDirCfgFilebeat + service + "_" + data["id"] + ".yml"
		CopyFilePath(loggingDirCfgFilebeatOriginal + service + ".yml", dst)
		setLoggingElkConfiguration(dst)

		client <- infos
	}
}

func removeLogging(service string, data map[string]string, client chan *ClientIN) {
	cfg := loadConfigFile("./service/logging.yml")

	// if had need deploy filebeat
	app := cfg.GetStringMapStringSlice(service)
	if len(app) > 0 {
		infos := &ClientIN{}
		infos.Action = "stop"
		infos.Data = map[string]string{}

		infos.Data["name"] = "filebeat_"+data["id"]
		client <- infos

		os.Remove(loggingDirCfgFilebeat + service + "_" + data["id"] + ".yml")

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
	//fmt.Println(name, body) // show pipeline
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

func setLoggingElkConfiguration(fileName string) {
	fd, err := ioutil.ReadFile(fileName)
	if err != nil {
		println("dest file not found:"+ fileName)
		log.Fatalln(err)
	}

	lines := strings.Split(string(fd), "\n")
	for i, line := range lines {
		if strings.Contains(line, "elk-ip") {
			lines[i] = "  hosts: [\"" + config.Config.Elasticsearch + "\"]"
			if config.Config.ElasticAuth {
				lines = append(lines, "  username: \"" + config.Config.ElasticUser + "\"")
				lines = append(lines, "  password: \"" + config.Config.ElasticPassword + "\"")
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}