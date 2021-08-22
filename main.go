package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main()  {
	argsWithProg := os.Args[1:]

	parts := strings.Split(argsWithProg[0], " ")

	fmt.Println(len(parts))
}

// readString считывает и возвращает строку из os.Stdin
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}
