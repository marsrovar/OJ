package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
)

type Question struct {
	ID       string
	Title    string
	Content  string
	CodeFunc string
	Ex       string
}

type Example struct {
	Ex string
}

var QP []Question
var QP_EX []Example

// var codeStr map[string]string
var CODE_STR []string

// func getExam(flag string) {
// 	if flag == "send" {
// 		socket.SendText("exam")
// 	} else if flag == "get" {
// 		examNum = 2

// 		QP = append(QP, Question{"4", "Reverse Integer", "Given a 32-bit signed integer, reverse digits of an integer."})
// 		QP = append(QP, Question{"5", "Two Sum", "Given an array of integers, return indices of the two numbers such that they add up to a specific target.\nYou may assume that each input would have exactly one solution, and you may not use the same element twice."})
// 		// codeStr = make(map[string]string)
// 		// codeStr["4"] = ""
// 		// codeStr["5"] = ""
// 		codeStr = append(codeStr, "")
// 		codeStr = append(codeStr, "")
// 	}
// }

var getQPcount = 1
var getQPnum = 1
var getQPnumExCount = 1
var getQPflag = 1
var getQPid = ""
var getQPtitle = ""
var getQPcontent = ""
var getQPcode = ""
var getQPex = "Z "
var getQPexCount = 0

func getQP(flag string, str string) {
	if flag == "send" {
		SOCKET.SendText("qp")
		SOCKET.SendText("1")
	} else if flag == "get" {
		if getQPcount == 1 { //題數
			getQPflag, _ = strconv.Atoi(str)
			getQPnum, _ = strconv.Atoi(str)
			getQPcount++
		} else if getQPcount == 2 { //example數
			getQPnumExCount, _ = strconv.Atoi(str)
			getQPcount++
		} else if getQPcount == 3 { //id
			getQPid = str
			getQPcount++
		} else if getQPcount == 4 { //標題
			getQPtitle = str
			getQPcount++
		} else if getQPcount == 5 { //內容
			getQPcontent = str
			getQPcount++
		} else if getQPcount == 6 { //code function
			getQPcode = str
			getQPcount++
		} else if getQPcount == 7 { //example
			getQPex += strconv.Itoa(getQPexCount) + " "
			QP_EX = append(QP_EX, Example{str})
			getQPexCount++
			getQPnumExCount--
			if getQPnumExCount == 0 {
				QP = append(QP, Question{getQPid, getQPtitle, getQPcontent, getQPcode, getQPex})
				CODE_STR = append(CODE_STR, getQPcode)
				getQPflag--
				getQPex = "Z "
				getQPcount = 2
			}
		}
	}
}

func checkLogin(g *gocui.Gui, str string) error {
	g.Update(func(g *gocui.Gui) error {
		if str == "T" {
			getQP("send", "")
		} else if str == "qp" {
			g.SetViewOnTop("qplist")
			g.SetCurrentView("qplist")
		} else {
			v, err := g.View("alert")
			if err != nil {
				return err
			}
			fmt.Fprintln(v, str)
			g.SetViewOnTop("alert")
		}
		return nil
	})
	return nil
}

func QPdata(g *gocui.Gui) error {
	g.Update(func(g *gocui.Gui) error {
		lMaxX, _ := g.Size()
		v, err := g.View("qplist")
		if err != nil {
			return err
		}
		v.Clear()
		viewQPlistAddLine(v, lMaxX, "Num", "ID", "Title", "STATUS")
		fmt.Fprintln(v, strings.Repeat("─", lMaxX))
		for i := 0; i < getQPnum; i++ {
			viewQPlistAddLine(v, lMaxX, strconv.Itoa(i+1), QP[i].ID, QP[i].Title, "X")
		}
		return nil
	})
	return nil
}

func codeAction(g *gocui.Gui, action string) {
	g.Update(func(g *gocui.Gui) error {
		if action == "Run Code" {
			v, _ := g.View("code")
			SOCKET.SendText("code")
			SOCKET.SendText("run")
			SOCKET.SendText(QP[SELECTED_POD_NUM-1].ID)
			SOCKET.SendText(v.ViewBuffer())
			g.Cursor = true
			g.SetViewOnBottom("codetool")
			g.SetCurrentView("code")
			changeStatusContext(g, "SL")
		} else if action == "Submit  Solution" {
			g.Cursor = true
			g.SetViewOnBottom("codetool")
			g.SetCurrentView("code")
			changeStatusContext(g, "SL")
		} else if action == "Reset" {
			v, _ := g.View("code")
			v.Clear()
			v.SetCursor(0, 0)
			CODE_STR[SELECTED_POD_NUM-1] = QP[SELECTED_POD_NUM-1].CodeFunc
			fmt.Fprintf(v, "%s", CODE_STR[SELECTED_POD_NUM-1])
			g.Cursor = true
			g.SetViewOnBottom("codetool")
			g.SetCurrentView("code")
			changeStatusContext(g, "SL")
		} else {
			g.Cursor = true
			g.SetViewOnBottom("codetool")
			g.SetCurrentView("code")
			changeStatusContext(g, "SL")
		}

		return nil
	})
}
