package main

import (
   "os"
   "time"
   "errors"
)

// Data structure of informations channels
// action stand for action event, service stand for services whish in output, info all informations
type infoIN struct{
   Action   string
   Services []string
   Data     map[string] string
}

//interface for agent setting
type AgentIn interface {
   addEventListener(c chan *infoIN, who string)
}

type sAgentIn struct {
   AgentIn AgentIn
}

//interface for service setting
type ServiceOut interface {
   setAction(action string, data map[string]string)
}

type sServiceOut struct {
   ServiceOut ServiceOut
}

// Exit listening and close stream
func closeListener()  {
   println("Exit")
   os.Exit(0)
}

func process(i *infoIN) {
   if i != nil {
	  errors.New("errors....")
   }
   // var agent ServiceOut

   // TODO !!!

   // service := sServiceOut{ServiceSTDOUT{}}
   // go service.setAction(i.Action, i.Data)

}

// start agent and open channels in and out stream
// input channel an listen to the structure value stream
// initialise des channels d'ecoute puis recupère les données du channel, traite les données inspecter puis envoi la configuration en sortie
func main() {
   // open input channel and listening
   listener := make(chan *infoIN)
   // START ALL AGENTS

   // agent := sAgentIn{AgendDocker{}}
   // go agent.startEventListener(listener)

   for {
	  response := <-listener
	  go process(response)
	  time.Sleep (time.Second * 1)
   }

   closeListener()
}