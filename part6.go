package main
import "fmt"

func main()  {
	grades:=[3]int{12,16,19}
	fmt.Println(grades[0])
	grades[1] = 18
	fmt.Println(grades)

}