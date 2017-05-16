package service

import (
   "fmt"
)

// Data strcuture for metrics service.
type ServiceMetrics struct {}


// Get action from the rooter in order to send to services metrics.
func (ServiceMetrics) GetAction(action string, data map[string]string) error {
   println("AGENT METRICS sortie .....")
   fmt.Println(action, data)
   setConfigServices(data["image"], data["application_type"], data["id"])

   return nil
}
