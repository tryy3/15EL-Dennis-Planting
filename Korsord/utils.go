package main

import "github.com/jroimartin/gocui"

func moveCursor(v *gocui.View) error {
	_, y := v.Size()

	var prev string

	for i := 0; i < y; i++ {
		str, err := v.Line(i)
		if err != nil || str == "" {
			return v.SetCursor(len(prev), i-1)
		}
		prev = str
	}
	return nil
}
