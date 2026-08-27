package gapbuf
import ("sort"; "sync")
type GapItem struct { Seq uint64; Data []byte }
type GapBuffer struct {
        mu sync.Mutex
        items []GapItem
        next uint64
}
func NewGapBuffer(start uint64) *GapBuffer { return &GapBuffer{next: start} }
func (g *GapBuffer) Insert(seq uint64, data []byte) {
        g.mu.Lock(); defer g.mu.Unlock()
        g.items = append(g.items, GapItem{Seq: seq, Data: append([]byte(nil), data...)})
        sort.Slice(g.items, func(i, j int) bool { return g.items[i].Seq < g.items[j].Seq })
}
func (g *GapBuffer) DrainOrdered() []GapItem {
        g.mu.Lock(); defer g.mu.Unlock()
        out := make([]GapItem, 0)
        for len(g.items) > 0 && g.items[0].Seq == g.next {
                it := g.items[0]
                g.items = g.items[1:]
                g.next++
                cp := it
                cp.Data = append([]byte(nil), it.Data...)
                out = append(out, cp)
        }
        return out
}
func (g *GapBuffer) NextExpected() uint64 { g.mu.Lock(); defer g.mu.Unlock(); return g.next }
