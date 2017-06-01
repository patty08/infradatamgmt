package client

import "github.com/sebastienmusso/infradatamgmt/service"

// *** Strategy Service ***

//Interface for service out settings
type ServiceOut interface {
   GetAction(action string, data map[string]string, client chan *service.ClientIN) error
}

// Structure for calling service
type sServiceOut struct {
   aServiceOut ServiceOut
}