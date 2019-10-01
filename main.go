package main

import (
	"fmt"
)

func main() {
	//var locale, greeting string
	var locale string
	// var languages = [4]string{"en", "es", "de", "fr"}/
	languages := make(map[string]string)
	// locale = languages[0]

	fmt.Printf("What language do you want? (specify two letters in lowercase) EX. en for English: ")
	// reader := bufio.NewReader(os.Stdin)
	// locale, _ := reader.ReadString('\n')
	fmt.Scanln(&locale)

	languages["en"] = "Hello"
	languages["es"] = "Hola"
	languages["de"] = "Guten Tag"
	languages["fr"] = "Bonjour"
	languages["ch"] = "Ni Hao"

	fmt.Println(languages[locale] + ", Go!")

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
