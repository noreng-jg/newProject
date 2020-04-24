package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	
	b,err:=ioutil.ReadFile("gbg")
	if err != nil{
		fmt.Print(err)
	}

	str:=string(b)
	fmt.Println(str)
	
}