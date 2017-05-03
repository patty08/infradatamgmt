package service

import (
   "fmt"
)

type ServiceDocker struct {}

func (ServiceDocker) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT DOCKER sortie .....")
   fmt.Println(action, data)
   return nil
}