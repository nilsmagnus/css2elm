package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gorilla/css/scanner"
)

func main() {
	fmt.Println("hei")

	fileName := flag.String("input", "", "")

	flag.Parse()

	if *fileName == "" {
		log.Fatal("'input' must be defined")
	}

	fileContents, err := ioutil.ReadFile(*fileName)

	if err != nil {
		log.Fatal("Could not read file")
	}

	s := scanner.New(string(fileContents))

	for {
		token := s.Next()
		if token.Type == scanner.TokenEOF || token.Type == scanner.TokenError {
			break
		}

	}
}
