package client

// *** Strategy Service ***

//Interface for service out settings
type ServiceOut interface {
   GetAction(action string, data map[string]string) error
}

// Structure for calling service
type sServiceOut struct {
   aServiceOut ServiceOut
}