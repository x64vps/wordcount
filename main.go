package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	argsWithProg := os.Args[1:]

	if argsWithProg[0] == "" {
		fmt.Println(0)
		return
	}

	parts := strings.Split(argsWithProg[0], " ")

	fmt.Println(len(parts))
}
