package service

import (
   "fmt"
)

// Data strcuture for metrics service.
type ServiceMetrics struct {}

<<<<<<< HEAD
// Function to take action string
func (ServiceMetrics) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT METRICS sortie .....")
=======
// Get action from the rooter in order to send to services metrics.
func (ServiceMetrics) GetAction(action string, data map[string]string) error {
   println("AGENT METRICS sortie .....")
>>>>>>> dee82e212a07b920d050360d0c910f991e60807c
   fmt.Println(action, data)
   setConfigServices(data["image"], data["application_type"], data["id"])

   return nil
}
