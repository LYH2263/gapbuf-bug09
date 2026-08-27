package gapbuf

type reliefMark struct {
	id       string
	relieved bool
}

func (e *Engine) Relief(id string, token string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.closed {
		return ErrClosed
	}
	if _, ok := e.items[id]; !ok {
		return ErrNotFound
	}
	if token == "" {
		return ErrInvalid
	}
	e.relief[id] = true
	return nil
}

func (e *Engine) IsRelieved(id string) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.relief[id]
}
