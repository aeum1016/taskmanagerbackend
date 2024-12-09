package util

import (
	"context"
	"fmt"
	"time"
)

// Schedule calls function `f` with a period `p` offsetted by `o`.
func Schedule(ctx context.Context, p time.Duration, o time.Duration, f func(time.Time)) {
	// Position the first execution
	first := time.Now().Truncate(p).Add(o)
	if first.Before(time.Now()) {
		first = first.Add(p)
	}
	fmt.Println(first)
	firstC := time.After(time.Until(first))

	t := &time.Ticker{C: nil}

	go func() {
		for {
			select {
			case v := <-firstC:
				// The ticker has to be started before f as it can take some time to finish
				t = time.NewTicker(p)
				f(v)
			case v := <-t.C:
				f(v)
			case <-ctx.Done():
				t.Stop()
				return
			}
		}
	}()
}
