package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	p := `password123`
	pa := `password1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)
	fmt.Println(bs)
	s := string(bs)
	fmt.Println(s)

	err = bcrypt.CompareHashAndPassword(bs, []byte(pa))
	if err != nil {
		fmt.Println("The password doesn't match")
	} else {
		fmt.Println("You're in!")
	}

}
