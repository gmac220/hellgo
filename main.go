package main

import (
	"fmt"
	"strings"
)

func main() {
	var locale string
	var available = []string{"English in en", "Spanish in es", "German in de", "French in fr", "Chinese in ch"}
	languages := make(map[string]string)

	fmt.Println("Select a language")

	for i := 0; i < len(available); i++ {
		fmt.Println(available[i])
	}

	fmt.Printf("What language do you want? EX. en for English: ")
	fmt.Scanln(&locale)

	if len(locale) != 2 {
		fmt.Printf("Please input only two letters: ")
		fmt.Scanln(&locale)
	}

	languages["en"] = "Hello"
	languages["es"] = "Hola"
	languages["de"] = "Guten Tag"
	languages["fr"] = "Bonjour"
	languages["ch"] = "Ni Hao"

	if locale == "" {
		fmt.Println("Yo, Go!")
	} else {
		fmt.Println(languages[strings.ToLower(locale)] + ", Go!")
	}

	/*
		if locale == "en" {
			greeting = "Hello"
		} else if locale == "es" {
			greeting = "Hola"
		} else if locale == "de" {
			greeting = "Guten Tag"
		} else {
			greeting = "Yo"
		}
	*/

	/*
		switch locale {
		case "en":
			greeting = "Hello"
		case "es":
			greeting = "Hola"
		case "de":
			greeting = "Guten Tag"
		case "fr":
			greeting = "Bonjour"
		default:
			greeting = "Yo"
		}
	*/

	//fmt.Printf(greeting + ", Go!\n")

}
