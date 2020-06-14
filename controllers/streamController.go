package controllers

import (
	"log"

	"github.com/thebrubaker/colony/game"
	"github.com/thebrubaker/colony/pb"
	"github.com/thebrubaker/colony/streams"
)

type Streams []*streams.Stream

type StreamController struct {
	streams map[game.GameKey]Streams
	actionc chan func()
	quitc   chan struct{}
}

// NewStreamController returns a controller that mediates new stream
// connections to a specific GameKey and broadcasts game state
// across all registered streams for a given GameKey.
func NewStreamController() *StreamController {
	sc := &StreamController{
		streams: make(map[game.GameKey]Streams),
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go sc.loop()
	return sc
}

func (sc *StreamController) loop() {
	for {
		select {
		case <-sc.quitc:
			return
		case f := <-sc.actionc:
			f()
		}
	}
}

func (sc *StreamController) Stop() {
	for key, streams := range sc.streams {
		log.Printf("closing stream %s", key)
		for _, stream := range streams {
			stream.Stop()
		}
		delete(sc.streams, key)
	}
	close(sc.quitc)
}

// Create stream will take a proto stream, transform it into
// a game stream, and register is with the controller. This allows
// the render loop to send a game's state to all registered streams
// for the given game's game.GameKey.
func (sc *StreamController) CreateStream(key game.GameKey, gs pb.GameService_StreamGameServer) *streams.Stream {
	c := make(chan *streams.Stream)
	sc.actionc <- func() {
		log.Printf("registering stream for game %s", key)
		s := streams.NewStream(gs)
		sc.streams[key] = append(sc.streams[key], s)
		c <- s
	}
	return <-c
}

// RemoveStream stops and deregisters the given stream on the given game.GameKey
// so that renders no longer propagate to the stream.
func (sc *StreamController) RemoveStream(key game.GameKey, s *streams.Stream) bool {
	removed := make(chan bool)
	sc.actionc <- func() {
		log.Printf("removing stream for game %s", key)
		s.Stop() // stop the stream

		streams, ok := sc.streams[key] // all streams for the given key
		if !ok {
			removed <- false
			return
		}
		needle := -1 // searching for index of stream to remove
		for i, c := range streams {
			if c == s {
				needle = i
			}
		}
		if needle == -1 { // not found
			removed <- false
			return
		}
		// de-register stream
		streams[needle] = streams[len(streams)-1]
		streams[len(streams)-1] = nil
		sc.streams[key] = streams[:len(streams)-1]

		removed <- true
	}
	return <-removed
}

func (sc *StreamController) Broadcast(key game.GameKey, g game.GameState) bool {
	c := make(chan bool)
	sc.actionc <- func() {
		streams, ok := sc.streams[key]
		if !ok {
			c <- false
			return
		}
		for _, stream := range streams {
			stream.Send(string(key), g)
		}
		c <- true
	}
	return <-c
}
