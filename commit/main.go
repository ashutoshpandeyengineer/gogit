package main

import (
	"fmt"
	"path/filepath"

	"io/ioutil"
	"log"
//	"os"
	"time"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)
func main() {
	////opening a directory where we want our git file to be created .  /// here name of the directory is abc in the same folder.
	r,err := git.PlainOpen("./abc")
	if err != nil {
		log.Fatal(err)
	}


	w,err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
   //we need a file to commit so lets create a new file inside the worktree of the project to get commited .///////////////////////////////////////
    filename := filepath.Join("./abc", "example-git-file")
	err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)
	if err != nil {
		log.Fatal(err)
	}


	_, err = w.Add("example-git-file")          //adding the file to the staging Area.
	if err != nil {
		log.Fatal(err)
	}

    status, err := w.Status()                 // verifiying the status of worktree using method status.
    if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)

    commit, err := w.Commit("example go-git commit", &git.CommitOptions{Author: &object.Signature{Name:"John Doe", Email:"john@doe.org" ,When:time.Now(),},})
	///////////////////////  " message of commit " ////////////////////////////////////////////// {credentials required to commit .}//////////////////////////////////////////                         
	if err != nil {
		log.Fatal(err)
	}
    
   


	obj, err := r.CommitObject(commit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(obj)
	//prints the current HEAD to verify that the file is succesfully commited.
	
	
}