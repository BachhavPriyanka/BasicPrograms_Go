package main
import "math/rand"
import  "fmt"
import "time"

//6

func main(){
	
	var arr [5]int
	var avg int
	sum := 0
	min := 10
	max := 50
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++{
		arr[i] = rand.Intn((max - min) + max)
		fmt.Println(arr[i])
		sum += arr[i]
	}
	avg = sum/5
	fmt.Println("Average : ",avg,"%") 
}