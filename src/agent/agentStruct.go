package agent

// Data structure of informations channels
// action stand for action event, service stand for services whish in output, info is all data information
type InfoIN struct{
	Action   string
	Services []string
	Data     map[string] string
}