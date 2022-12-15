package main
import "fmt"

//3

func main(){
	var celcius, fahrenheit float32
	fmt.Println("Enter Temperature in celcius : ")
	fmt.Scanln(&celcius)
	fahrenheit = ((celcius*9)/5) + 32
	fmt.Println("Fahrenheit :", fahrenheit)
}