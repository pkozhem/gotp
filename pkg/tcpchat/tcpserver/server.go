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
	var rooms = &sync.Map{}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		client := &domain.TCPClient{Id: uuid.New().String(), Conn: conn}
		go tcpclient.HandleClient(client, rooms)
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
