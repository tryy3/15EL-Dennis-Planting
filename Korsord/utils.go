package main

import "github.com/jroimartin/gocui"

func moveCursor(v *gocui.View) {
	_, y := v.Size()

	var prev string

	for i := 0; i < y; i++ {
		str, err := v.Line(i)
		if err != nil || str == "" {
			err = v.SetCursor(len(prev), i-1)
			if err != nil {
				panic(err)
			}
		}
		prev = str
	}
}
