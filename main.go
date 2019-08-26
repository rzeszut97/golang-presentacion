package main // Folder name

import ( // Libraries
	"fmt"
	"time"
)

type Persona struct {
	Nombre string
}

func main() {
	var start interface{}
	fmt.Println("")
	fmt.Println(" ██████╗  ██████╗ ██╗      █████╗ ███╗   ██╗ ██████╗               ██████╗ ██████╗ ███████╗███████╗███████╗███╗   ██╗████████╗ █████╗  ██████╗██╗ ██████╗ ███╗   ██╗")
	fmt.Println("██╔════╝ ██╔═══██╗██║     ██╔══██╗████╗  ██║██╔════╝               ██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝████╗  ██║╚══██╔══╝██╔══██╗██╔════╝██║██╔═══██╗████╗  ██║")
	fmt.Println("██║  ███╗██║   ██║██║     ███████║██╔██╗ ██║██║  ███╗    █████╗    ██████╔╝██████╔╝█████╗  ███████╗█████╗  ██╔██╗ ██║   ██║   ███████║██║     ██║██║   ██║██╔██╗ ██║")
	fmt.Println("██║   ██║██║   ██║██║     ██╔══██║██║╚██╗██║██║   ██║    ╚════╝    ██╔═══╝ ██╔══██╗██╔══╝  ╚════██║██╔══╝  ██║╚██╗██║   ██║   ██╔══██║██║     ██║██║   ██║██║╚██╗██║")
	fmt.Println("╚██████╔╝╚██████╔╝███████╗██║  ██║██║ ╚████║╚██████╔╝              ██║     ██║  ██║███████╗███████║███████╗██║ ╚████║   ██║   ██║  ██║╚██████╗██║╚██████╔╝██║ ╚████║")
	fmt.Println(" ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝               ╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝")
	fmt.Println("")
	fmt.Println("Buen día " + getName() + "! Has ingresando al maravilloso universo de Golang...") // Stack
	fmt.Println("Presione una tecla para continuar...")
	fmt.Scanf("%s", &start)

	// Pointer - Example
	fmt.Println("Ejecutando Pointer Example")
	pointerExample()
	fmt.Println()

	// Go Routine - Example
	fmt.Println("Ejecutando Goroutines Example")
	goroutineExample()
	fmt.Println()

	// Channel - Example
	fmt.Println("Ejecutando Channel Example")
	channelExample()
	fmt.Println()

	persona := new(Persona)
	persona.despedirse()
}

/* func FuncName ( Parameters ) Returns */
func getName() string {
	name := "NN"
	fmt.Println("Ingrese su nombre: ")
	fmt.Scanf("%s", &name)
	fmt.Println()
	return name
}

func pointerExample() {
	a := 100
	b := &a            // a Pointer (memory location)
	fmt.Println(a, *b) // 100, Pointer data
	fmt.Println(&a, b)
	pointerModify(b)
	defer fmt.Println(a, *b)
	defer fmt.Println(&a, b)
}

func pointerModify(b *int) {
	*b = 10
}

func goroutineExample() {
	// Race
	go firstRunner()
	go secondRunner()
	go thirdRunner()
	time.Sleep(2 * time.Second)
}

func firstRunner() {
	fmt.Println("Llegó el corredor 1")
}

func secondRunner() {
	fmt.Println("Llegó el corredor 2")
}

func thirdRunner() {
	fmt.Println("Llegó el corredor 3")
}

func channelExample() {
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		ch1 <- 2000
	}()
	fmt.Println(<-ch1)

	ch2 := make(chan string)
	go func() {
		defer close(ch2)
		ch2 <- "A"
		ch2 <- "N"
		ch2 <- "D"
		ch2 <- "R"
		ch2 <- "E"
		ch2 <- "A"
		ch2 <- "N"
		ch2 <- "I"
	}()
	for n := range ch2 {
		fmt.Printf("%s", n)
		fmt.Println()
	}
}

func (p Persona) despedirse() {
	fmt.Println(p.Nombre)
}
