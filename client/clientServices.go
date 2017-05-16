package client

// *** Strategy Service ***

//Interface for service out settings
type ServiceOut interface {
   SetAction(action string, data map[string]string) error
}

// Structure for calling service
type sServiceOut struct {
   aServiceOut ServiceOut
}

// *** End Strategy Service ***
