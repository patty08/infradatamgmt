package service

import (
   "fmt"
   "errors"
)

// Data structure for stdout service
type ServiceStdout struct {}

// Function to take action string
func (ServiceStdout) SetAction(action string, data map[string]string) error {
   if action == "" {
	  return errors.New("No action event")
   }
   fmt.Printf("action: %s \n data: %s \n", action, data)

   return nil
}
