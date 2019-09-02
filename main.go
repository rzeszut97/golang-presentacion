package main // Folder name

import ( // Libraries

	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gadielMa/golang-presentacion/examples"
	"github.com/gorilla/mux"
)

type Persona struct {
	Nombre     string
	Apellido   string
	DNI        int
	Telefono   string
	Ejercicios bool
}

var subdito = new(Persona)

// Colors
var (
	greenFmt    = color.New(color.FgGreen)
	cyanFmt     = color.New(color.FgCyan)
	boldFmt     = color.New(color.Bold)
	cyanBoldFmt = color.New(color.FgCyan, color.Bold)
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)
	log.Fatal(http.ListenAndServe(":8000", r))

	cyanFmt.Println("\n ██████╗  ██████╗ ██╗      █████╗ ███╗   ██╗ ██████╗               ██████╗ ██████╗ ███████╗███████╗███████╗███╗   ██╗████████╗ █████╗  ██████╗██╗ ██████╗ ███╗   ██╗")
	cyanFmt.Println("██╔════╝ ██╔═══██╗██║     ██╔══██╗████╗  ██║██╔════╝               ██╔══██╗██╔══██╗██╔════╝██╔════╝██╔════╝████╗  ██║╚══██╔══╝██╔══██╗██╔════╝██║██╔═══██╗████╗  ██║")
	cyanFmt.Println("██║  ███╗██║   ██║██║     ███████║██╔██╗ ██║██║  ███╗    █████╗    ██████╔╝██████╔╝█████╗  ███████╗█████╗  ██╔██╗ ██║   ██║   ███████║██║     ██║██║   ██║██╔██╗ ██║")
	cyanFmt.Println("██║   ██║██║   ██║██║     ██╔══██║██║╚██╗██║██║   ██║    ╚════╝    ██╔═══╝ ██╔══██╗██╔══╝  ╚════██║██╔══╝  ██║╚██╗██║   ██║   ██╔══██║██║     ██║██║   ██║██║╚██╗██║")
	cyanFmt.Println("╚██████╔╝╚██████╔╝███████╗██║  ██║██║ ╚████║╚██████╔╝              ██║     ██║  ██║███████╗███████║███████╗██║ ╚████║   ██║   ██║  ██║╚██████╗██║╚██████╔╝██║ ╚████║")
	cyanFmt.Println(" ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝               ╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═══╝   ╚═╝   ╚═╝  ╚═╝ ╚═════╝╚═╝ ╚═════╝ ╚═╝  ╚═══╝")

	cyanBoldFmt.Print("\nIngrese su nombre: ")
	fmt.Scanf("%s", &subdito.Nombre)

	cyanBoldFmt.Print("\nBuen día ")
	boldFmt.Print(subdito.Nombre, "!")
	cyanBoldFmt.Print(" Has ingresando al maravilloso universo de Golang.\n")
	continuar()

	// Pointer - Example
	genericExample("Pointer", examples.PointerExample)

	// Go Routine - Example
	genericExample("Go Routine", examples.GoroutineExample)

	// Channel - Example
	genericExample("Channel", examples.ChannelExample)

	subdito.despedirse()
}

/* func FuncName ( Parameters ) Returns */
func (p Persona) despedirse() {
	fmt.Println("Ahora... Manos a la obra", p.Nombre, "!\n")
}

func continuar() {
	cyanFmt.Println("\nPresione una tecla para continuar...")
	fmt.Scanf("%s", &subdito.Ejercicios)
}

func genericExample(example string, function func()) {
	greenFmt.Println("#################################################")
	greenFmt.Println("######### Ejecutando " + example + " Example #########")
	greenFmt.Println("#################################################")
	function()
	continuar()
}
