package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(a , b int){
	  defer wg.Done()
	  fmt.Println("jam an barabar ast ba" , a+b)
	}(5,3)

	wg.Wait()
}