package handlers

import (
	"github.com/lonelyevil/kook"
)

func RegistryHandlers(s *kook.Session, handlers ...any) {
	for _, e := range handlers {
		s.AddHandler(e)
	}
}
