package tcp

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
)

type EchoHandler struct {
	activeConn sync.Map

	closing atomic.Bool
}

func (e EchoHandler) Handle(ctx context.Context, conn net.Conn) {

	//TODO implement me
	panic("implement me")
}

func (e EchoHandler) Close() error {
	//TODO implement me
	panic("implement me")
}
