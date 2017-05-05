package client

// *** Strategy Service ***

//interface for service out settings
type ServiceOut interface {
   SetAction(action string, data map[string]string) error
}

// Structure for calling service
type sServiceOut struct {
   aServiceOut ServiceOut
}

// *** End Strategy Service ***
