package main

import "fmt"
import "strconv"

//Develope a program that reads a number from user as a string and print 10 times of that 
//number as a string only..suppose the number entered 12 output will be 120

func main(){
	var num string
	var zero int
	fmt.Println("Enter a number : ")
	fmt.Scanln(&num)
	zero = 0
	fmt.Println(num + strconv.Itoa(zero))

}
