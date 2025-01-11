package ports

import (
	"context"
	"math/rand"
	"sync"
)

type inPort struct {
	basePort
	// here we can add fields
}

func newInPort(id int) *inPort {
	return &inPort{
		basePort: basePort{
			id:       id,
			portType: PortTypeIn,
			stopCh:   make(chan struct{}),
		},
	}
}

func (p *inPort) Start(
	ctx context.Context,
	wg *sync.WaitGroup,
) {
	go func() {
		wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-p.stopCh:
				// can add graceful shutdown or some logic
				return
			}
		}
	}()
}
func (p *inPort) Read() (
	int,
	error,
) {
	val := rand.Intn(2)
	return val, nil
}
