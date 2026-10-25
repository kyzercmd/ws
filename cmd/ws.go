package main

import "github.com/gorilla/websocket"

type client struct {
	socket  *websocket.Conn
	receive chan []byte
	group   *group
}

type group struct {
	clients map[*client]bool
	join    chan *client
	leave   chan *client
	forward chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newGroup() *group {
	return &group{
		clients: make(map[*client]bool),
		join:    make(chan *client),
		leave:   make(chan *client),
		forward: make(chan []byte),
	}
}

func (g *group) run() {
	select {
	case client := <-g.join:
		g.clients[client] = true
	case client := <-g.leave:
		delete(g.clients, client)
		close(client.receive)
	case msg := <-g.forward:
		for client := range g.clients {
			client.receive <- msg
		}
	}
}

// client sends message into group's forward channel
func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.group.forward <- msg
	}
}

// client receives message from the receive channel over socket
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.receive {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
