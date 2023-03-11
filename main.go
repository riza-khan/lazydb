package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	// Create a new gocui view
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '1', gocui.ModNone, view1); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("view1", 'g', gocui.ModNone, view1Print); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '2', gocui.ModNone, view2); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("view2", 'g', gocui.ModNone, view2Print); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// maxX, maxY := g.Size()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	// Create a new view with the name "myView"
	if v, err := g.SetView("view1", 0, 0, 20, 10); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}

		// Set the default background color of the view to nothing
		v.Title = "View 1"
		v.Wrap = false

		if _, err := g.SetCurrentView("view1"); err != nil {
			return err
		}
		// fmt.Fprintln(v, "View 1")
	}

	// Set a second view with the name "myOtherView"
	if v, err := g.SetView("view2", 25, 0, 45, 10); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}

		v.Title = "View 2"
		v.Wrap = false
	}

	return nil
}

func view1(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView("view1"); err != nil {
		return err
	}

	return nil
}

func view1Print(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(v, "Foobar")

	return nil
}

func view2Print(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintln(v, "Goobar")

	return nil
}

func view2(g *gocui.Gui, v *gocui.View) error {
	if _, err := g.SetCurrentView("view2"); err != nil {
		return err
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
