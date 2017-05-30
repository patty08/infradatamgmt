package service

import (
   "os"
   "testing"
   "github.com/stretchr/testify/assert"
   "io/ioutil"
)

// Copy source file in argument to destination source. The path where the file must be paste (see: dirOriginal, dirCustom)
func TestCopyFile(t *testing.T) {
   file, err := os.Create("./testfile.yml")
   assert.NotNil(t, err)
   s := "Test confifuration.go"
   file.WriteString(s)

   ofile, err := os.OpenFile( "testfile.yml",0,777)
   assert.Error(t, err)
   defer ofile.Close()


   err = CopyFile(file,"copyfile.yml")
   assert.Nil(t,err)
   _, err = os.Open("copyfile.yml")
   assert.NotNil(t, err)

   testf, err := os.OpenFile( "testfile.yml",0,777)
   assert.Error(t, err)

   cpf, err := os.OpenFile( "copyfile.yml",0,777)
   assert.Error(t, err)

   src, err := ioutil.ReadAll(testf)
   dst, err := ioutil.ReadAll(cpf)
   t.Log(src)
   //assert.Equal(t, s, src, "wrong string")
   assert.Equal(t, src, dst, "bad sector file")

}

func TestServiceMetrics_GetAction(t *testing.T) {


}

// Change agent's name in parameters to file name (eg: surikator -> surikator.yml).
func TestformatNameConfig(t *testing.T){

}