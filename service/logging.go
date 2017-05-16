package service

import (
   "fmt"
)

type ServiceLogging struct {}

func (ServiceLogging) GetAction(action string, data map[string]string) error {
   fmt.Println("AGENT LOGGING sortie .....")
   fmt.Println(action, data)
   return nil
}