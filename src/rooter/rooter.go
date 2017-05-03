package rooter

import (
   "os"
   "time"
   "errors"
   "agent"
   "service"
   "fmt"
   "github.com/docker/docker/api/types"
)

//interface for agent setting
type AgentIn interface {
   AddEventListener(c chan *agent.InfoIN, who string) error
}

type sAgentIn struct {
   AgentIn AgentIn
}

//interface for service out settings
type ServiceOut interface {
   SetAction(action string, data map[string]string) error
}

//interface for service docker settings
type ServiceDocker interface {
   SetAction(action string, data map[string]string) error
}

//interface for service logging settings
type ServiceLogging interface {
   SetAction(action string, data map[string]string) error
}

//interface for service metrics settings
type ServiceMetrics interface {
   SetAction(action string, data map[string]string) error
}

type sServiceOut struct {
   aServiceOut ServiceOut
}

type sServiceDocker struct {
   aServiceOut ServiceDocker
}

type sServiceLogging struct {
   aServiceOut ServiceDocker
}

type sServiceMetrics struct {
   aServiceOut ServiceDocker
}

// Exit listening and close stream
func closeListener()  {
   println("Exit")
   os.Exit(0)
}

func process(i *agent.InfoIN) {
   fmt.Println("recupéré: ", i.Action,i.Services, i.Data)
   if i == nil {
	  errors.New("errors....")
   }

   for k := 0 ; k <= len(i.Services)-1; k++{
	  switch i.Services[k]{
	  case "STDOUT":
		 {
			agent := sServiceOut{service.ServiceStdout{}}
			go agent.aServiceOut.SetAction(i.Action, i.Data)
		 }
	  case "DOCKER":
		 {
			agent := sServiceDocker{service.ServiceDocker{}}
			go agent.aServiceOut.SetAction(i.Action, i.Data)
		 }
	  case "LOGGING":
		 {
			agent := sServiceLogging{service.ServiceLogging{}}
			go agent.aServiceOut.SetAction(i.Action, i.Data)
		 }
	  case "METRICS":
		 {
			agent := sServiceMetrics{service.ServiceMetrics{}}
			go agent.aServiceOut.SetAction(i.Action, i.Data)
		 }
	  }
   }
   fmt.Println("fin process...")
}


// start agent and open channels in and out stream
// input channel an listen to the structure value stream
// initialise des channels d'ecoute puis recupère les données du channel, traite les données inspecter puis envoi la configuration en sortie
func Start() {
   // open input channel and listening
   listener := make(chan *agent.InfoIN)
   // START ALL AGENTS

   agentDocker := agent.AgentDocker{}
   a := sAgentIn{agentDocker}
   go a.AgentIn.AddEventListener(listener, "unix:///var/run/docker.sock")
   for {
	  go process(<- listener)
	  time.Sleep (time.Second * 1)
   }
   closeListener()
   fmt.Println("fin...")
}