package models

import "errors"

type User struct {
	//En caso de usar json, xml o yaml solo cambiar en la notacion
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users []User

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{Id: 1, Username: "Alejandro", Password: "1234"}
	users[user.Id] = user
}

func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("El usuario no se encuentra en el map")
}

func SaveUser(user User) User {
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}

func UpdateUser(user User, username, password string) User {
	user.Username = username
	user.Password = password
	users[user.Id] = user
	return user
}

func DeleteUser(id int) {
	delete(users, id)
}
