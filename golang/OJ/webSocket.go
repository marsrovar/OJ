package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/sacOO7/gowebsocket"
	"golang.org/x/crypto/bcrypt"
)

func IsSame(str string, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str)) == nil
}

func webSocketConn(g *gocui.Gui) {
	flag := ""

	SOCKET.OnConnectError = func(err error, socket gowebsocket.Socket) {
		// log.Fatal("Received connect error - ", err)
		log.Fatal("Server is not open")
	}

	SOCKET.OnConnected = func(socket gowebsocket.Socket) {
		// log.Println("Connected to server")
	}

	SOCKET.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		// log.Println("Received message - " + message)
		if flag == "login" {
			if IsSame("T", message) == true {
				checkLogin(g, "T")
				flag = ""
			} else {
				checkLogin(g, "ASD")
				flag = ""
			}
		} else if flag == "qp" {
			getQP("get", message)
			if getQPflag == 0 {
				QPdata(g)
				checkLogin(g, "qp")
				flag = ""
			}

		} else if flag == "code" {

		}

		switch message {
		case "login":
			flag = "login"
		case "qp":
			flag = "qp"
		case "code":
			flag = "code"
		}
	}

	SOCKET.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		// log.Println("Received ping - " + data)
	}

	SOCKET.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		// log.Println("Received pong - " + data)
	}

	SOCKET.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	SOCKET.Connect()

}
