package streams

import (
	"encoding/json"

	"github.com/thebrubaker/colony/game"
	"github.com/thebrubaker/colony/pb"
)

type Stream struct {
	gs      pb.GameService_StreamGameServer
	actionc chan func()
	quitc   chan struct{}
}

func NewStream(gs pb.GameService_StreamGameServer) *Stream {
	s := &Stream{
		gs:      gs,
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go s.loop()
	return s
}

func (s *Stream) loop() {
	for {
		select {
		case <-s.gs.Context().Done():
			return
		case <-s.quitc:
			return
		case f := <-s.actionc:
			f()
		}
	}
}

func (s *Stream) Stop() {
	close(s.quitc)
}

func (s *Stream) Send(key string, g game.GameState) bool {
	c := make(chan bool)
	s.actionc <- func() {
		data, err := json.Marshal(g)

		if err != nil {
			c <- false
			return
		}

		s.gs.Send(&pb.GameState{
			GameKey: key,
			Json:    string(data),
		})

		c <- true
	}
	return <-c
}
