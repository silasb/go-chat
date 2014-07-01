// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

var channels = make(map[string]*hub)

//type channels struct {
//	hubs map[*hub]bool
//	register chan
//}

// hub maintains the set of active connections and broadcasts messages to the
// connections.
type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

//var channels = channels{
//	register: make(string channel, chan *hub),
//	unregister: make(string channel, chan *hub),
//}

//var h = hub{
//	broadcast:   make(chan []byte),
//	register:    make(chan *connection),
//	unregister:  make(chan *connection),
//	connections: make(map[*connection]bool),
//}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			log.Println("client connect: ", c)

			h.connections[c] = true
		case c := <-h.unregister:
			log.Println("client disconnect: ", c)

			delete(h.connections, c)
			close(c.send)
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
		}
	}
}
