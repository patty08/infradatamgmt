package service

import (
   "fmt"
)

type ServiceLogging struct {}

func (ServiceLogging) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT LOGGING sortie .....")
   fmt.Println(action, data)
   return nil
}