package gapbuf
import "errors"
var (
        ErrClosed = errors.New("gapbuf: closed")
        ErrInvalid = errors.New("gapbuf: invalid")
        ErrNotFound = errors.New("gapbuf: not found")
        ErrConflict = errors.New("gapbuf: conflict")
)
