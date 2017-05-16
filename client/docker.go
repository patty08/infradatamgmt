package client

import (
   "github.com/sebastienmusso/infradatamgmt/agent"
   "github.com/sebastienmusso/infradatamgmt/service"
   "io"
   "os"
   "strings"
   "fmt"
   "bytes"
   "log"
   "io/ioutil"
)

const dirOriginal string = "./rooter/configuration/metricbeat/conf/original/"
const dirCustom string = "./rooter/configuration/metricbeat/conf/custom/"

type ClientDocker struct {}

// function choose the action to set for the agent services
// services is the list of services to activate
// data is all informations data send by the container
// todo: refactor setAction() for all services

func (ClientDocker) SetAction(info *agent.InfoIN) error {
   for k := 0 ; k <= len(info.Services)-1; k++{
	  switch info.Services[k]{
	  case "stdout":
		 {
			service := sServiceOut{service.ServiceStdout{}}
			go service.aServiceOut.GetAction(info.Action, info.Data)
		 }
	  case "logging":
		 {
			service := sServiceOut{service.ServiceLogging{}}
			go service.aServiceOut.GetAction(info.Action, info.Data)
		 }
	  case "metric":
		 {
			service := sServiceOut{service.ServiceMetrics{}}
			go service.aServiceOut.GetAction(info.Action, info.Data)
			// todo: setting configuration for processors
			setConfigServices(info.Data["image"], info.Data["application_type"], info.Data["id"])
		 }
	  }
   }

   return nil
}

// format the host id to string hosts: ["ID"]
func formatHostName(hostid string) string{
   if hostid == "" {
	  fmt.Println("host name is empty")
   }
   name := bytes.NewBufferString(  "hosts: [\""+hostid)
   name.WriteString("\"]")
   return name.String()
}

// function to replace host line in the configuration file to the id of the container in /custom/file.yml
func setidConfiguration(idContainer string, nomconf string) {

   // format host
   frmHost := formatHostName(idContainer)

   // get file name from the name agent
   // check if the configuration is available in host
   fd, err := ioutil.ReadFile(dirCustom +nomconf+"_"+idContainer+".yml")
   if err != nil {
	  println("dest file not found:"+ dirCustom + nomconf + "_" + idContainer+ ".yml")
	  log.Fatalln(err)
   }

   lines := strings.Split(string(fd), "\n")
   for i, line := range lines {
	  if strings.Contains(line, " hosts: [") {
		 lines[i] = frmHost
	  }
   }
   output := strings.Join(lines, "\n")
   err = ioutil.WriteFile(dirCustom +string(nomconf+"_"+idContainer+".yml"), []byte(output), 0644)
   if err != nil {
	  log.Fatalln(err)
   }
}

// copy file configuration to /usr/share/metricbeat/custom/, open the file and set host to container's id
// nomconf is the name of the image docker hub
// application_type must be set by user for setting configuration processors,
// if not nomconf is checked for matching similar file configuration
// id is the id of the container
func setConfigServices(image string, application_type string, id string)  {
   if application_type == "" {
	  application_type = image
   }
   application_type = strings.ToLower(application_type)
   // check if configuration exist
   println(application_type)

   // get file name from name agent: name.yml
   agentName := formatNameConfig(application_type)

   // check if the configuration is available in host
   file, err := os.OpenFile(dirOriginal +string(agentName),0,777)
   if err != nil {
	  fmt.Println("src file not found:" +dirOriginal +agentName)
	  log.Fatal(err)
   }
   // close at last
   defer file.Close()

   CopyFile(file, dirCustom +string(agentName)+"_"+id+".yml")
   setidConfiguration(id, string(agentName))

   println(">> configuration is set: "+string(agentName))

   // replace the id in the file configuration with the ids in data
   // find host: [" and replace to host : ["id"] in the custom configuration

   file.Close()
}

// change agent name in parameters to file name (eg: docker -> docker.yml)
func formatNameConfig(a string) string {
   if a == "" {
	  fmt.Println("configuration name corrupted ")
   }
   name := bytes.NewBufferString(a)
   name.WriteString(".yml")
   return name.String()
}

// copy file source to destination, in is the file to copy, dst is the path where the file must be paste
func CopyFile(in io.Reader, dst string) (err error) {

   // Does file already exist? Skip
   if _, err := os.Stat(dst); err == nil {
	  return nil
   }

   out, err := os.Create(dst)
   if err != nil {
	  fmt.Println("Error creating file", err)
	  return
   }

   defer func() {
	  cerr := out.Close()
	  if err == nil {
		 err = cerr
	  }
   }()
   var bytes int64
   if bytes, err = io.Copy(out, in); err != nil {
	  fmt.Println("io.Copy error")
	  return
   }
   fmt.Println(bytes)

   err = out.Sync()
   return

}


