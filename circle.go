package main
import "fmt"

//2. Write a program which takes diameter (float32) of a circle and computes perimeter and 
//area of the circle

func main(){
	var area, radius, perimeter float32
	fmt.Println("Enter the radius : ")
	fmt.Scanln(&radius)
	area = 3.14 * radius * radius
	perimeter = 2 * (22 / 7.0) * radius
	fmt.Println("Area of circle : ", area)
	fmt.Println("Perimeter of circle : ", perimeter)
}