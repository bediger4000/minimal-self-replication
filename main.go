package main

import (
	"log"
	"minimal-self-replication/lexing"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("need an input file name on command line")
	}
	fileName := os.Args[1]

	buf, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lxr := lexing.NewCommandLexer([]rune(string(buf)))
	variables := make(map[string][]byte)

	for str := lxr(); str != ""; str = lxr() {
		lexing.Execute(str, variables)
	}
}
