package main

import (
	"testing"
)

func TestDespedida(t *testing.T) {
	gadiel := &Persona{Nombre: "Gadiel"}
	gadiel.despedirse()
}
