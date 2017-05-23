package service

import (
   "fmt"

)

// Data strcuture for metrics service.
type ServiceMetrics struct {}


// Get action from the rooter in order to send to services metrics.
func (ServiceMetrics) GetAction(action string, data map[string]string) error {
   switch action {
   case "stop":
	   detachContainer(data)
   case "delete":
	   detachContainer(data)
   case "start":
      setConfigServices(data["image"], data["application_type"], data["id"], data["ip"])
   }
   return nil
}
