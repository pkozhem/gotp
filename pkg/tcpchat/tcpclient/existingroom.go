package tcpclient

import (
	"gotp/pkg/tcpchat/domain"
	"sync"
)

func enterToExistingRoom(client *domain.TCPClient, rooms *sync.Map) {
	if client.Room != nil {
		client.WritelnString("You are already in the room")
		return
	}
	client.WriteString("Enter room ID for entering the room: ")
	input := client.ReadString()
	client.EnterRoomById(input, rooms)
	ChatToRoom(client)
}
