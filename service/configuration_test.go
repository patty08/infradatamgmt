package service

import (
   "os"
   "testing"
   "github.com/stretchr/testify/assert"
)

// Copy source file in argument to destination source. The path where the file must be paste (see: dirOriginal, dirCustom)
func TestCopyFile(t *testing.T) {
   file, err := os.Create("./testfile.yml")
   assert.New(t)
   s := "fichier test"

   file.WriteString(s)
   println(file)

   assert.NotNil(t,file, "empty file")

   ofile, err := os.OpenFile( "./testfile.yml",0,777)
   assert.NotNil(t, ofile)

   err = CopyFile(ofile,"./copyfile.yml")
   assert.Nil(t,err)

   _, err = os.Open("./copyfile.yml")
   assert.NotNil(t, err)

   defer ofile.Close()
}

func TestServiceMetrics_GetAction(t *testing.T) {


}
