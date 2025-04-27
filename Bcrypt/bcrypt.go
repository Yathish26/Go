package main

import(
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	var password string
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	passwordBytes := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	fmt.Println("Your bcrypt hashed password is:")
	fmt.Println(string(hashedPassword))
}

