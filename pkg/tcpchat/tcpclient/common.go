package tcpclient

import "gotp/pkg/tcpchat/domain"

func ChatToRoom(client *domain.TCPClient){
	for {
		msg := client.ReadString()
		recievers := client.Room.Members
		for _, member := range recievers {
			member.WritelnString(msg)
		}
	}
}
