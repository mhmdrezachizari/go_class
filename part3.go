package main
import "fmt"

func main()  {
	var number int = 1
	for number != 0{
		fmt.Print("enter number")
		fmt.Scan(&number)
		fmt.Println("number:" , number)
	}
	fmt.Println("finished")
}