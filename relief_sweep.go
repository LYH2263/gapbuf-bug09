package gapbuf

import "context"

func (e *Engine) ReliefSweep(ctx context.Context) (int, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closed {
		return 0, ErrClosed
	}
	n := 0
	for _, id := range e.order {
		if err := ctx.Err(); err != nil {
			return n, err
		}
		if id != "" {
			n++
		}
	}
	return n, nil
}
