//go:generate mockgen -destination mock_resource/mock_handler.go github.com/falcosecurity/cloud-native-security-hub/pkg/event Handler

package event

import (
	"log"
)

type Handler interface {
	Dispatch(event Interface)
}

type asyncHandler struct {
	eventChannel chan Interface
}

func (h asyncHandler) Dispatch(event Interface) {
	h.eventChannel <- event
}

func NewHandler() Handler {
	handler := &asyncHandler{
		eventChannel: make(chan Interface),
	}
	go handleAsyncEvents(handler)
	return handler
}

func handleAsyncEvents(handler *asyncHandler) {
	for event := range handler.eventChannel {
		err := event.Handle()
		if err != nil {
			log.Println(err)
		}
	}
}
