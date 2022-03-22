package cmd

import (
	"fmt"
	"os"
	"strings"
)

func handleSignals() {
	exit := func(success bool) {
		os.Exit(1)
	}
	for {
		select {
		case osSignal := <-globalOSSignalCh:
			fmt.Println("Exiting on signal: ", strings.ToUpper(osSignal.String()))
			exit(true)
		}
	}
}
