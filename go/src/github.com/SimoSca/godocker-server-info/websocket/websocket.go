package websocket

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

type Message struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}

// Global variable to collect all clients and Broadcasts to them
var clients []Client

func onWsConnect(ws *websocket.Conn) {
	defer ws.Close()
	client := NewClient(ws)
	clients = addClientAndGreet(clients, client)
	client.listen()
}

func broadcast(msg *Message) {
	fmt.Printf("Broadcasting %+v\n", msg)
	for _, c := range clients {
		c.ch <- msg
	}
}

func addClientAndGreet(list []Client, client Client) []Client {
	clients = append(list, client)
	websocket.JSON.Send(client.connection, Message{"Server", "Welcome!"})
	return clients
}

func heartBeat() {
	for range time.Tick(time.Second * 10) {
		fmt.Println("Trigger HeartBeat")
		msg := &Message{
			Author: "Server",
			Body:   "Broadcast message",
		}
		broadcast(msg)
	}
}

func WsHandler(ws *websocket.Conn) {
	fmt.Println("Receive data from : ", ws.LocalAddr())
	fmt.Println("Sending data to : ", ws.RemoteAddr())

	//
	go heartBeat()

	onWsConnect(ws)

	// var err error

	// ws loop
	// for {
	// 	var reply string

	// 	if err = websocket.Message.Receive(ws, &reply); err != nil {
	// 		fmt.Println("Can't receive")
	// 		break
	// 	}

	// 	fmt.Println("Received back from client: " + reply)

	// 	msg := "Received:  " + reply
	// 	fmt.Println("Sending to client: " + msg)

	// 	if err = websocket.Message.Send(ws, msg); err != nil {
	// 		fmt.Println("Can't send")
	// 		break
	// 	}
	// }
}
