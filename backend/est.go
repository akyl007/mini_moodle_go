package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash := "$2a$10$glKtqXH5KgkAuTkVxhP30elh5.oKD9a5q3Q0OtG5Ux2NlvnW8Mxwq"
	password := "testpassword"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("Пароли не совпадают:", err)
	} else {
		fmt.Println("Пароли совпадают")
	}
}
