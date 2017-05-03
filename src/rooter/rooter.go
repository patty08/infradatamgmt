package rooter

import (
   "os"
   "time"
   "errors"
   "agent"
   "client"
   "io/ioutil"
   "log"
   "fmt"
)

// *** Strategy Agent ***

// interface for agent setting
type AgentIn interface {
   AddEventListener(c chan *agent.InfoIN, who string) error
}
// structure for calling agent
type sAgentIn struct {
   AgentIn AgentIn
}

// *** End Strategy Agent ***
// *** Strategy client ***

//interface for client out settings
type ClientOut interface {
   SetAction(info *agent.InfoIN) error
}

// Structure for calling client
type sClientOut struct {
   aClientOut ClientOut
}

// *** End Strategy client ***

// function close listening
func closeListener()  {
   println("Exit")
   os.Exit(0)
}

func process(i *agent.InfoIN) {
   //fmt.Println("recupéré: ", i.Action,i.Services, i.Data)
   if i == nil {
	  errors.New("errors....")
   }
   //namefile := "service/config.txt"
   //parseConfig(namefile)

   fmt.Println(i.Data)
   switch i.Data["client"]{
   case "docker":
	  {
		///agent := sServiceDocker{client.ClientDocker{}}
		//go agent.aClientDocker.SetAction(i.Action,i.Services, i.Data)
		 agent := sClientOut{client.ClientDocker{}}
		 go agent.aClientOut.SetAction(i)
	  }
   default:
	  fmt.Println("aaa", i.Data)
	  agent := sClientOut{client.ClientDocker{}}
	  go agent.aClientOut.SetAction(i)
   }
}

func parseConfig(namefile string) {
   stream, err := ioutil.ReadFile(namefile)
   if err != nil {
	  log.Fatal(err)
   }
   lireFichier := string(stream)
   fmt.Println(lireFichier)
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
}