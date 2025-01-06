package chat

import (
	"fmt"
	"gotp/pkg/tcpchat/tcpclient"
	"gotp/pkg/tcpchat/domain"
	"net"
	"sync"

	"github.com/google/uuid"
)

func StartTCPChat(protocol, host, port string) {
	listener := createListener(protocol, host, port)
	defer listener.Close()
	var connMap = &sync.Map{}
	var rooms = &sync.Map{}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		id := uuid.New().String()
		client := &domain.TCPClient{Id: id, Conn: conn}
		connMap.Store(client.Id, client)
		go tcpclient.HandleClient(client, connMap, rooms)
	}
}

func createListener(protocol, host, port string) net.Listener {
	listener, err := net.Listen(protocol, fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Starting TCP server on %s:%s\n", host, port)
	return listener
}
