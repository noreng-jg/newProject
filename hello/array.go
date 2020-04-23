package main

import "fmt"
//soma de coordenadas cartesianas
func main(){
	var arr1= [3]int{2,1,1}
	var arr2= [3]int{4,2,1}
	var arr3= new([3]int)//novo array sem atribuicao
	for i:=0; i<len(arr1); i++{
		arr3[i]=arr1[i]+arr2[i]
	}

	fmt.Println(arr3)
}
