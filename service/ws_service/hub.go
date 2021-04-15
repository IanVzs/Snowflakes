// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ws

import "github.com/IanVzs/Snowflakes/pkg/logging"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
	hubID      string
}

func newHub(hubID string) *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		hubID:      hubID,
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				HManager.checker <- h.hubID
				logging.Debugf("client leave from: %s >", h.hubID)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

type HubManager struct {
	Hubs    map[string]*Hub
	checker chan string
	runner  chan string
}

var HManager = HubManager{
	Hubs:    make(map[string]*Hub),
	checker: make(chan string),
	runner:  make(chan string),
}

func NewHub(hubID string) *Hub {
	hub := newHub(hubID)
	HManager.Hubs[hubID] = hub
	HManager.runner <- hubID
	logging.Debug("NewHub hubID -> runner")
	return hub
}

// 关闭已无Clinet的Hub回收
func (hm *HubManager) run() {
	logging.Debug("HubManager running.")
	for {
		select {
		case hubID, ok := <-HManager.runner:
			if !ok {
				logging.Debug("HManager.runner: !ok")
			} else {
				logging.Debugf("HManager(%s).runner: -ok", hubID)
				go HManager.Hubs[hubID].run()
			}

		case hubID, ok := <-HManager.checker:
			logging.Infof("client leave hubID: %s <", hubID)
			if !ok {
				logging.Debug("HManager.checker: !ok")
			} else {
				if len(hm.Hubs[hubID].clients) == 0 {
					logging.Infof("hub close: %s", hubID)
					delete(hm.Hubs, hubID)
				}
			}
		}
	}
}
