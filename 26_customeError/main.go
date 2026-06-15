package main

import (
	"fmt"
	"os"
	"strings"
)

type NoTextError struct {
	Message string
}

func (n *NoTextError) Error() string {
	return n.Message
}
func loadFile(fileName string) (string, error) {
	if !strings.HasSuffix(fileName, ".txt") {
		return "", &NoTextError{
			"opening non txt file",
		}
	}
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return fileName, nil
}
func main() {
	//s,err:=loadFile("go.mod")
	//  will return
	//  syntax:
	// error:  opening non txt file

	//  s,err:=loadFile("go.txt")
	//  will return
	// syntax:
	// error:  open go.txt: The system cannot find the file specified.

	s, err := loadFile("go.txt")
	// error,ls is used to check if the error is of a specific type, in this case, os.ErrNotExist.
	//  If the error is of that type, it prints "File does not exist". Otherwise, it prints "Unexpected error".
	//  Finally, it prints the syntax and the error message.
	if error.ls(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("Unexpected error")
	}

    // var noTextError *NoTextError
	// if error.ls(err, os.ErrNotExist) {
	// 	fmt.Println("File is not a text file")
	// }else if error.As(err, &noTextError){
	// 	fmt.Println("Text Error " + noTextError.Message)
	// }
	fmt.Println("syntax: ", s)
	fmt.Println("error: ", err)
}
