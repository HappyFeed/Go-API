package models

import "errors"

type ValidationError error

var (
	errorUsername      = ValidationError(errors.New("El username no debe estar vacio"))
	errorShortUsername = ValidationError(errors.New("El username es muy corto"))
	errorLongUsername  = ValidationError(errors.New("El username es muy extenso"))

	errorEmail = ValidationError(errors.New("Formato invalido de email"))

	errorPasswordEncryption = ValidationError(errors.New("No es posible cifrar el texto"))

	errorLogin = ValidationError(errors.New("Usuario o password incorrecto"))
)
