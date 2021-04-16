// Simple Message Echo Server
// Author: Arthur Zhang, Ian
// Create Date: 20190101
// Modify Date: 20210415
// link: https://www.jianshu.com/p/f058fdbdea58

package ws

import (
	"encoding/json"

	"github.com/IanVzs/Snowflakes/pkgs/logging"
	"github.com/gorilla/websocket"
)

// EchoClientManager is a websocket manager
type EchoClientManager struct {
	EchoClients map[*EchoClient]bool
	Broadcast   chan []byte
	Register    chan *EchoClient
	Unregister  chan *EchoClient
}

// EchoClient is a websocket client
type EchoClient struct {
	ID     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message is an object for websocket message which is mapped to json type
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Start is to start a ws server
func (manager *EchoClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.EchoClients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.EchoClients[conn]; ok {
				close(conn.Send)
				delete(manager.EchoClients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.EchoClients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.EchoClients, conn)
				}
			}
		}
	}
}

// Send is to send ws message to ws client
func (manager *EchoClientManager) Send(message []byte, ignore *EchoClient) {
	for conn := range manager.EchoClients {
		if conn != ignore {
			logging.Debug(message)
			conn.Send <- message
		}
	}
}

func (c *EchoClient) Read() {
	defer func() {
		Manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			c.Socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.ID, Content: string(message)})
		Manager.Broadcast <- jsonMessage
	}
}

func (c *EchoClient) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			logging.Infof("Write to %s\t-\t%s", c.ID, string(message))
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// Manager define a echo server manager
var Manager = EchoClientManager{
	Broadcast:   make(chan []byte),
	Register:    make(chan *EchoClient),
	Unregister:  make(chan *EchoClient),
	EchoClients: make(map[*EchoClient]bool),
}
