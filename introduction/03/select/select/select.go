package selectexample

import "fmt"

func Routine(a, b *chan string) {
	for {
		select {
		case x := <-*a:
			fmt.Printf("a: %s\n", x)
		case y := <-*b:
			fmt.Printf("b: %s\n", y)
		}
	}
}
