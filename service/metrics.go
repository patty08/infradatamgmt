package service

import (
   "fmt"
)

type ServiceMetrics struct {}

func (ServiceMetrics) GetAction(action string, data map[string]string) error {
   fmt.Println("AGENT METRICS sortie .....")
   fmt.Println(action, data)

   //setConfigServices(data["image"], data["application_type"], data["id"])

   return nil
}