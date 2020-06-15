package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"../models"
)

var user *models.User

const (
	id           = 1
	username     = "alejandro_gpg"
	password     = "1234"
	passwordHash = "$2a$10$kdqO/H4SfP6GVaHyPhNoAexyRdMieOB9RRirc6v8VRNd/JvFlkUPC"
	email        = "alsjdask@gmail.com"
	createdDate  = "2020-06-14"
)

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(username, password, email)
	if err != nil {
		t.Error("No es posible crear el objeto", err)
	}

}

func TestSave(t *testing.T) {
	user, _ := models.NewUser(randomUsername(), password, email)
	if err := user.Save(); err != nil {
		t.Error("No es posible crear el usuario", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser(randomUsername(), password, email)
	if err != nil {
		t.Error("No es posible insertar el objeto", nil)
	}
}

func TestUniqueUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	if err == nil {
		t.Error("Es posible usernames duplicados")
	}
}

func TestDuplicateUsername(t *testing.T) {
	_, err := models.CreateUser(username, password, email)
	if err.Error() != "Error 1062: Duplicate entry 'alejandro_gpg' for key 'username'" {
		t.Error("Es posible tener un username duplicado en la base de datos!")
	}
}

func TestGetUser(t *testing.T) {
	user := models.GetUserById(id)
	if !equalsUser(user) || !equalsCreateDate(user.GetCreatedDate()) {
		t.Error("No es posible obtener el usuario")
	}
}

func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("No es posible obtener a los usuario")
	}
}

func TestPassword(t *testing.T) {
	user, _ := models.NewUser(username, password, email)
	if user.Password == password || len(user.Password) != 60 {
		t.Error("No es posible cifrar el password")
	}
}

func equalsUser(user *models.User) bool {
	return user.Username == username && user.Email == email
}

func equalsCreateDate(date time.Time) bool {
	t, _ := time.Parse("2020-01-02", createdDate)
	return t == date
}

/*
func TestDelete(t *testing.T) {
	if err := user.Delete(); err != nil {
		t.Error("No es posible borrar el usuario")
	}
}
*/
func randomUsername() string {
	return fmt.Sprintf("%s/%d", username, rand.Intn(1000))
}

func TestLogin(t *testing.T) {
	if valid := models.Login(username, password); !valid {
		t.Error("no es posible realizar el login")
	}
}

func TestNoLogin(t *testing.T) {
	if valid := models.Login(randomUsername(), password); valid {
		t.Error("Es posible realizar el login con parametros malos")
	}
}

func TestValidEmail(t *testing.T) {
	if err := models.ValidEmail(email); err != nil {
		t.Error("Validacion erronea en el email", err)
	}
}

func TestInValidEmail(t *testing.T) {
	if err := models.ValidEmail("dsajkajdkja.com"); err == nil {
		t.Error("Validacion erronea en el email")
	}
}

func TestUserNameLength(t *testing.T) {
	newUsername := username
	for i := 0; i < 10; i++ {
		newUsername += newUsername
	}

	_, err := models.NewUser(newUsername, password, email)
	if err == nil || err.Error() != "Username muy largo, maximo 30 caracteres" {
		t.Error("Es posible generar un usuario con username largo")
	}
}
