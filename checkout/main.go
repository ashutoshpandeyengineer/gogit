package main
import (
	"fmt"
     "github.com/go-git/go-git/v5"
	"log"
	"gopkg.in/src-d/go-git.v4/plumbing")

func main() {
	r,err := git.PlainClone("./directory",false,&git.CloneOptions{URL : "url" ,})
	if err != nil {
		log.Fatal(err)
	}
    ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ref.Hash())

	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
    // ... checking out to commit
	err = w.Checkout(&git.CheckoutOptions{
		Hash:plumbing.NewHash (commit),
	})
	if err != nil {
		log.Fatal(err)
	}

	//... retrieving the commit being pointed by HEAD, it shows that the
	// repository is pointing to the giving commit in detached mode
	ref, err = r.Head()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ref.Hash())

	fmt.Println("hello")

		
	
}