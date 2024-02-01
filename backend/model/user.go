package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
}

func (user *User) SetPassword(password string) error {

	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4); err != nil {
		log.Println("Failed to set password for user " + user.Email)
		return err
	} else {
		user.Password = hashedPassword
	}
	log.Println("Password set" + string(user.Password))
	return nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
