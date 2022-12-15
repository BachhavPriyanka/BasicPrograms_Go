package main
import "fmt"
import "strings"

//4

func main(){
	var place string
	fmt.Println("Enter place name where you would like to visit : ")
	fmt.Scanln(&place)
	fmt.Println("Upper Case : ", strings.ToUpper(place))
	fmt.Println("Lower Case : ", strings.ToLower(place))
}