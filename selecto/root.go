package selecto

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/jroimartin/gocui"
)

var (
	KeyQ = rune(113)
	KeyJ = rune(106)
	KeyK = rune(107)
)

type Result struct {
	Line *string
	Error error
}

type Selecto struct {
	source chan string
	result chan Result
	gui *gocui.Gui
	selected bool
}

func NewSelecto(r io.Reader) (*Selecto, error) {
	scanner := bufio.NewScanner(r)
	source := make(chan string, 20)
	go func() {
		for scanner.Scan() {
			source <- scanner.Text()
		}
	}()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, fmt.Errorf("cannot create gui %w", err)
	}

	s := &Selecto{
		source: source,
		result: make(chan Result, 1),
		gui: g,
	}

	g.SetManagerFunc(s.layout)

	if err = s.keybinding(g); err != nil {
		return nil, fmt.Errorf("cannot set keybindings %w", err)
	}

	return s, nil
}

func (s *Selecto) WaitForSelect() *Result {
	defer s.gui.Close()
	err := s.gui.MainLoop()
	if err != nil && errors.Is(gocui.ErrQuit, err) && !s.selected {
		nothing := ""
		return &Result {
			Line: &nothing,
		}
	}
	if err != nil && !errors.Is(gocui.ErrQuit, err) {
		return &Result{
			Line: nil,
			Error: fmt.Errorf("stopped with %w", err),
		}
	}

	result := <- s.result
	return &result
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
		v.Title = "Select..."

		go func() {
			for line := range s.source {
				if line == "" { continue }
				fmt.Fprintln(v, line)
				g.Update(func(g *gocui.Gui) error { return nil })
			}
		}()

		if _, err = g.SetCurrentView("list"); err != nil {
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

	s.selected = true
	s.result <- Result{
		Line: &line,
	}

	return quit(g, v)
}

func (s *Selecto) keybinding(g *gocui.Gui) error {
	if err := g.SetKeybinding("list", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", KeyQ, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", gocui.KeyEnter, gocui.ModNone, s.getLine); err != nil {
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

	return nil
}
