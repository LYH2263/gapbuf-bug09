package gapbuf

import "github.com/LYH2263/go-gapbuf/internal/audit"

func (e *Engine) Audit() *audit.Logger {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.audit
}
