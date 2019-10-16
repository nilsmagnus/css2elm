package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
	"github.com/iancoleman/strcase"
)

func main() {

	fileName := flag.String("input", "", "The css-file to generate style-functions from")
	functionPrefix := flag.String("prefix", "view", "Prefix for the style-functions")

	flag.Parse()

	if *fileName == "" {
		log.Fatal("'input' must be defined")
	}

	fileContents, err := ioutil.ReadFile(*fileName)

	if err != nil {
		log.Fatal("Could not read file")
	}

	cleanCss, err := parser.Parse(string(fileContents))

	if err != nil {
		log.Fatalf("Could not parse css file: %s", err.Error())
	}

	fmt.Printf("module ElmFromCss exposing(..)\n")
	fmt.Printf("import Html exposing (Attribute)\n")
	fmt.Printf("import Html.Attributes exposing (style)\n")
	for _, v := range cleanCss.Rules {
		fmt.Printf("\n\n%s\n\n", elmStyleFunction(*functionPrefix, v))
	}

}

func elmStyleFunction(functionPrefix string, v *css.Rule) string {
	var sb strings.Builder
	cleanName := strings.ReplaceAll(v.Prelude, ".", "")
	functionName := fmt.Sprintf("%s%s", functionPrefix, strcase.ToCamel(cleanName))

	sb.WriteString(fmt.Sprintf("%s: List (Attribute msg)\n", functionName))
	sb.WriteString(fmt.Sprintf("%s = \n", functionName))
	for k, component := range v.Declarations {
		separator := "    ,"
		if k == 0 {
			sb.WriteString(fmt.Sprintf("    ["))
			separator = ""
		}
		sb.WriteString(fmt.Sprintf("%sstyle \"%s\" \"%s\"\n", separator, component.Property, component.Value))
	}

	sb.WriteString(fmt.Sprintf("    ]\n"))

	return sb.String()
}
