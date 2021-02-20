package debug

import (
	"encoding/json"

	tm "github.com/buger/goterm"
	"github.com/thebrubaker/colony/game"
)

type Stream struct {
	actionc chan func()
	quitc   chan struct{}
}

func NewStream() *Stream {
	s := &Stream{
		actionc: make(chan func()),
		quitc:   make(chan struct{}),
	}
	go s.loop()
	return s
}

func (s *Stream) loop() {
	for {
		select {
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
		data, err := json.MarshalIndent(g, "", "    ")

		if err != nil {
			c <- false
			return
		}

		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Println(string(data))
		tm.Flush()

		c <- true
	}
	return <-c
}
