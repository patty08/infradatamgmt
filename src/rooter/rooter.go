package rooter

import (
	"agent"
	"client"

	"errors"
	"time"
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

func process(i *agent.InfoIN) error {
	if i == nil {
		return errors.New("InfoIn Structure Error on process")
	}
	//namefile := "service/config.txt"
	//parseConfig(namefile)
	var err error
	switch i.Data["client"] {
	case "docker":
		agent := sClientOut{client.ClientDocker{}}
		err = agent.aClientOut.SetAction(i)
	default:
		agent := sClientOut{client.ClientDocker{}}
		err = agent.aClientOut.SetAction(i)
	}

	return err
}

// TODO implementation config file for rooter
/*func parseConfig(namefile string) {
	stream, err := ioutil.ReadFile(namefile)
	if err != nil {
		log.Fatal(err)
	}
	lireFichier := string(stream)
	fmt.Println(lireFichier)
}*/

// start agent and open channels in and out stream
// input channel an listen to the structure value stream
// initialise des channels d'ecoute puis recupère les données du channel, traite les données inspecter puis envoi la configuration en sortie
func Start() {
	// open input channel and listening
	listener := make(chan *agent.InfoIN)
	// START ALL AGENTS
	a := sAgentIn{agent.AgentDocker{}}
	go a.AgentIn.AddEventListener(listener, "unix:///var/run/docker.sock")
	for {
		go process(<-listener)
		time.Sleep(time.Second * 1)
	}
}
