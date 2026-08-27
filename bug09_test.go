package gapbuf_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/LYH2263/go-gapbuf"
)

func TestBug09_RotateClosesOldAudit(t *testing.T) {
	d := filepath.Join(os.TempDir(), "audit_rotate")
	_ = os.RemoveAll(d)
	p1 := filepath.Join(d, "a.log")
	p2 := filepath.Join(d, "b.log")
	e, err := gapbuf.New(gapbuf.Options{AuditPath: p1})
	if err != nil {
		t.Fatal(err)
	}
	defer e.Close()
	if e.Audit() == nil {
		t.Fatal("no audit")
	}
	if err := e.Audit().Rotate(p2); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(p1); err != nil {
		t.Fatalf("remove old: %v", err)
	}
}
