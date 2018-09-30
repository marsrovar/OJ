package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/sacOO7/gowebsocket"
)

var SOCKET = gowebsocket.New("ws://127.0.0.1:1121/")

var SELECTED_POD_NUM = 0

func moveViewCursorDown(g *gocui.Gui, v *gocui.View, allowEmpty bool) error {
	cx, cy := v.Cursor()
	ox, oy := v.Origin()
	nextLine, err := getNextViewLine(g, v)
	if err != nil {
		return err
	}
	if !allowEmpty && nextLine == " " {
		return nil
	}
	if err := v.SetCursor(cx, cy+1); err != nil {
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}
	return nil
}

func moveViewCursorUp(g *gocui.Gui, v *gocui.View, dY int) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if cy > dY {
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}
	}
	return nil
}

func getViewLine(g *gocui.Gui, v *gocui.View) (string, error) {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy); err != nil {
		line = ""
	}

	return line, err
}

func getNextViewLine(g *gocui.Gui, v *gocui.View) (string, error) {
	var line string
	var err error

	_, cy := v.Cursor()
	if line, err = v.Line(cy + 1); err != nil {
		line = ""
	}

	return line, err
}

func getQPlistNameFromLine(line string) string {
	if line == "" {
		return ""
	}

	i := strings.Index(line, " ")
	if i == -1 {
		return line
	}

	return line[0:i]
}

func getSelectedQPlist(g *gocui.Gui) (string, error) {
	v, err := g.View("qplist")
	if err != nil {
		return "", err
	}
	l, err := getViewLine(g, v)
	if err != nil {
		return "", err
	}
	p := getQPlistNameFromLine(l)

	return p, nil
}

func showViewCode(g *gocui.Gui) error {
	p, err := getSelectedQPlist(g)
	if err != nil {
		return err
	}
	codeV, err := g.View("code")
	if err != nil {
		return err
	}
	codeV.SetCursor(0, 0)
	SELECTED_POD_NUM, err = strconv.Atoi(p)
	fmt.Fprintf(codeV, "%s", CODE_STR[SELECTED_POD_NUM-1])

	vLc, err := g.View("content")
	if err != nil {
		return err
	}
	vLc.Clear()
	fmt.Fprintf(vLc, "%s", QP[SELECTED_POD_NUM-1].Content)
	vLc.SetCursor(0, 0)
	vLc1, err := g.View("example")
	if err != nil {
		return err
	}
	vLc1.Clear()
	strEx := strings.Split(QP[SELECTED_POD_NUM-1].Ex, " ")
	for i := 1; i < len(strEx)-1; i++ {
		n, _ := strconv.Atoi(strEx[i])
		fmt.Fprintf(vLc1, "%s", QP_EX[n].Ex+"\n\n")
	}

	g.SetViewOnTop("code")
	g.SetViewOnTop("content")
	g.SetViewOnTop("example")
	g.SetCurrentView("code")

	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	// g.SelFgColor = gocui.ColorRed

	g.SetManagerFunc(uiLayout)

	if err := uiKey(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func uiLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	loginAlert(g, maxX, maxY)
	viewTitle(g, maxX, maxY)
	viewStatusBar(g, maxX, maxY)
	viewCodeTool(g, maxX, maxY)
	viewCode(g, maxX, maxY)
	viewQPlist(g, maxX, maxY)
	viewEmpty(g, maxX, maxY)
	viewLogin(g, maxX, maxY)

	return nil
}
