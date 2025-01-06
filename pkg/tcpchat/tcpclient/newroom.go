package tcpclient

import (
	"gotp/pkg/tcpchat/domain"
	"sync"
)

func createNewRoom(client *domain.TCPClient, rooms *sync.Map) {
	if client.Room != nil {
		client.WritelnString("You are already in the room")
		return
	}
	client.WriteString("Enter room name: ")
	input := client.ReadString()
	room := &domain.Room{Owner: client, Name: input}
	rooms.Store(client.Id, room)
	client.EnterRoom(room)
	ChatToRoom(client)
}
