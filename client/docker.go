package client

import (
   "github.com/sebastienmusso/infradatamgmt/agent"
   "github.com/sebastienmusso/infradatamgmt/service"
   "io"
   "os"
   "strings"
   "src/github.com/pkg/errors"
   "fmt"
   "bytes"
   "log"
)

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
			go service.aServiceOut.SetAction(info.Action, info.Data)
		 }
	  case "logging":
		 {
			service := sServiceOut{service.ServiceLogging{}}
			go service.aServiceOut.SetAction(info.Action, info.Data)
		 }
	  case "metric":
		 {
			service := sServiceOut{service.ServiceMetrics{}}
			go service.aServiceOut.SetAction(info.Action, info.Data)
			// todo: setting configuration for processors
			setConfigServices(info.Data["image"], info.Data["sys"])
		 }
	  }
   }

   return nil
}

// copie la configuration demandée dans /usr/share/metricbeat/custom/
// args systeme = label system activate or not activate
func setConfigServices(nomconf string, systeme string)  {

   // todo : nom passé en label
   // todo : system.yml à effacer
   // check if configuration exist
   conf := hasConfiguration(nomconf)

   // get file name from name agent
   agentName := formatNameConfig(strings.ToLower(nomconf))

   // check if the configuration is available in host
   file, err := os.OpenFile("./rooter/configuration/metricbeat/conf/original/"+string(agentName),0,777)
   if err != nil {
	  fmt.Println("src file not found: ./configuration/metricbeat/conf/original/"+agentName)
	  log.Fatal(err)
   }
   fmt.Println(file)

   defer file.Close()

   // Set system configuration system when system label is on, -l sys=on
   if systeme == "on"{
	  fileSys := getSystemConfig()
	  CopyFile(fileSys, "./rooter/configuration/metricbeat/conf/custom/system.yml")
   }

   if conf != "" {
	  CopyFile(file, "./rooter/configuration/metricbeat/conf/custom/"+string(conf)+".yml")
   } else {
	  CopyFile(file, "./rooter/configuration/metricbeat/conf/custom/"+string(agentName))
   }
   println("CONFIGURATION set: "+string(agentName))
   file.Close()

}

// change agent name in parameters to file name (eg: docker -> docker.yml)
func formatNameConfig(a string) string {
   if a == "" {
	  fmt.Println("name corrupted ")
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

// check if the name has configuration file
func hasConfiguration(nomconf string) string{
   if nomconf == "" {
	  errors.New("input nomconf parameters null")
   }
   // input configurations available
   isConfig := []string{"apache","docker","mongoDB","mySQL","nginx","php","postgreSQL","redis"}
   var conf string
   for i := range isConfig{
	  if strings.HasPrefix(nomconf, isConfig[i]){
		 conf = isConfig[i]
		 fmt.Println("ConfName--> : "+conf)
	  }
   }

   return string(conf)
}

func getSystemConfig() io.Reader{
   file, err := os.OpenFile("./configuration/metricbeat/conf/original/system.yml",0,777)
   if err == nil {
	  log.Fatal(err)
   }
   return file
}