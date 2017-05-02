package rooter

import (
    "os"
	"time"
	"errors"
	"agent"
)

//interface for agent setting
type AgentIn interface {
	addEventListener(c chan *InfoIN, who string) error
}

type sAgentIn struct {
	AgentIn AgentIn
}

//interface for service setting
type ServiceOut interface {
	setAction(action string, data map[string]string) error
}

type sServiceOut struct {
	ServiceOut ServiceOut
}

// Exit listening and close stream
func closeListener()  {
	println("Exit")
	os.Exit(0)
}

func process(i *agent.InfoIN) {
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
func Start() {
	// open input channel and listening
	listener := make(chan *agent.InfoIN)
	// START ALL AGENTS

	a := sAgentIn{agent.AgentDocker{}}
	go a.AgentIn.addEventListener(listener, "unix:///var/run/docker.sock")

	for {
		response := <-listener
		go process(response)
		time.Sleep (time.Second * 1)
	}

	closeListener()
}