package service

import (
   "fmt"
   "errors"
)

// data structure for stdout service
type ServiceStdout struct {}

// function to take action sting
func (ServiceStdout) SetAction(action string, data map[string]string) error {
   if action != "" {
	  errors.New("errors stodout....")
   }
   fmt.Printf("action: %s \n data: %s \n", action, data)
   fmt.Println("->")

   return nil
}
