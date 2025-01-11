package server

import (
	"IntelvisionTest/ports"
	"context"
	"fmt"
	"sync"
)

type Server struct {
	inPorts  []ports.Port
	outPorts []ports.Port

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func NewServer(numIn, numOut int) *Server {
	ctx, cancel := context.WithCancel(context.Background())

	s := &Server{
		inPorts: make(
			[]ports.Port,
			numIn,
		),
		outPorts: make(
			[]ports.Port,
			numOut,
		),
		ctx:    ctx,
		cancel: cancel,
	}

	// Create n inputs
	for i := 0; i < numIn; i++ {
		s.inPorts[i] = ports.NewPort(
			ports.PortTypeIn,
			i,
		)
	}
	// Create n outs
	for i := 0; i < numOut; i++ {
		s.outPorts[i] = ports.NewPort(
			ports.PortTypeOut,
			i,
		)
	}

	return s
}

// Start all ports in own goroutines
func (s *Server) Start() {
	s.wg.Add(len(s.inPorts) + len(s.outPorts))
	for _, p := range s.inPorts {
		p.Start(
			s.ctx,
			&s.wg,
		)
	}
	for _, p := range s.outPorts {
		p.Start(
			s.ctx,
			&s.wg,
		)
	}
	s.wg.Wait()
}

// Stop all goroutines with ports
func (s *Server) Stop() {
	s.cancel()
	for _, p := range s.inPorts {
		p.Stop()
	}
	for _, p := range s.outPorts {
		p.Stop()
	}
}

func (s *Server) Read(portNumber int) (
	int,
	error,
) {
	if portNumber < 1 || portNumber > len(s.inPorts) {
		return 0, fmt.Errorf(
			"invalid input port number: %d",
			portNumber,
		)
	}
	reader, ok := ports.AsReader(s.inPorts[portNumber-1])
	if !ok {
		return 0, fmt.Errorf(
			"port #%d is not a reader",
			portNumber,
		)
	}
	val, err := reader.Read()
	return val, err
}

func (s *Server) Write(
	portNumber,
	transactionID,
	value int,
) error {
	if portNumber < 1 || portNumber > len(s.outPorts) {
		return fmt.Errorf(
			"invalid out port number: %d",
			portNumber,
		)
	}
	writer, ok := ports.AsWriter(s.outPorts[portNumber-1])
	if !ok {
		return fmt.Errorf(
			"port #%d is not a writer",
			portNumber,
		)
	}
	return writer.Write(
		value,
		transactionID,
	)
}
