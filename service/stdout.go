package service

import (
   "fmt"
   "errors"
)

// Data structure for stdout service
type ServiceStdout struct {}


// Get action from the rooter in order to send to services stdout.
func (ServiceStdout) GetAction(action string, data map[string]string, client chan *ClientIN) error {
   if action == "" {
	  return errors.New("No action event")
   }
   fmt.Printf("action: %s \n data: %s \n", action, data)

   return nil
}
