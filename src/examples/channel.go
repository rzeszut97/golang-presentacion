package examples

import "fmt"

func ChannelExample() {
	ch1 := make(chan bool)

	go func() {
		defer close(ch1)
		ch1 <- true
	}()

	fmt.Println("Creamos un channel:", <-ch1)
	fmt.Println("")

	ch2 := make(chan string)

	go func(ch chan string) {
		ch <- "	A"
		ch <- "	N"
		ch <- "	D"
		ch <- "	R"
		ch <- "	E"
		ch <- "	A"
		ch <- "	N"
		ch <- "	I"
		defer close(ch)
	}(ch2)

	for n := range ch2 {
		fmt.Printf("%s\n", n)
	}
}
