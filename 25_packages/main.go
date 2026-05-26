package main //<-- enerty part

// this packages are custom made by us
import (
	"fmt"

	"github.com/Mritunjay2005/GOlang/auth"
	"github.com/Mritunjay2005/GOlang/users"
	"github.com/fatih/color"
)

func main() { //<-- entry point of the go progrma
	// we just need to ceate a mod
	// the best practice is to write the name of the git repo name
	//command is --> go mod init github.com/Mritunjay2005/GOlang

	//we are able to export this package because this have starting letter as capital
	auth.LoginWithCrediential("piyush", "2007")

	session := auth.GetSession()
	fmt.Println("session is ", session)

	user := users.User{
		Email: "piyush@gmail.com",
		Name:  "Piyush",
	}
	fmt.Println(user.Email)

	// but if we need to import packages from outside of then -->3rd party packages

	//to import the package the command is --> go get "link of the package"
	//       or
	// we can past the link of the package in the import section and then run go mod tidy --> it will automatically install the package
	color.Green(user.Email)
}
