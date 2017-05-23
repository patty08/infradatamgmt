// Developed by Treeptik team @2017

/*
	Package agent contains all the agent use to monitor.

	Docker Agent

	Docker agent is used to listen docker events and parse the information into a standard format.
	The agent listen to the docker socket.The agent use label to select events.
	Only events with the following label are parsed by the agent:
		"monitor=enabled"
	Information are parsed into the following format:
		  Action: string
		  Service: []string
		  Data: map[string]string{}


*/
package agent
