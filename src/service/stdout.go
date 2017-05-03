package service

import (
   "fmt"
)

type ServiceStdout struct {}

func (ServiceStdout) SetAction(action string, data map[string]string) error {
   fmt.Println("AGENT STDOUT sortie .....")
   fmt.Println(action, data)
   return nil
}