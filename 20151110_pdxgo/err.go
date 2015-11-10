package main

import "fmt"
import "errors"

type gitError struct {
	repo, commit string
}

func (e gitError) Error() string {
	return fmt.Sprintf("Error in repo (%s) with commit %s", e.repo, e.commit)
}

func a() error { return gitError{repo: "github.com...", commit: "123"} }
func b() error { return errors.New("Error in repo (github.com...) with commit 123") }

func main() {
	err := a()
	_, ok := err.(gitError)
	fmt.Println(err)
	fmt.Println(ok)

	err = b()
	_, ok = err.(gitError)
	fmt.Println(err)
	fmt.Println(ok)
}
