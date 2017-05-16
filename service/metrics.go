package service

import (
   "fmt"
)

// Data strcuture for metrics service.
type ServiceMetrics struct {}

// Function to take action string
func (ServiceMetrics) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT METRICS sortie .....")
   fmt.Println(action, data)
   return nil
}