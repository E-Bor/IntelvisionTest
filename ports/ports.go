package ports

import (
	"context"
	"sync"
)

type PortType int

const (
	PortTypeIn = iota
	PortTypeOut
)

type Port interface {
	Start(
		ctx context.Context,
		wg *sync.WaitGroup,
	)
	Stop()
	Type() PortType
	ID() int
}

type PortReader interface {
	Read() (
		int,
		error,
	)
}

type PortWriter interface {
	Write(
		value,
		transactionId int,
	) error
}

type basePort struct {
	id       int
	portType PortType
	stopCh   chan struct{}
}

func (p *basePort) ID() int { // cos we wanna using type embedding
	return p.id
}

func (p *basePort) Type() PortType {
	return p.portType
}

func (p *basePort) Stop() {
	close(p.stopCh)
}

func NewPort(
	pt PortType,
	id int,
) Port {
	switch pt {
	case PortTypeIn:
		return newInPort(id)
	case PortTypeOut:
		return newOutPort(id)
	default:
		panic("Unknown port type") // I don`t now should we skip unknown port type or make failure
	}
}

func AsReader(p Port) (
	PortReader,
	bool,
) {
	r, ok := p.(PortReader)
	return r, ok
}

func AsWriter(p Port) (
	PortWriter,
	bool,
) {
	w, ok := p.(PortWriter)
	return w, ok
}
