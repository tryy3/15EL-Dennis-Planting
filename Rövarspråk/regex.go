package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Function main är en inbyggd funktion, liknande main funtionen i Java.
func main() {
	buffer := bufio.NewReader(os.Stdin) // Gör en new Reader på os.Stdin (i princip, gör en ny reader på kommando konsolen)

	// Initialisera regex som letar efter konsonanter.
	reg := regexp.MustCompile("((?i)[b-df-gj-np-tv-xz])")

	// Fråga efter en sträng och spara det i variabeln input
	fmt.Print("Skriv en sträng: ")
	input, err := buffer.ReadString('\n')

	// Kolla om det var något error, borde inte vara något, men kan alltid vara bra att kolla.
	if err != nil {
		// Om det finns ett error, gör en panic på det
		panic(err)
	}

	// Byt ut varje matchning i regex med match grupp 1 plus ett o plus match grupp 1
	// I princip, leta efter varje konsonant, när den har hittat ett
	// Byt ut tecknet mot tecken + o + tecken
	// Sen skriv ut strängen
	fmt.Println("Rövarspråk:", reg.ReplaceAllStringFunc(input, func(s string) string {
		return s + "o" + strings.ToLower(s)
	}))
}
