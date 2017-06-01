package service

// Data structure of informations channels.
// Message for orchestrator client
type ClientIN struct{
	Action   string
	Data     map[string] string
}