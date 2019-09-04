package examples

import (
	"fmt"
	"time"
)

func GoroutineExample() {
	// Race
	go firstRunner()
	go secondRunner()
	go thirdRunner()
	go fourthRunner()
	go fifthRunner()
	time.Sleep(1 * time.Second)
}

func firstRunner() {
	fmt.Println("Race -> Llegó el corredor 1")
}

func secondRunner() {
	fmt.Println("Race -> Llegó el corredor 2")
}

func thirdRunner() {
	fmt.Println("Race -> Llegó el corredor 3")
}

func fourthRunner() {
	fmt.Println("Race -> Llegó el corredor 4")
}

func fifthRunner() {
	fmt.Println("Race -> Llegó el corredor 5")
}
