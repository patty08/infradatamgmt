package service

import (
	"fmt"
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"os"
)

// configuration directory
const metricDirOriginal string = "./rooter/configuration/metricbeat/conf/original/"
const metricDirCustom string = "./rooter/configuration/metricbeat/conf/custom/"

// Data strcuture for metrics service.
type ServiceMetrics struct {}


// Get action from the rooter in order to send to services metrics.
func (ServiceMetrics) GetAction(action string, data map[string]string, client chan *ClientIN) error {
   switch action {
   case "stop":
	   detachMetricConfiguration(data)
   case "delete":
	   detachMetricConfiguration(data)
   case "start":
	   setConfigMetricServices(data["image"], data["application_type"], data["id"], data["ip"])
   }
   return nil
}

// Formatted the host id to string value:
//	hosts: ["ID"].
func formatMetricHostName(hostIP string) string{
	if hostIP == "" {
		fmt.Println("host name is empty")
	}
	name := bytes.NewBufferString(  "hosts: [\""+hostIP)
	name.WriteString("\"]")
	return name.String()
}

// Copy configuration file to /usr/share/metricbeat/custom/, open the file and setidConfiguration.
// 	image is the name of the image docker hub.
// 	Application_type must be set by user for setting configuration processors,
// 	if not nomconf is checked for matching similar file configuration.
// 	id is the id of the container.
func setConfigMetricServices(image string, application_type string, id string, ip string)  {
	if application_type == "" {
		application_type = image
	}
	application_type = strings.ToLower(application_type)
	// check if configuration exist
	println(application_type)

	// get file name from name agent: name.yml
	agentName := formatNameConfig(application_type)

	// check if the configuration is available in host
	file, err := os.OpenFile(metricDirOriginal +string(agentName),0,777)
	if err != nil {
		fmt.Println("src file not found:" +metricDirOriginal +agentName)
		println(err)
	}
	// close at last
	defer file.Close()

	CopyFile(file, metricDirCustom +application_type+"_"+id+".yml")

	// replace the id in the file configuration with the ids in data
	// find host: [" and replace to host : ["id"] in the custom configuration
	setMetricIpConfiguration(id, application_type, ip)
	println(">> configuration is set: "+string(agentName))
}

// Function to replace the line hosts in the configuration file to the id of the container in /custom/file.yml.
// ipContainer is the ip adress of the running container
// image is the base image of the container (application_type)
// ipContainer is the ip adress of the running container
func setMetricIpConfiguration(idContainer string, image string, ipContainer string) {
	// format host
	frmHost := formatMetricHostName(ipContainer)
	println(frmHost)

	// get file name from the name agent
	// check if the configuration is available in host
	fd, err := ioutil.ReadFile(metricDirCustom +image+"_"+idContainer+".yml")
	if err != nil {
		println("dest file not found:"+ metricDirCustom + image + "_" + idContainer+ ".yml")
		log.Fatalln(err)
	}

	lines := strings.Split(string(fd), "\n")
	for i, line := range lines {
		if strings.Contains(line, " hosts: [") {
			lines[i] = frmHost
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(metricDirCustom +string(image+"_"+idContainer+".yml"), []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// Delete configuration file when the container is removed or stopped.
func detachMetricConfiguration(data map[string]string){
	err:= os.Remove(metricDirCustom +data["application_type"]+"_"+data["id"]+".yml")
	if err != nil {
		fmt.Println(err)
		return
	}
}