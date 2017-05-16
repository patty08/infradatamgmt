package service

import (
   "fmt"
   "errors"
)

// Data structure for stdout service
type ServiceStdout struct {}

<<<<<<< HEAD
// Function to take action string
func (ServiceStdout) SetAction(action string, data map[string]string) error {
=======
// Get action from the rooter in order to send to services stdout.
func (ServiceStdout) GetAction(action string, data map[string]string) error {
>>>>>>> dee82e212a07b920d050360d0c910f991e60807c
   if action == "" {
	  return errors.New("No action event")
   }
   fmt.Printf("action: %s \n data: %s \n", action, data)

   return nil
}
