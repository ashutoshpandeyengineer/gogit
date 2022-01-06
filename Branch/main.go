package main

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)
func main() {
	r,err := git.PlainClone("./directory",false,&git.CloneOptions{
		URL: "url" ,
	})
    if err != nil {
		log.Fatal(err)
	}
    // getting the head of the tree .
	headRef, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}

	//Create a new plumbing.HashReference object with the name of the branch and the hash from the HEAD. The reference name should be a full reference
	ref :=plumbing.NewHashReference("refs/heads/my-branch",headRef.Hash())

	// The created reference is saved in the storage.
	err = r.Storer.SetReference(ref)
	if err != nil {
		log.Fatal(err)
	}
	// Or deleted from it.
	err = r.Storer.RemoveReference(ref.Name())
	if err != nil {
		log.Fatal(err)
	}
}