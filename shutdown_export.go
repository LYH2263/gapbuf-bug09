package gapbuf

import (
	"encoding/json"
	"os"

	"github.com/LYH2263/go-gapbuf/internal/clone"
)

func (e *Engine) ShutdownExport(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	rows := make([]*Record, 0, len(e.order))
	for _, id := range e.order {
		r := e.items[id]
		rows = append(rows, &Record{
			ID: r.ID, Payload: clone.Bytes(r.Payload), Meta: clone.StringMap(r.Meta), At: r.At,
		})
	}
	b, err := json.Marshal(rows)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}
