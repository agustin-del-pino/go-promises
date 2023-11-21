package main

import (
	"fmt"

	gopromise "github.com/agustin-del-pino/go-promises/pkg/go-promise"
)

func CreatePromise(i int) gopromise.Promise[int] {
	return gopromise.New(func() (int, error) {
		return i, nil
	})
}

func main() {
	p := make([]gopromise.Promise[int], 5)
	
	for i := range p {
		p[i] = CreatePromise(i)
	}

	r, err := gopromise.AllSettled(p).Await()

	fmt.Printf("value: %v\nerror: %s", r, err)
}