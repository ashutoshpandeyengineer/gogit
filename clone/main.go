package main

import (
	"fmt"
	"os"
    "log"
	"io"
	"io/ioutil"
	"github.com/go-git/go-git/v5"
	
)
func main() {
	
	r,err:=git.PlainClone("./ashu", false ,&git.CloneOptions{ URL: "https://github.com/ashutoshpandeyengineer/go-git.git" ,Progress : os.Stdout,})
	//********************path of the localstorage*****************Url of the directory from where you want to clone*********************
	if err != nil {log.Fatal(err)}
		
	file , err:=os.Create("./ashu.txt")
	if err != nil { log.Fatal(err) }

	length, err := io.WriteString(file , r.Storer.NewEncodedObject().Hash().String())
	if err != nil { log.Fatal(err) }
	
	fmt.Println("length is: ", length)
	//used to print the length of the array .
    defer file.Close()

    readFile("./ashu.txt")
}
func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	if err != nil { log.Fatal(err) }
	

	fmt.Println("Text data inside the file is \n", string(databyte))

}