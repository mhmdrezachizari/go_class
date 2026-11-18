package main
import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(3)
	for i:=0 ; i<3 ;i++{
		go func(){
		  defer wg.Done()
		  fmt.Println("yek kar anjam shod")
		}()
	}
	wg.Wait()
	fmt.Println("har se kar tamom shodand")
}