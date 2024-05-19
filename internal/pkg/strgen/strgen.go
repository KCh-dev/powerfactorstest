package strgen

import (
	"goapp/pkg/util" // Ensure this is the correct path to your util package
	"sync"
	"time"
)

type StringGenerator struct {
	strChan     chan<- string  // String output channel.
	quitChannel chan struct{}  // Quit.
	running     sync.WaitGroup // Running.
}

func New(strChan chan<- string) *StringGenerator {
	s := StringGenerator{}
	s.strChan = strChan
	s.quitChannel = make(chan struct{})
	s.running = sync.WaitGroup{}
	return &s
}

// Start string generator. Stop() must be called at the end.
func (s *StringGenerator) Start() error {
	s.running.Add(1)
	go s.mainLoop()
	return nil
}

func (s *StringGenerator) Stop() {
	close(s.quitChannel)
	s.running.Wait()
}

func (s *StringGenerator) mainLoop() {
	defer s.running.Done()
	for {
		select {
		case s.strChan <- util.RandHexString(20): // 20 characters of hex (10 bytes)
		case <-s.quitChannel:
			return
		default:
		}
		time.Sleep(1 * time.Second)
	}
}
