package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

var (
	KeyQ = rune(113)
	KeyJ = rune(106)
	KeyK = rune(107)
)

type Selecto struct {
	lines        []string
	selectedLine string
}

func NewSelecto(lines []string) *Selecto {
	return &Selecto{
		lines: lines,
	}
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}

	cx, cy := v.Cursor()
	text, err := v.Line(cy + 1)
	if err != nil {
		return err
	}
	if text == "" {
		return nil
	}

	if err := v.SetCursor(cx, cy+1); err != nil {
		return err
	}

	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}

	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func (s *Selecto) getLine(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	line, err := v.Line(cy)
	if err != nil {
		return err
	}

	s.selectedLine = line

	return gocui.ErrQuit
}

func (s *Selecto) keybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, s.quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", KeyQ, gocui.ModNone, s.quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", KeyJ, gocui.ModNone, cursorDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", KeyK, gocui.ModNone, cursorUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", gocui.KeyEnter, gocui.ModNone, s.getLine); err != nil {
		return err
	}

	return nil
}

func (s *Selecto) Start() (string, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.SetManagerFunc(s.layout)

	if err = s.keybinding(g); err != nil {
		log.Panicln(err)
	}

	err = g.MainLoop()
	if err != nil && !errors.Is(gocui.ErrQuit, err) {
		log.Panicln(err)
	}

	return s.selectedLine, nil
}

func (s *Selecto) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	g.Cursor = false
	if v, err := g.SetView("list", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.FgColor = gocui.ColorGreen
		v.Title = "Select row..."

		for _, line := range s.lines {
			if line == "" {
				continue
			}

			fmt.Fprintln(v, line)
		}

		if _, err = g.SetCurrentView("list"); err != nil {
			log.Panicln(err)
		}
	}

	return nil
}

func (s *Selecto) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	selecto := NewSelecto(lines)

	selected, err := selecto.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(selected)
}
