package ports

import (
	"context"
	"fmt"
	"sync"
)

type outPort struct {
	basePort
	// here we can add fields
}

func newOutPort(id int) *outPort {
	return &outPort{
		basePort: basePort{
			id:       id,
			portType: PortTypeOut,
			stopCh:   make(chan struct{}),
		},
	}
}

func (p *outPort) Start(
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

func (p *outPort) Write(
	value int,
	transactionID int,
) error {
	fmt.Printf(
		"[OUT #%d] Write: transaction=%d, value=%d\n",
		p.ID(),
		transactionID,
		value,
	)
	return nil
}
