package tworoom

import (
	"github.com/LucaHhx/nano/component"
	"github.com/LucaHhx/nano/session"
)

var (
	// All services in master server
	Services = &component.Components{}

	roomService = newChatRoomService()
)

func init() {
	Services.Register(roomService)
}

func OnSessionClosed(s *session.Session) {
	roomService.userDisconnected(s)
}
