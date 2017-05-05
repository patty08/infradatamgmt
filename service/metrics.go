package service

import (
   "fmt"
)

type ServiceMetrics struct {}

func (ServiceMetrics) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT METRICS sortie .....")
   fmt.Println(action, data)
   return nil
}