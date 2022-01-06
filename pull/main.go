package main 
import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"log"
	

)

func main() {
    r,err := git.PlainOpen("./directory")
	if err != nil {log.Fatal(err)}


	w,err := r.Worktree()
	if err != nil {log.Fatal(err)}
	

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {log.Fatal(err)}
    fmt.Println("Hello")

	ref, err := r.Head()
	if err != nil {log.Fatal(err)}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {log.Fatal(err)}

	fmt.Println(commit)
    
	 

}