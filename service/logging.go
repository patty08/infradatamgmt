package service

import (
   "fmt"
)

// Data structure for logging service.
type ServiceLogging struct {}

<<<<<<< HEAD
// Function to take action string
func (ServiceLogging) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT LOGGING sortie .....")
=======
// Get action from the rooter in order to send to services logging.
func (ServiceLogging) GetAction(action string, data map[string]string) error {
   println("AGENT LOGGING sortie .....")
>>>>>>> dee82e212a07b920d050360d0c910f991e60807c
   fmt.Println(action, data)
   // todo: check how to proceed
   setConfigServices(data["image"], data["application_type"], data["id"])
   return nil
}