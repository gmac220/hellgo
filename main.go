package main

import (
	"fmt"
	"os"
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

	if len(os.Args) == 1 {
		fmt.Printf("What language do you want?: ")
		fmt.Scanln(&locale)
	} else {
		locale = os.Args[1]
	}

	if len(locale) != 2 {
		fmt.Printf("Please input only two letters: ")
		fmt.Scanln(&locale)
	}

	languages["en"] = "Hello"
	languages["es"] = "Hola"
	languages["de"] = "Guten Tag"
	languages["fr"] = "Bonjour"
	languages["ch"] = "Ni Hao"

	output := languages[strings.ToLower(locale)]

	if output == "" {
		output = "Yo"
	}

	fmt.Println(output, ", Go!")
}
