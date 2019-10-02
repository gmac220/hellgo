package main

import (
	"fmt"
	"os"
	// "strings"
)

var translations map[string]string

func translate(input string) string {
	return translations[input]
}

func argswitch(input []string) string {
	var locale string
	if len(input) == 1 {
		help()
	} else {
		locale = os.Args[1]
	}

	return locale
}

func help() {
	fmt.Println("usage: hellgo args1")
	fmt.Println("args: ")

	for index := range translations {
		fmt.Println(index)
	}

	os.Exit(0)
}

func main() {
	translations = make(map[string]string)
	translations["en"] = "Hello"
	translations["es"] = "Hola"
	translations["de"] = "Guten Tag"
	translations["fr"] = "Bonjour"
	translations["ch"] = "Ni Hao"

	// input := argswitch(os.Args)
	// locale := strings.ToLower(input)
	locale := argswitch(os.Args)

	output := translate(locale)

	if output == "" {
		output = "Yo"
	}

	fmt.Println(output, ", Go!")
}
