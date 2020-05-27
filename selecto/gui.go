package selecto

import (
	"fmt"
	"os"

	"github.com/jroimartin/gocui"
)

const (
	BORDER_SIZE = 2
)

var DEBUG = os.Getenv("DEBUG") // turn on view debug info

// Move cursor up
// When cursor is near to top border it scrolls Origin and stay few lines before the border
func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}

	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if DEBUG != "" {
		v.Title = title(v)
	}


	if oy > 0 && cy <= BORDER_SIZE {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			if err = v.SetCursor(cx, cy-1); err != nil {
				return fmt.Errorf("cannot set cursor %w", err)
			}
			return fmt.Errorf("cannot set origin %w", err)
		}

		return nil
	}

	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func title(v *gocui.View) string {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	sx, sy := v.Size()
	bl := len(v.BufferLines())
	vbl := len(v.ViewBufferLines())

	return fmt.Sprintf("Select... (wh(%d, %d), origin(%d,%d), cur(%d, %d), buf(%d), vbuf(%d))", sx, sy, ox,oy, cx,cy, bl, vbl)
}

// Move cursor down
// When cursor is near to bottom border it scrolls Origin and stay few lines before the border
func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v == nil {
		return nil
	}

	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	_, sy := v.Size()

	if DEBUG != "" {
		v.Title = title(v)
	}

	// get text in next line
	text, err := v.Line(cy+1)
	if err != nil {
		return fmt.Errorf("vLine %w", err)
	}
	if text == "" {
		return nil
	}

	if cy >= sy - BORDER_SIZE {
		if err = v.SetOrigin(ox, oy+1); err != nil {
			if err = v.SetCursor(cx, cy+1); err != nil {
				return fmt.Errorf("cannot set cursor %w", err)
			}

			return fmt.Errorf("cannot set origin %w", err)
		}

		return nil
	}

	if err := v.SetCursor(cx, cy+1); err != nil {
		return fmt.Errorf("cannot set cursor %w", err)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
