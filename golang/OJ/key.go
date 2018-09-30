package main

import "github.com/jroimartin/gocui"

var keys []Key = []Key{
	Key{"", gocui.KeyCtrlC, actionGlobalQuit},
	// Key{"", gocui.KeyCtrlA, actionViewPodsLogsDownA},
	Key{"loginEmail", gocui.KeyTab, actionViewLoginTab},
	Key{"loginEmail", gocui.KeyEnter, actionViewLoginEn},
	Key{"loginPassword", gocui.KeyTab, actionViewLoginTab},
	Key{"loginPassword", gocui.KeyEnter, actionViewLoginEn},
	Key{"loginBtn", gocui.KeyTab, actionViewLoginTab},
	Key{"loginBtn", gocui.KeyEnter, actionViewLoginBtn},
	Key{"qplist", gocui.KeyArrowUp, actionViewQPlistUp},
	Key{"qplist", gocui.KeyArrowDown, actionViewQPlistDown},
	Key{"qplist", gocui.KeyEnter, actionViewCode},
	Key{"code", gocui.KeyCtrlN, actionCodeTool},
	Key{"code", gocui.KeyCtrlL, actionViewCodeHide},
	Key{"code", gocui.KeyArrowUp, actionViewCodeUp},
	Key{"code", gocui.KeyArrowDown, actionViewCodeDown},
	Key{"code", gocui.KeyTab, actionViewCodeTab},
	Key{"code", gocui.KeyCtrlO, actionViewCodeNewLine},
	// Key{"content", gocui.KeyArrowDown, actionViewCodeDown},
	Key{"codetool", gocui.KeyArrowUp, actionViewCodeToolUp},
	Key{"codetool", gocui.KeyArrowDown, actionViewCodeToolDown},
	Key{"codetool", gocui.KeyEnter, actionViewCodeToolSelect},
	Key{"codetool", gocui.KeyCtrlN, actionCodeTool},
}

type Key struct {
	viewname string
	key      interface{}
	handler  func(*gocui.Gui, *gocui.View) error
}

func uiKey(g *gocui.Gui) error {
	for _, key := range keys {
		if err := g.SetKeybinding(key.viewname, key.key, gocui.ModNone, key.handler); err != nil {
			return err
		}
	}
	return nil
}
