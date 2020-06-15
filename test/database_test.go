package test

import (
	"testing"

	"../models"
)

func TestConnection(t *testing.T) {
	connection := models.GetConnection()
	if connection == nil {
		t.Error("No es posile realizar la conexion", nil)
	}
}
