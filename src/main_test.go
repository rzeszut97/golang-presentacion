package main

import "testing"

// Sintaxis ´*_test.go´ ´Test*´

func TestDespedida(t *testing.T) {
	gadiel := &Persona{Nombre: "Gadiel"}
	gadiel.despedirse()
}
