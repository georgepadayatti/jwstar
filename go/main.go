package main

import (
	"fmt"

	"github.com/georgepadayatti/jwstar/go/cmd"
)

func main() {

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
