package main

import (
	"fmt"

	"github.com/codebyankita/podcast/auth"
	"github.com/codebyankita/podcast/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("codebyankita", "secret")
	session := auth.GetSession()

	fmt.Println("session", session)

	user := user.User{
		Email: "user@email.com",
		// Name:  "John Doe",
	}

	// fmt.Println(user.Email, user.Name)
	color.Green(user.Email)

}
