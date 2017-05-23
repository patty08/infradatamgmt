package service

import (
   "fmt"
   "os"
)

// Data strcuture for metrics service.
type ServiceMetrics struct {}


// Get action from the rooter in order to send to services metrics.
func (ServiceMetrics) GetAction(action string, data map[string]string) error {
   println("AGENT METRICS sortie .....")
   fmt.Println(action, data)
   switch action {
   case "delete":
      detachContainer(data)
   default:
      setConfigServices(data["image"], data["application_type"], data["id"], data["ip"])
   }
   return nil
}
// Todo: Delete all container's configurations (file, etc)
func detachContainer(data map[string]string){
   err:= os.Remove(dirCustom +data["application_type"]+"_"+data["id"]+".yml")
   if err != nil {
      fmt.Println(err)
      return
   }
   println("Detached..")
}