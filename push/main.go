package main
import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	

)

func main() {
	//opening the directory to perform git operation.
	r,err := git.PlainOpen("./directory")
	if err != nil {
		log.Fatal(err)
	}
	// By default it will push to the remote origin.
    err = r.Push(&git.PushOptions{})  //push command declared here 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully  Done")
}