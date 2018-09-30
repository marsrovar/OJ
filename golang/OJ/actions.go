package main

import (
	"strconv"

	"github.com/jroimartin/gocui"
)

var CODETOOL_DISPLAYED bool = false
var LOGIN_NUM int = 1

func actionGlobalQuit(g *gocui.Gui, v *gocui.View) error {
	SOCKET.Close()
	return gocui.ErrQuit
}

func actionViewLoginTab(g *gocui.Gui, v *gocui.View) error {
	if LOGIN_NUM == 1 {
		g.Cursor = true
		g.SetCurrentView("loginPassword")
		LOGIN_NUM++
	} else if LOGIN_NUM == 2 {
		g.Cursor = false
		g.SelFgColor = gocui.ColorRed
		g.SetCurrentView("loginBtn")
		LOGIN_NUM++
	} else {
		g.Cursor = true
		g.SelFgColor = gocui.ColorGreen
		g.SetCurrentView("loginEmail")
		LOGIN_NUM = 1
	}
	return nil
}

func actionViewLoginBtn(g *gocui.Gui, v *gocui.View) error {
	email, err := g.View("loginEmail")
	if err != nil {
		return err
	}
	viewEmail := email.ViewBuffer()
	password, err := g.View("loginPassword")
	if err != nil {
		return err
	}
	viewPassword := password.ViewBuffer()
	webSocketConn(g)

	SOCKET.SendText("login")
	SOCKET.SendText(viewEmail)
	SOCKET.SendText(viewPassword)

	return nil
}

//不跳行
func actionViewLoginEn(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func actionViewQPlistUp(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorUp(g, v, 2)
	return nil
}

func actionViewQPlistDown(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorDown(g, v, false)
	return nil
}

func actionViewCode(g *gocui.Gui, v *gocui.View) error {
	err := showViewCode(g)
	g.Cursor = true
	changeStatusContext(g, "SL")
	return err
}

func actionViewCodeUp(g *gocui.Gui, v *gocui.View) error {
	vLc, err := g.View("code")
	if err != nil {
		return err
	}
	moveViewCursorUp(g, vLc, 0)
	return nil
}

func actionViewCodeDown(g *gocui.Gui, v *gocui.View) error {
	// vLc, err := g.View("content")
	vLc, err := g.View("code")
	if err != nil {
		return err
	}
	moveViewCursorDown(g, vLc, false)
	return nil
}

func actionViewCodeNewLine(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	l, _ := v.Line(cy)
	v.SetCursor(len(l), cy)
	v.EditNewLine()
	return nil
}

func actionViewCodeTab(g *gocui.Gui, v *gocui.View) error {
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')
	v.EditWrite(' ')

	return nil
}

// func actionViewPodsLogsDownA(g *gocui.Gui, v *gocui.View) error {
// 	g.SetCurrentView("content")
// 	return nil
// }

func actionViewCodeHide(g *gocui.Gui, v *gocui.View) error {
	g.SetViewOnBottom("code")
	g.SetViewOnBottom("content")
	g.SetViewOnBottom("example")
	g.SetCurrentView("qplist")
	p, err := getSelectedQPlist(g)
	if err != nil {
		return err
	}
	// codeV, err := g.View("code")
	// if err != nil {
	// 	return err
	// }
	num, err := strconv.Atoi(p)
	CODE_STR[num-1] = v.ViewBuffer()
	v.Clear()

	g.Cursor = false
	changeStatusContext(g, "D")

	return nil
}

func actionCodeTool(g *gocui.Gui, v *gocui.View) error {
	vn := "codetool"
	if !CODETOOL_DISPLAYED {
		v, _ := g.View(vn)
		v.SetCursor(0, 0)
		g.Cursor = false
		g.SetViewOnTop(vn)
		g.SetCurrentView(vn)
		changeStatusContext(g, "SE")
	} else {
		g.Cursor = true
		g.SetViewOnBottom(vn)
		g.SetCurrentView("code")
		changeStatusContext(g, "SL")
	}

	CODETOOL_DISPLAYED = !CODETOOL_DISPLAYED
	return nil
}

func actionViewCodeToolUp(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorUp(g, v, 0)
	return nil
}

func actionViewCodeToolDown(g *gocui.Gui, v *gocui.View) error {
	moveViewCursorDown(g, v, false)
	return nil
}

func actionViewCodeToolSelect(g *gocui.Gui, v *gocui.View) error {
	action, _ := getViewLine(g, v)
	codeAction(g, action)
	CODETOOL_DISPLAYED = !CODETOOL_DISPLAYED

	// changeStatusContext(g, "SL")
	// actionGlobalToggleViewNamespaces(g, v)
	return nil
}
