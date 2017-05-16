package service

import (
   "fmt"
)

// Data structure for logging service.
type ServiceLogging struct {}

// Function to take action string
func (ServiceLogging) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT LOGGING sortie .....")
   fmt.Println(action, data)
   return nil
}