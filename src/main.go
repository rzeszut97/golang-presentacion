package main // Folder name

import ( // Libraries

	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gadielMa/golang-presentacion/src/examples"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", GetHealth).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Escuchando conexiones en el puerto :", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
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

type M map[string]interface{}

var GetHealth = func(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusOK, generateJson("UP", ""))
}

func generateJson(status string, msjError string) M {
	return M{
		"status": M{
			"code":        status,
			"description": "",
		},
		"details": M{
			"sqlServerHealthIndicator": M{
				"status": M{
					"code":        status,
					"description": msjError,
				},
				"details": M{
					"host":    "",
					"version": "",
				},
			},
		},
	}
}
