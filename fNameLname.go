package main

import "fmt"
//7
func main(){
	var fName, lName string
	fmt.Println(" Enter first name : ")
	fmt.Scanln(&fName)
	fmt.Println(" Enter Last name : ")
	fmt.Scanln(&lName)
	fmt.Println(fName,"",lName)
}