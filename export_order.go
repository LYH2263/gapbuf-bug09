package gapbuf

func (e *Engine) ExportOrder() []string {
	e.mu.Lock()
	defer e.mu.Unlock()
	out := make([]string, len(e.order))
	copy(out, e.order)
	return out
}
