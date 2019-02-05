package rtpengine

import (
	"fmt"
	"io"
)

// Engine handles a stream of incoming RTP/RTCP packets and does the following
//
// - reorders and removes duplicate RTP packets
// - responds/emits NACKs
// - emits sender/reciever reports
type Engine struct {
	nextConn io.ReadWriter
}

// NewEngine creates a new engine
func NewEngine(conn io.ReadWriter) (*Engine, error) {
	if conn == nil {
		return nil, fmt.Errorf("conn must not be nil")
	}

	return &Engine{nextConn: conn}, nil
}

// Read TODO
func (e *Engine) Read(b []byte) (int, error) {
	return 0, fmt.Errorf("TODO")
}

// Write TODO
func (e *Engine) Write(b []byte) (int, error) {
	return 0, fmt.Errorf("TODO")
}

// Stop tears down the Engine and clears any state
func (e *Engine) Stop() {
}
