package main
import "fmt"
import "time"
//8

func main(){
	time := time.Now()
	fmt.Println("Date :", time.Format("2006-01-02"))	
	fmt.Println("Time :", time.Format("15:04:05"))
}