package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//read file
	//methode 1:-open the file and read the file if needed one by one
	f, err := os.Open("example.txt")
	if err != nil {
		//log the error
		panic(err)
	}
	fillInfo, err := f.Stat()
	if err != nil {
		//log the error
		panic(err)
	}
	defer f.Close() //best to use as soon as possible after opening the file
	fmt.Println("file information :", fillInfo.Name())
	fmt.Println("file information :", fillInfo.Size())
	fmt.Println("file information :", fillInfo.IsDir())
	fmt.Println("file information :", fillInfo.ModTime())
	//methode 2:- retreive data in one reading
	buf := make([]byte, fillInfo.Size())
	d, err := f.Read(buf)
	if err != nil {
		//log the error
		panic(err)
	}
	fmt.Println("data", d, " ", buf)
	for i := 0; i < int(fillInfo.Size()); i++ {
		fmt.Println("data", d, string(buf[i]))
	}
	//methode 3:- store file in one go not best for large files
	data, err := os.ReadFile("example")
	if err != nil {
		//log the error
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Printf("\n ----FOLDERS---- \n")
	//read folder

	folder, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	defer folder.Close()
	//dir, err := folder.ReadDir(-1)
	dir, err := folder.ReadDir(2) //positive value indicate how much the out filenno.will be if we put value<0 then it will show all the filesand folder present in that directory
	for _, fi := range dir {
		fmt.Println(fi.Name(), fi.IsDir())
	}
	//create file
	f1, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	//write in file
	//methode 1:- directly inserting the data
	f1.WriteString("hi go")
	f1.WriteString(" nice language")

	//methode 2:-inserting the data from byte
	byte := []byte(" hello go glang")
	f1.Write(byte)

	//print the content of exaple2.txt file
	da, err := os.ReadFile("example2.txt")
	if err != nil {
		//log the error
		panic(err)
	}
	fmt.Println(string(da))

	//read and write to another file(streaming fashion)
	sourceFile, err := os.Open("example.txt") //source file
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	destFile, err := os.Create("exapmle3.txt") //destination file
	if err != nil {
		panic(err)
	}
	defer destFile.Close()
	reader := bufio.NewReader(sourceFile) //reader
	writer := bufio.NewWriter(destFile)   //writer
	for {
		b, err := reader.ReadByte() //read byte by byte
		if err != nil {
			if err.Error() != "EOF" { //check that we have not reached end of  file
				panic(err)
			}
			break //if reached end of file the break the infinite loop
		}
		werr := writer.WriteByte(b) //write byte by byte
		if werr != nil {
			panic(werr)
		}
	}
	writer.Flush() //flush if anything is left
	fmt.Println("data transferred...")

	//delete file

	//sample file creation for deleating it and then commenting it to rundelete commands

	// 	f2,r:=os.Create("example4.txt")
	// if r!=nil{
	// 	panic(r)
	// }
	// defer f2.Close()
	// fmt.Println("file created")

	//delete file command
	er := os.Remove("example4.txt")
	if er != nil {
		panic(er)
	}
	fmt.Println("file deleted")
}
