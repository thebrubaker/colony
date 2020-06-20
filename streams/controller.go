package streams

import (
	"log"

	"github.com/thebrubaker/colony/game"
	"github.com/thebrubaker/colony/keys"
	"github.com/thebrubaker/colony/pb"
)

type Streams []*Stream

// Controller manages all streams
type Controller struct {
	streams map[keys.GameKey]Streams
	actionc chan func()
	quitc   chan struct{}
}

// NewController returns a controller that mediates new stream
// connections to a specific GameKey and broadcasts game state
// across all registered streams for a given GameKey.
func NewController() *Controller {
	sc := &Controller{
		streams: make(map[keys.GameKey]Streams),
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go sc.loop()
	return sc
}

func (sc *Controller) loop() {
	for {
		select {
		case <-sc.quitc:
			return
		case f := <-sc.actionc:
			f()
		}
	}
}

// Stop ends all streams and closes the controller
func (sc *Controller) Stop() {
	for key, streams := range sc.streams {
		log.Printf("closing stream %s", key)
		for _, stream := range streams {
			stream.Stop()
		}
		delete(sc.streams, key)
	}
	close(sc.quitc)
}

// CreateStream will take a proto stream, transform it into
// a game stream, and register is with the controller. This allows
// the render loop to send a game's state to all registered streams
// for the given game's keys.GameKey.
func (sc *Controller) CreateStream(key keys.GameKey, gs pb.GameService_StreamGameServer) *Stream {
	c := make(chan *Stream)
	sc.actionc <- func() {
		log.Printf("registering stream for game %s", key)
		s := NewStream(gs)
		sc.streams[key] = append(sc.streams[key], s)
		c <- s
	}
	return <-c
}

// RemoveStream stops and deregisters the given stream on the given keys.GameKey
// so that renders no longer propagate to the stream.
func (sc *Controller) RemoveStream(key keys.GameKey, stream *Stream) {
	c := make(chan struct{})
	sc.actionc <- func() {
		log.Printf("removing stream for game %s", key)
		stream.Stop() // stop the stream

		streams, ok := sc.streams[key] // all streams for the given key
		if !ok {
			close(c)
			return
		}
		needle := -1 // searching for index of stream to remove
		for i, s := range streams {
			if s == stream {
				needle = i
			}
		}
		if needle == -1 { // not found
			close(c)
			return
		}
		// de-register stream
		streams[needle] = streams[len(streams)-1]
		streams[len(streams)-1] = stream
		sc.streams[key] = streams[:len(streams)-1]

		close(c)
	}
	<-c
}

// Broadcast sends the game state across all connected streams for the given GameKey.
func (sc *Controller) Broadcast(key keys.GameKey, g game.GameState) {
	c := make(chan struct{})
	sc.actionc <- func() {
		streams, ok := sc.streams[key]
		if !ok {
			close(c)
			return
		}
		for _, stream := range streams {
			stream.Send(string(key), g)
		}
		close(c)
	}
	<-c
}
