// Generator is a function that returns a channel. We launch the goroutine from
// inside the generator.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("blah!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You said: %q\n", <-c)
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)

			n := time.Duration(rand.Intn(1e3))
			time.Sleep(n * time.Millisecond)
		}
	}()
	return c
}
