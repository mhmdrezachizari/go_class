package main
import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		fmt.Println("chai dar hal amade shodan")
		time.Sleep(2*time.Second)
		fmt.Println("chai  amade shod")
	}()
	fmt.Println("man miz ra michinam")
	wg.Wait()
	fmt.Println("mitavanid chai benoshid")
}