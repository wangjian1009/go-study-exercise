package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type myServer struct{}

func (server myServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.String() {
	case "/":
		fmt.Fprintf(w, "get server")
	case "/test":
		fmt.Fprintf(w, "test")
	default:
		fmt.Fprintf(w, "unknow http")
	}

}

func main() {
	// 创建errgroup
	group, ctx := errgroup.WithContext(context.Background())

	// 创建http服务
	var s myServer
	se := http.Server{
		Handler: s,
		Addr:    ":9090",
	}
	http.Handle("/", s)

	// 启动http服务
	group.Go(func() error {
		defer fmt.Println("http server stoped")
		return se.ListenAndServe()
	})

	// 监听系统信号
	group.Go(func() error {
		signCh := make(chan os.Signal, 1)
		signal.Notify(signCh)

		var sig os.Signal
		var exitCancel context.CancelFunc

		for {
			select {
			case sig = <-signCh:
				switch sig {
				case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
					fmt.Println("received system fatal signal ", sig)

					if exitCancel == nil {
						// 没有在退出过程中, 则启动退出流程
						exitCtx, c := context.WithTimeout(context.Background(), 100*time.Millisecond)
						exitCancel = c

						group.Go(func() error {
							return se.Shutdown(exitCtx)
						})
					} else {
						exitCancel()
						return errors.New("find return signal,exit")
					}
				default:
					fmt.Println("received system signal ", sig, " ignore")
				}
			case <-ctx.Done():
				return nil
			}
		}
	})

	if err := group.Wait(); err != nil {
		fmt.Println("all goroutine are dead get errors:", err)
	}
}
