package main

import "fmt"

func main()  {
	scores := []int{12,16,19}
	fmt.Println(scores[0])
	fmt.Println(len(scores))
	scores = append(scores, 20)
		fmt.Println(scores)

}