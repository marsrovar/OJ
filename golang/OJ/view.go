package main

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
	"github.com/willf/pad"
)

func viewTitle(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("title", -1, -1, lMaxX, 1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.BgColor = gocui.ColorDefault | gocui.AttrReverse
		v.FgColor = gocui.ColorDefault | gocui.AttrReverse
	}

	return nil
}

func viewStatusBar(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("status", -1, lMaxY-2, lMaxX, lMaxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.BgColor = gocui.ColorBlack
		v.FgColor = gocui.ColorWhite
		// changeStatusContext(g, "D")
	}

	return nil
}

func viewEmpty(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("empty", -1, 1, lMaxX, lMaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

	return nil
}

func viewLogin(g *gocui.Gui, lMaxX int, lMaxY int) error {
	w := lMaxX / 4
	h := lMaxY / 4
	minX := (lMaxX / 2) - (w / 2)
	minY := (lMaxY / 2) - (h / 2)
	maxX := minX + w
	maxY := minY + h

	if v, err := g.SetView("login", minX, minY, maxX, maxY+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Configure view
		v.Title = " login "
		v.Frame = true
		fmt.Fprintln(v, " ")
		fmt.Fprintln(v, "Email")
		fmt.Fprintln(v, " ")
		fmt.Fprintln(v, " ")
		fmt.Fprintln(v, "Password")

		changeStatusContext(g, "LOGIN")
	}

	if v, err := g.SetView("loginEmail", minX+10, minY+1, maxX-5, maxY-5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		g.Cursor = true
		v.Editable = true
		g.SetCurrentView("loginEmail")

	}

	if v, err := g.SetView("loginPassword", minX+10, minY+4, maxX-5, maxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Editable = true
		v.Mask = '*'
	}

	if v, err := g.SetView("loginBtn", minX+2, minY+7, minX+8, maxY+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		fmt.Fprintln(v, "Login")
	}

	return nil
}

func loginAlert(g *gocui.Gui, lMaxX int, lMaxY int) error {
	w := lMaxX / 4
	h := lMaxY / 4
	minX := (lMaxX / 2) - (w / 2) + 2
	minY := (lMaxY / 2) - (h / 2) - 3
	maxX := minX + w - 4
	maxY := minY

	if v, err := g.SetView("alert", minX, minY, maxX, maxY+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
	}

	return nil
}

func viewQPlist(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("qplist", -1, 1, lMaxX, lMaxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		v.SetCursor(0, 2)
		// go viewPodsShowWithAutoRefresh(g)
	}

	return nil
}

// var c1 = make(chan string)

// // Auto refresh view pods
// func viewPodsShowWithAutoRefresh(g *gocui.Gui) {
// 	// c := getConfig()
// 	t := time.NewTicker(time.Second * 1)
// 	// go viewPodsRefreshList(g)
// 	for {
// 		select {
// 		case <-t.C:
// 			debug(g, fmt.Sprintf("View pods: Refreshing (%ds)", "a"))
// 		}
// 	}
// }

func viewQPlistAddLine(v *gocui.View, maxX int, num, name, title, status string) {
	wN := maxX - 34
	if wN < 45 {
		wN = 45
	}
	line := pad.Right(num, 5, " ") +
		pad.Right("|", 1, " ") +
		pad.Right(name, 8, " ") +
		pad.Right("|", 1, " ") +
		pad.Right(title, wN+5, " ") +
		pad.Right("|", 1, " ") +
		pad.Right(status, 15, " ")
	fmt.Fprintln(v, line)
}

func viewCode(g *gocui.Gui, lMaxX int, lMaxY int) error {
	if v, err := g.SetView("code", 2, 2, int((lMaxX-4)/2)+1, lMaxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Code "
		// v.Autoscroll = true
		v.Editable = true
		v.Wrap = true
	}

	minX := int(lMaxX / 2)
	minY := 2
	maxX := lMaxX - 4
	maxY := int(lMaxY / 2)
	if v, err := g.SetView("content", minX, minY, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Content "
		v.Wrap = true
		v.Highlight = true
	}

	if v, err := g.SetView("example", minX, maxY, lMaxX-4, lMaxY-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Example "
		v.Wrap = true
	}

	return nil
}

func viewCodeTool(g *gocui.Gui, lMaxX int, lMaxY int) error {
	w := lMaxX / 4
	h := lMaxY / 4
	minX := (lMaxX / 2) - (w / 2)
	minY := (lMaxY / 2) - (h / 2)
	maxX := minX + w
	maxY := minY + h

	if v, err := g.SetView("codetool", minX, minY, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Run Code")
		fmt.Fprintln(v, "Submit  Solution")
		fmt.Fprintln(v, "Reset")
		fmt.Fprintln(v, "Exit")
	}

	return nil
}

func changeStatusContext(g *gocui.Gui, c string) error {
	lMaxX, _ := g.Size()
	v, err := g.View("status")
	if err != nil {
		return err
	}

	v.Clear()
	v1, err := g.View("title")
	if err != nil {
		return err
	}

	v1.Clear()
	i := lMaxX + 4
	b := ""
	ab := ""
	switch c {
	case "LOGIN":
		i = 70 + i
		b = b + frameText("Tab") + " Exchange   "
		b = b + frameText("Enter") + " Login   "
	case "D":
		i = 100 + i
		b = b + frameText("↑") + " Up   "
		b = b + frameText("↓") + " Down   "
		b = b + frameText("Enter") + " Show Question   "
	case "SE":
		i = i + 125
		b = b + frameText("↑") + " Up   "
		b = b + frameText("↓") + " Down   "
		b = b + frameText("Enter") + " Select   "
		b = b + frameText("CTRL+N") + " Hide Code Tool   "
	case "SL":
		p, err := getSelectedQPlist(g)
		if err != nil {
			return err
		}
		num, err := strconv.Atoi(p)
		ab = QP[num-1].Title
		i = i + 150
		b = b + frameText("↑") + " Up   "
		b = b + frameText("↓") + " Down   "
		b = b + frameText("CTRL+L") + " Hide Question   "
		b = b + frameText("CTRL+O") + " New Line   "
		b = b + frameText("CTRL+N") + " Show Code Tool   "
	}
	b = b + frameText("CTRL+C") + " Exit"

	fmt.Fprintln(v, pad.Left(b, i, " "))
	fmt.Fprintln(v1, "    "+ab)

	return nil
}
