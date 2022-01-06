package main

import (
	"log"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/storage/memory"
)
func main() {
	//create a new repository
	//"git init"
	r,err := git.Init(memory.NewStorage(),nil)
	if err != nil {
		log.Fatal(err)
	}
	// Add a new remote, with the default fetch refspec
	//Info("git remote add example https://github.com/git-fixtures/basic.git")
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "example",
		URLs: []string{"https://github.com/git-fixtures/basic.git"},
	})
	if err != nil {
		log.Fatal(err)
	}
	// List remotes from a repository
	// "git remotes -v"

	list, err := r.Remotes()
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range list {
		fmt.Println(r)
	}

	// Fetch using the new remote
	// "git fetch example"
	err = r.Fetch(&git.FetchOptions{
		RemoteName: "example",
	})
	if err != nil {
		log.Fatal(err)
	}
    // Delete the example remote
	//"git remote rm example

	err = r.DeleteRemote("example")
	if err != nil {
		log.Fatal(err)
	}
}