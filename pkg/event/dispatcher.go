//go:generate mockgen -destination mock_resource/mock_dispatcher.go github.com/falcosecurity/cloud-native-security-hub/pkg/event Dispatcher

package event

import (
	"log"
)

type Dispatcher interface {
	Dispatch(event Interface)
}

type asyncDispatcher struct {
	eventChannel chan Interface
}

func (h asyncDispatcher) Dispatch(event Interface) {
	h.eventChannel <- event
}

func NewEventDispatcher() Dispatcher {
	handler := &asyncDispatcher{
		eventChannel: make(chan Interface),
	}
	go dispatchAsyncEvents(handler)
	return handler
}

func dispatchAsyncEvents(dispatcher *asyncDispatcher) {
	for event := range dispatcher.eventChannel {
		err := event.Handle()
		if err != nil {
			log.Println(err)
		}
	}
}
