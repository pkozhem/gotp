package tcpclient

import (
	"fmt"
	"gotp/pkg/tcpchat/domain"
	"sync"
)

const (
	RoomClientChooseCreateNew string = "1"
	RoomClientChooseConnectExisting string = "2"
)

func HandleClient(client *domain.TCPClient, connMap *sync.Map, rooms *sync.Map) {
	defer func() {
		client.Conn.Close()
		connMap.Delete(client.Id)
	}()
	client.WritelnString("Welcome to the TCP chat!")
	client.WriteString("You can:\n  1) Create new room\n  2) Connect to the existing room\nYour choice: ")	
	for {
		input := client.ReadString()
		handleClientChooseRoom(input, client, rooms)
	}
}

func handleClientChooseRoom(input string, client *domain.TCPClient, rooms *sync.Map) {
	switch input {
	case RoomClientChooseCreateNew:
		createNewRoom(client, rooms)
	case RoomClientChooseConnectExisting:
		listAvailableRooms(client, rooms)
		enterToExistingRoom(client, rooms)
	}
}

func listAvailableRooms(client *domain.TCPClient, rooms *sync.Map) {
	client.WritelnString("List of rooms:")
	rooms.Range(func(key, value interface{}) bool {
		client.WritelnString(fmt.Sprintf("  ID: %v. Name: %v", key, value.(*domain.Room).Name))
		return true
	})
}
