package test

import (
	"testing"
)

func TestHolaMundo(t *testing.T) {
	str := "Hola mundo"
	if str != "Hola mundo" {
		t.Error("No es posible saludar a los usuarios", nil)
	}
}
