package main

import (
	"fmt"
	"os"
)

func translate(input string) string {
	translations := make(map[string]string)
	translations["en"] = "Hello"
	translations["es"] = "Hola"
	translations["de"] = "Guten Tag"
	translations["fr"] = "Bonjour"
	translations["ch"] = "Ni Hao"

	return translations[input]
}

func argswitch(input []string) string {
	var locale string
	if len(input) == 1 {
		fmt.Printf("What language do you want?: ")
		fmt.Scanln(&locale)
	} else {
		locale = os.Args[1]
	}

	return locale
}

func main() {
	locale := argswitch(os.Args)
	var available = []string{"English in en", "Spanish in es", "German in de", "French in fr", "Chinese in ch"}

	fmt.Println("Select a language")

	for i := 0; i < len(available); i++ {
		fmt.Println(available[i])
	}

	if len(locale) != 2 {
		fmt.Printf("Please input only two letters: ")
		fmt.Scanln(&locale)
	}

	output := translate(locale)

	if output == "" {
		output = "Yo"
	}

	fmt.Println(output, ", Go!")
}
