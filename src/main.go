package main // Folder name

import ( // Libraries

	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gadielMa/golang-presentacion/src/examples"
	"github.com/gorilla/mux"
	"gopkg.in/matryer/respond.v1"
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

var GetHealth = func(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusOK, "Jose Maria Listooooortiii")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", GetHealth)

	fmt.Println("Escuchando conexiones en el puerto: 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

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
	fmt.Println("Ahora... Manos a la obra", p.Nombre)
	fmt.Println("")

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
