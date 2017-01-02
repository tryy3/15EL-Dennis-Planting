package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/huandu/xstrings"
	"github.com/jroimartin/gocui"
)

type Tecken struct {
	Siffra  uint8
	Bokstav string
}

var (
	DefaultEditor *Editor
)

var (
	WIDTH   = -1
	HEIGHT  = -1
	LAYOUT  = [][]*Tecken{}
	LISTA   = []*Tecken{}
	viewArr = []string{"viewList", "viewCMD"}
	active  = 1
	gui     *gocui.Gui
)

func updateList(g *gocui.Gui) error {
	v, err := g.View("viewList")
	v.Clear()

	if err != nil {
		return err
	}

	if HEIGHT != -1 && WIDTH != -1 {
		fmt.Fprintf(v, "WIDTH: %d\n", WIDTH)
		fmt.Fprintf(v, "HEIGHT: %d\n\n", HEIGHT)
	}

	if len(LISTA) > 0 {
		for i := 0; i < len(LISTA); i++ {
			tecken := LISTA[i]
			bokstav := tecken.Bokstav
			if bokstav == "" {
				bokstav = "Null"
			}
			fmt.Fprintf(v, "%d: %s\n", tecken.Siffra, bokstav)
		}
	}
	return nil
}

func updateMain(g *gocui.Gui) error {
	v, err := g.View("viewMain")
	v.Clear()

	if err != nil {
		return err
	}

	if HEIGHT != -1 && WIDTH != -1 {
		for y := 0; y <= HEIGHT+1; y++ {
			for x := 0; x <= WIDTH; x++ {
				if y == 0 || y == HEIGHT+1 {
					fmt.Fprint(v, "#  ")
					continue
				}
				str := "#"
				if x == WIDTH {
					fmt.Fprint(v, xstrings.LeftJustify(str, 3, " "))
					continue
				}
				t := LAYOUT[y-1][x]
				if x != 0 {
					str = " "
				}

				if t == nil {
					fmt.Fprint(v, xstrings.LeftJustify(str, 3, " "))
					continue
				}

				if t.Bokstav != "" {
					str += t.Bokstav
				} else {
					str += strconv.Itoa(int(t.Siffra))
				}
				fmt.Fprint(v, xstrings.LeftJustify(str, 3, " "))
			}
			fmt.Fprint(v, "\n")
		}
	}
	fmt.Fprintf(v, "%#v", LAYOUT)
	return nil
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	if active == 0 {
		active = 1
	} else {
		active = 0
	}

	if _, err := setCurrentViewOnTop(g, viewArr[active]); err != nil {
		return err
	}
	g.Cursor = true
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func layout(g *gocui.Gui) error {
	gui = g
	maxX, maxY := g.Size()
	if v, err := g.SetView("viewMain", 0, 0, int(0.85*float32(maxX)), int(0.9*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Main layout"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("viewList", int(0.85*float32(maxX))+1, 0, maxX-1, int(0.9*float32(maxY))); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Lista"
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true
	}
	if v, err := g.SetView("viewCMD", 0, int(0.9*float32(maxY))+1, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Terminal"
		v.Wrap = true
		v.Autoscroll = true
		v.Editable = true
		v.Editor = DefaultEditor

		fmt.Fprintln(v, DefaultEditor.Frågor[0].Text(DefaultEditor))
		err := moveCursor(v)
		if err != nil {
			return err
		}

		if _, err = setCurrentViewOnTop(g, "viewCMD"); err != nil {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
