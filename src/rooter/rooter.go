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

//noinspection GoImportUsedAsName
func process(i *agent.InfoIN) {
   if i == nil {
	  errors.New("errors....")
   }
   // TODO: match configuration settings into input file
   //namefile := "service/config.txt"
   //parseConfig(namefile)

   fmt.Println(i.Data)
   switch i.Data["client"]{
   //noinspection GoImportUsedAsName
   case "docker":
	  {
		///agent := sServiceDocker{client.ClientDocker{}}
		//go agent.aClientDocker.SetAction(i.Action,i.Services, i.Data)
		 agent := sClientOut{client.ClientDocker{}}
		 go agent.aClientOut.SetAction(i)
	  }
   //noinspection GoImportUsedAsName
   default:
	  fmt.Println("aaa", i.Data)
	  agent := sClientOut{client.ClientDocker{}}
	  go agent.aClientOut.SetAction(i)
   }
}

//noinspection GoUnusedFunction
func parseConfig(name string) {
   stream, err := ioutil.ReadFile(name)
   if err != nil {
	  log.Fatal(err)
   }
   input := string(stream)
   fmt.Println(input)
}

// start agent and open channels in and out stream
// input channel an listen to the structure value stream
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