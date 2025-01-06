package domain

import (
	"bufio"
	"fmt"
	"net"
	"slices"
	"strings"
	"sync"
)

type IClient interface {
	WriteString(s string)
	ReadString() string
	EnterRoom(room *Room)
}

type TCPClient struct {
	Id string
	Room *Room
	Conn net.Conn
}

func (c *TCPClient) WriteString(s string) {
	c.Conn.Write([]byte(s))
}

func (c *TCPClient) WritelnString(s string) {
	s = fmt.Sprintf("%s\n", s)
	c.WriteString(s)
}

func (c *TCPClient) ReadString() string {
	input, err := bufio.NewReader(c.Conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	return input
}

func (c *TCPClient) EnterRoom(room *Room) {
	if c.Room == nil {
		c.Room = room
	}
	room.Members = append(room.Members, c)
	c.WritelnString(fmt.Sprintf("Connected to the room %s.\nRoom members:", room.Name))
	for _, member := range room.Members {
		c.WriteString(fmt.Sprintf("%s\n", member.Id))
	}
}

func (c *TCPClient) EnterRoomById(id string, rooms *sync.Map) {
	room, ok := rooms.Load(id)
	if !ok {
		c.WritelnString("No room with this ID")
		return
	}
	c.EnterRoom(room.(*Room))
}

func (c *TCPClient) ExitRoom() {
	if c.Room == nil {
		c.WritelnString("No need to exit a room - you didn't entered one")
	}
	idx := slices.IndexFunc(c.Room.Members, func(cl *TCPClient) bool {
		return cl == c
	})
	c.Room.Members = append(c.Room.Members[:idx], c.Room.Members[idx+1:]...)
}
