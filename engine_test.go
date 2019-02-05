package rtpengine_test

import (
	"testing"

	"github.com/pions/rtpengine"
)

type testNextConn struct {
}

func (t *testNextConn) Read(b []byte) (int, error) {
	return 0, nil
}

func (t *testNextConn) Write(b []byte) (int, error) {
	return 0, nil
}

func TestEngineNew(t *testing.T) {
	if _, err := rtpengine.NewEngine(nil); err == nil {
		t.Fatal("nil conn was accepted by NewEngine")
	}

	engine, err := rtpengine.NewEngine(&testNextConn{})
	if err != nil {
		t.Fatal(err)
	} else if engine == nil {
		t.Fatal("NewEngine did not return an error, but the engine was nil")
	}
}
