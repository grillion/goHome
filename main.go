package main

import (

	"github.com/grillion/goHome/controllers/httpServices"
	"github.com/grillion/goHome/db"
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
)

func init() {

}

func main() {

	pid := syscall.Getpid()

	// Close DB when exiting
	defer db.CloseSession()

	// Channels for application tasks
	inform := make(chan int) // mPower inform event channel
	sigs  := make(chan os.Signal, 1) // OS Signal channel
	tick := time.Tick(1000 * time.Millisecond) // 1 second ticker for timed Application operations
	exitCode := -1

	// Start http services
	// +--- Static Files
	// +--- API End Points
	// +--- mPower inform handler
	go httpServices.Start()

	// Redirect OS Signals to channel
	signal.Notify(sigs)

	fmt.Println("Application ready (", pid, ")")

	// Handle operations until told to exit ( or we crash )
	for exitCode < 0 {
		select {
		case newSignal := <-sigs:

			switch newSignal {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("hungup")

			// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("interrupt")
				exitCode = 0

			// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				exitCode = 0

			// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				exitCode = 0

			default:
				fmt.Println("Unknown signal.")
				exitCode = 1
			}

		case <-inform:
			fmt.Println("inform")
		case <-tick:
			fmt.Print(".")
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

	fmt.Println("Cleaning up")

	os.Exit(exitCode)
}
