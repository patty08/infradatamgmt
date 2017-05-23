package service

import (
   "os"
   "fmt"
   "io"
   "bytes"
   "io/ioutil"
   "log"
   "strings"
)

// configuration directory
const dirOriginal string = "./rooter/configuration/metricbeat/conf/original/"
const dirCustom string = "./rooter/configuration/metricbeat/conf/custom/"

// Formatted the host id to string value:
//	hosts: ["ID"].
func formatHostName(hostIP string) string{
   if hostIP == "" {
	  fmt.Println("host name is empty")
   }
   name := bytes.NewBufferString(  "hosts: [\""+hostIP)
   name.WriteString("\"]")
   return name.String()
}

// Function to replace the line hosts in the configuration file to the id of the container in /custom/file.yml.
// ipContainer is the ip adress of the running container
// image is the base image of the container (application_type)
// ipContainer is the ip adress of the running container
func setidConfiguration(idContainer string, image string, ipContainer string) {
   // format host
   frmHost := formatHostName(ipContainer)
   println(frmHost)

   // get file name from the name agent
   // check if the configuration is available in host
   fd, err := ioutil.ReadFile(dirCustom +image+"_"+idContainer+".yml")
   if err != nil {
	  println("dest file not found:"+ dirCustom + image + "_" + idContainer+ ".yml")
	  log.Fatalln(err)
   }

   lines := strings.Split(string(fd), "\n")
   for i, line := range lines {
	  if strings.Contains(line, " hosts: [") {
		 lines[i] = frmHost
	  }
   }
   output := strings.Join(lines, "\n")
   err = ioutil.WriteFile(dirCustom +string(image+"_"+idContainer+".yml"), []byte(output), 0644)
   if err != nil {
	  log.Fatalln(err)
   }
}

// Copy configuration file to /usr/share/metricbeat/custom/, open the file and setidConfiguration.
// 	image is the name of the image docker hub.
// 	Application_type must be set by user for setting configuration processors,
// 	if not nomconf is checked for matching similar file configuration.
// 	id is the id of the container.
func setConfigServices(image string, application_type string, id string, ip string)  {
   if application_type == "" {
	  application_type = image
   }
   application_type = strings.ToLower(application_type)
   // check if configuration exist
   println(application_type)

   // get file name from name agent: name.yml
   agentName := formatNameConfig(application_type)

   // check if the configuration is available in host
   file, err := os.OpenFile(dirOriginal +string(agentName),0,777)
   if err != nil {
	  fmt.Println("src file not found:" +dirOriginal +agentName)
	  println(err)
   }
   // close at last
   defer file.Close()

   CopyFile(file, dirCustom +application_type+"_"+id+".yml")

   // replace the id in the file configuration with the ids in data
   // find host: [" and replace to host : ["id"] in the custom configuration
   setidConfiguration(id, application_type, ip)
   println(">> configuration is set: "+string(agentName))
}

// Change agent's name in parameters to file name (eg: docker -> docker.yml).
func formatNameConfig(a string) string {
   if a == "" {
	  fmt.Println("configuration name corrupted ")
   }
   name := bytes.NewBufferString(a)
   name.WriteString(".yml")
   return name.String()
}

// Copy source file in argument to destination source. The path where the file must be paste (see: dirOriginal, dirCustom)
func CopyFile(in io.Reader, dst string) (err error) {

   // Does file already exist? Skip
   if _, err := os.Stat(dst); err == nil {
	  return nil
   }

   out, err := os.Create(dst)
   if err != nil {
	  fmt.Println("Error creating file", err)
	  return
   }

   defer func() {
	  cerr := out.Close()
	  if err == nil {
		 err = cerr
	  }
   }()
   var bytes int64
   if bytes, err = io.Copy(out, in); err != nil {
	  fmt.Println("io.Copy error")
	  return
   }
   fmt.Println(bytes)

   err = out.Sync()
   return
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

