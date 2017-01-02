package main

import (
	"errors"
	"fmt"
	"strconv"

	"strings"

	"unicode"

	"unicode/utf8"

	"github.com/jroimartin/gocui"
)

type Editor struct {
	Frågor []*Fråga
	Nummer int
	Buf    string
}

type Fråga struct {
	Text  func(*Editor) string
	Check func(string) error
	Svar  func(string)
}

func (e *Editor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		e.Buf += string(ch)
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		e.Buf += string(ch)
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		if e.Buf == "" {
			return
		}
		e.Buf = e.Buf[:len(e.Buf)-1]
		v.EditDelete(true)
		// ...
	case key == gocui.KeyEnter:
		if e.Nummer >= len(e.Frågor) {
			// TODO: Gör något här, typ skriv något error eller bara lägg till en new line.
			return
		}

		fråga := e.Frågor[e.Nummer]
		err := fråga.Check(e.Buf)
		if err != nil {
			v.Clear()
			if err.Error() != "" {
				fmt.Fprintln(v, err)
			}
			fmt.Fprintln(v, fråga.Text(e))
			err = moveCursor(v)
			if err != nil {
				panic(err)
			}
			e.Buf = ""
			return
		}

		fråga.Svar(e.Buf)
		e.Nummer++
		e.Buf = ""
		v.Clear()
		if e.Nummer >= len(e.Frågor) {
			fmt.Fprintln(v, "Inga fler frågor!")
			return
		}
		fmt.Fprintln(v, e.Frågor[e.Nummer].Text(e))
		err = moveCursor(v)
		if err != nil {
			panic(err)
		}
	}
}

func init() {
	DefaultEditor = &Editor{
		Frågor: []*Fråga{
			{
				func(*Editor) string { return "Var god och skriv ut hur många kolumer layouten ska ha:" },
				func(str string) error {
					if _, err := strconv.Atoi(strings.Trim(str, " ")); err != nil {
						return errors.New("Det verkar som antalet du angav inte är ett nummer, var god och försök igen.")
					}
					return nil
				},
				func(str string) {
					WIDTH, _ = strconv.Atoi(strings.Trim(str, " "))
				},
			},
			{
				func(*Editor) string { return "Var god och skriv ut hur många rader layouten ska ha:" },
				func(str string) error {
					if _, err := strconv.Atoi(strings.Trim(str, " ")); err != nil {
						return errors.New("Det verkar som antalet du angav inte är ett nummer, var god och försök igen.")
					}
					return nil
				},
				func(str string) {
					HEIGHT, _ = strconv.Atoi(strings.Trim(str, " "))
					gui.Execute(func(g *gocui.Gui) error {
						for y := 0; y < HEIGHT; y++ {
							var kolumn []*Tecken
							for x := 0; x < WIDTH; x++ {
								kolumn = append(kolumn, nil)
							}
							LAYOUT = append(LAYOUT, kolumn)
						}

						if err := updateList(g); err != nil {
							return err
						}
						if err := updateMain(g); err != nil {
							return err
						}
						return nil
					})
				},
			},
			{
				func(*Editor) string { return "Var god och skriv antalet siffror du vill ha mellan 1, 27: " },
				func(str string) error {
					num, err := strconv.Atoi(strings.Trim(str, " "))
					if err != nil {
						return errors.New("Det verkar som antalet du angav inte är ett nummer, var god och försök igen.")
					}

					if num <= 0 {
						return errors.New("Antalet du angav är mindre än 0.")
					}
					if num > 27 {
						return errors.New("Antalet du angav är högre än 27.")
					}
					return nil
				},
				func(str string) {
					num, _ := strconv.Atoi(strings.Trim(str, " "))

					for i := 1; i <= num; i++ {
						LISTA = append(LISTA, &Tecken{uint8(i), ""})
					}

					gui.Execute(func(g *gocui.Gui) error {
						return updateList(g)
					})
				},
			},
			{
				func(*Editor) string {
					return "Var god och skriv ut en siffra och x,y position, separera det med en komma: "
				},
				func(str string) error {
					l := strings.Split(str, ",")

					if len(l) < 3 {
						return errors.New("Det verkar som du har skrivit in för få argument, var god och försök igen.")
					}

					if len(l) > 3 {
						return errors.New("Det verkar som du har skrivit in för många argument, var god och försök igen.")
					}

					siffra, err := strconv.Atoi(strings.Trim(l[0], " "))

					if err != nil {
						return errors.New("Det verkar som antalet du angav inte är ett nummer, var god och försök igen.")
					}

					if siffra > len(LISTA) {
						return errors.New("Siffran du angav är högre än längden på Listan.")
					}

					if siffra < 1 {
						return errors.New("Siffran du angav är mindre än 1.")
					}

					x, err := strconv.Atoi(strings.Trim(l[1], " "))

					if err != nil {
						return errors.New("Det verkar som antalet du angav inte är ett nummer, var god och försök igen.")
					}

					if x < 0 {
						return errors.New("X position kan inte vara mindre än 0.")
					}

					if x >= WIDTH {
						return errors.New("X position kan inte vara högre än WIDTH.")
					}

					y, err := strconv.Atoi(strings.Trim(l[2], " "))

					if err != nil {
						return errors.New("Det verkar som antalet du angav inte är ett nummer, var god och försök igen.")
					}

					if y < 0 {
						return errors.New("y position kan inte vara mindre än 0.")
					}

					if y >= HEIGHT {
						return errors.New("y position kan inte vara högre än HEIGHT.")
					}

					LAYOUT[y][x] = LISTA[siffra-1]
					gui.Execute(func(g *gocui.Gui) error {
						return updateMain(g)
					})

					for y := 0; y < len(LAYOUT); y++ {
						for x := 0; x < len(LAYOUT[y]); x++ {
							if LAYOUT[y][x] == nil {
								return errors.New("")
							}
						}
					}
					return nil
				},
				func(str string) {
					return
				},
			},
			{
				func(*Editor) string {
					return "Var god och skriv ut en siffra och en bokstav, separera det med en komma: "
				},
				func(str string) error {
					l := strings.Split(str, ",")

					if len(l) < 2 {
						return errors.New("Det verkar som du har skrivit in för få argument, var god och försök igen.")
					}

					if len(l) > 2 {
						return errors.New("Det verkar som du har skrivit in för många argument, var god och försök igen.")
					}

					siffra, err := strconv.Atoi(strings.Trim(l[0], " "))

					if err != nil {
						return errors.New("Det verkar som antalet du angav inte är ett nummer, var god och försök igen.")
					}

					if siffra > len(LISTA) {
						return errors.New("Siffran du angav är högre än längden på Listan.")
					}

					if siffra < 1 {
						return errors.New("Siffran du angav är mindre än 1.")
					}

					r, antal := utf8.DecodeRuneInString(l[1])
					if antal != len(l[1]) {
						return errors.New("Du kan bara skriva en bokstav inte flera, var god och försök igen.")
					}
					if !unicode.IsLetter(r) {
						return errors.New("Du kan bara skriva en bokstav, inte en symbol, var god och försök igen.")
					}

					LISTA[siffra-1].Bokstav = strings.ToUpper(l[1])
					gui.Execute(func(g *gocui.Gui) error {
						if err := updateList(g); err != nil {
							return err
						}
						if err := updateMain(g); err != nil {
							return err
						}
						return nil
					})
					return errors.New("")
				},
				func(str string) {

				},
			},
		},
	}
}
