package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func bcryptSome() {
	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(bs)

	check := "password123"

	err = bcrypt.CompareHashAndPassword(bs, []byte(check))
	if err != nil {
		fmt.Println("You can't login")
		return
	}
	fmt.Println("You're logged in")

}
