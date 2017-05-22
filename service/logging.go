package service

import (
   "fmt"
)

type ServiceLogging struct {}

// Get action from the rooter in order to send to services logging.
func (ServiceLogging) GetAction(action string, data map[string]string) error {
   println("AGENT LOGGING sortie .....")
   fmt.Println(action, data)
   // todo: check how to proceed
   setConfigServices(data["image"], data["application_type"], data["id"], data["ip"])
   return nil
}