package domain

import "sync"

type Room struct {
	Owner *TCPClient
	Members []*TCPClient
	Name string
}

func DestroyRoom(r *Room, rooms *sync.Map) {
	r.Owner = nil
	r.Name = ""
	for _, member := range r.Members {
		member.Room = nil
	}
	r.Members = nil
}
