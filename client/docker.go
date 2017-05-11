package client

import (
   "github.com/sebastienmusso/infradatamgmt/agent"
   "github.com/sebastienmusso/infradatamgmt/service"
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
		 }
	  }
   }

   return nil
}

func setConfigServices(nomconf string, idContainer string, volumeContainer string)  {

}