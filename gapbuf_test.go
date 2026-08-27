package gapbuf_test

import (
	"testing"

	"github.com/LYH2263/go-gapbuf"
)

func newEngine(t *testing.T) *gapbuf.Engine {
	e, err := gapbuf.New(gapbuf.Options{})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = e.Close() })
	return e
}
