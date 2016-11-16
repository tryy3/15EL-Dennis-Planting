package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// stringInSlice tar emot en string och en string slice
// loopar igenom string slicen och sen kollar om
// str är med i list
// I java så skulle detta vara samma sak som att köra
// #Arrays::Contains
func stringInSlice(str string, list []string) bool {
	str = strings.ToUpper(str) // Jag la till denna linje då jag inte bryr mig om Lowercase/uppercase.
	for _, tecken := range list {
		// Kolla om tecken (varje element i list) är likadant som str.
		// Om det är likadant, return true, annars fortsätt.
		//fmt.Println(tecken, str, tecken == str)
		if str == tecken {
			return true
		}
	}
	// Om funktionen inte har slutat än, return false.
	return false
}

// Function main är en inbyggd funktion, liknande main funtionen i Java.
func main() {
	// Initialisera variabler
	var rövarspråk string               // Kommer att innehålla slut resultatet
	buffer := bufio.NewReader(os.Stdin) // Gör en new Reader på os.Stdin (i princip, gör en ny reader på kommando konsolen)
	konsonanter := []string{"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Z"}

	// Fråga efter en sträng, sen spara svaret i inteRövarspråk.
	fmt.Print("Skriv en sträng: ")
	inteRövarspråk, err := buffer.ReadString('\n')

	// Kolla om det var något error, borde inte vara något, men kan alltid vara bra att kolla.
	if err != nil {
		// Om det finns ett error, gör en panic på det
		panic(err)
	}

	// Gör inteRövarspråk till en slice efter varje tecken (inklusive mellanrum)
	inteRövarspråkSlice := strings.Split(inteRövarspråk, "")

	// Loopa igenom varje tecken från slicen.
	for _, tecken := range inteRövarspråkSlice {
		// Kolla om tecknet är en ny linje om det finns, hoppa över den.
		if tecken == "\n" || tecken == "\r" {
			continue
		}
		if stringInSlice(tecken, konsonanter) {
			rövarspråk += tecken + "o" + strings.ToLower(tecken)
		} else {
			rövarspråk += tecken
		}
	}
	fmt.Println("Rövarspråk:", rövarspråk)
}
