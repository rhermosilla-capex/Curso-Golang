package waitgroups

import (
	"fmt"
	"sync"
)

func WGRoutine(a, b chan func() string, wg *sync.WaitGroup) {
	for {
		select {
		case x := <-a:
			value := x()
			fmt.Println(value)
			wg.Done()
		case y := <-b:
			value := y()
			fmt.Println(value)
			wg.Done()
		}
	}
}
