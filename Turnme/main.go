package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func reverse(text string) string {
	// Konvertera strängen till en rune slice.
	// Så att man kan loopa över den.
	textRune := []rune(text)
	returnRune := []rune{}

	// Loopar igenom rune slicen baklänges och lägger till runen i den andra rune slicen.
	for i := len(textRune) - 1; i >= 0; i-- {
		returnRune = append(returnRune, textRune[i])
	}

	// Konvertera rune slicen till en string och returna den.
	return string(returnRune)
}

func main() {
	// Gör en ny Reader som läser ifrån os.Stdin
	// så att man kan få en sträng ifrån t.ex. CMD.
	buffer := bufio.NewReader(os.Stdin)

	// Skapa en regex som letar efter whitespaces.
	reg := regexp.MustCompile(`(\s)`)

	// Fråga efter en sträng och spara den i en variabel.
	fmt.Print("Skriv något: ")
	turn, err := buffer.ReadString('\n')

	// Kolla om det vart något error.
	if err != nil {
		panic(err)
	}

	// Ta bort alla mellanslag från strängen.
	noSpace := reg.ReplaceAllStringFunc(turn, func(s string) string {
		return ""
	})

	// Vänd på strängen.
	reversed := reverse(noSpace)

	// Kolla om strängen är ett anagram!
	if strings.ToLower(noSpace) == strings.ToLower(reversed) {
		fmt.Println("Strängen är en anagram!")
	} else {
		fmt.Println("Strängen är inte en anagram!")
	}
}
