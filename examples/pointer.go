package examples

import "fmt"

func PointerExample() {
	a := "SoyUnaVariable"
	p := &a // Pointer to Memory Location

	fmt.Println("a: ", a)
	fmt.Println("&a: ", &a)
	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)
	fmt.Println("")

	pointerModify(p)

	// Stack
	defer fmt.Println("a: ", a)
	defer fmt.Println("&a: ", &a)
	defer fmt.Println("p: ", p)
	defer fmt.Println("*p: ", *p)
}

func pointerModify(p *string) {
	*p = "SoyUnaNuevaVariable"
}
