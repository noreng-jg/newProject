//exemplo de map com post
package main

import "fmt"


func main(){
post:=make(map[int]string)//id e titulo
post[1]="Primeiro post"
fmt.Println(post[1])
}



