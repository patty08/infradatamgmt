package service

import (
   "os"
   "fmt"
   "io"
   "bytes"
	"github.com/spf13/viper"
)

// Change agent's name in parameters to file name (eg: surikator -> surikator.yml).
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

// Search and load yalm config file
func loadConfigFile(file string) *viper.Viper{
	cfg := viper.New()
	cfg.SetConfigFile(file)
	err := cfg.ReadInConfig()
	if err != nil {
		_, e := os.Create(file)
		if e != nil {
			panic(e)
		}
	}

	return cfg
}