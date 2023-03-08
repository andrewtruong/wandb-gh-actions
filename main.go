package main

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	cp "github.com/otiai10/copy"
)

func main() {
	// Load Repo
	r, err := git.PlainOpen(".")
	checkErr(err)

	// Create Branch
	ref, err := createBranch(r)
	checkErr(err)

	// Get Worktree
	w, err := r.Worktree()
	checkErr(err)

	// Checkout Branch
	err = w.Checkout(&git.CheckoutOptions{Branch: ref.Name()})
	checkErr(err)

	// Copy Files
	pwd, _ := os.Getwd()
	src := filepath.Clean(filepath.Join(pwd, "test"))
	dst := filepath.Clean(filepath.Join(pwd, ".github/workflows"))
	err = cp.Copy(src, dst)
	checkErr(err)

	// Add Files
	_, err = w.Add(".github/workflows")
	checkErr(err)

	// Push
	auth, _ := ssh.NewSSHAgentAuth("git")
	err = r.Push(&git.PushOptions{
		Auth: auth,
	})
	checkErr(err)

}
func checkErr(err error) {
	if err != nil {
		// log.Fatal("Error: ", err)
		panic(err)
	}
}
func createBranch(r *git.Repository) (*plumbing.Reference, error) {
	branch := plumbing.NewBranchReferenceName("test")
	headRef, err := r.Head()
	checkErr(err)

	ref := plumbing.NewHashReference(branch, headRef.Hash())
	err = r.Storer.SetReference(ref)
	checkErr(err)

	return ref, err
}
