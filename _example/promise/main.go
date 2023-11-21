package main

import (
	"fmt"

	gopromise "github.com/agustin-del-pino/go-promises/pkg/go-promise"
)



func main() {
	p := gopromise.New(func() (string, error) {
		return "Hello, World!", nil
	})

	r, err := p.Await()

	fmt.Printf("value: %s\nerror: %s", r, err)
}