package main

import (
	"github.com/spf13/afero"
	"fmt"
)


func main(){
	fmt.Println("nibba")
	aferoTest()
}

func aferoTest(){
	var AppFs = afero.NewMemMapFs()
	AppFs.Open("/Users/misael/Desktop/Servidor/tests")
	fs := new(afero.NewMemMapFs)
	f, err := afero.TempFile(fs,"","ioutil-test")
}