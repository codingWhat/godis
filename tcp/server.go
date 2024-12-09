package tcp

import (
	"context"
	"github.com/codingWhat/godis/interface/tcp"
	"github.com/codingWhat/godis/lib/logger"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Config struct {
	Address string
}

func ListenAndServerWithSignal(config *Config, handler tcp.Handler) error {

	listen, err := net.Listen("tcp", config.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen")
	closeCh := make(chan struct{})
	go func() {
		sigCh := make(chan os.Signal)
		signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

		for {
			select {
			case <-sigCh:
				close(closeCh)
				time.Sleep(2 * time.Second)
			}
		}
	}()

	return ListenAndServe(listen, handler, closeCh)
}

func ListenAndServe(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) error {
	ctx := context.Background()

	defer func() {
		listener.Close()
		handler.Close()
	}()

	var wait sync.WaitGroup
	for {
		select {
		case <-closeChan:
			break
		default:
		}

		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		wait.Add(1)
		go func() {
			defer func() {
				recover()
				wait.Done()
			}()
			handler.Handle(ctx, conn)
		}()
	}

	wait.Wait()
}
