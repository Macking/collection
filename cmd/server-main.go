package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	globalOSSignalCh = make(chan os.Signal, 1)
)

func serverMain() {
	signal.Notify(globalOSSignalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go handleSignals()

	handler, err := configureServerHandler()
	if err != nil {
		fmt.Println("Can't config server ", err)
	}

	server := &http.Server{
		Addr:    ":8089",
		Handler: handler,
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Can't start server ", err)
	}
	<-globalOSSignalCh

}
