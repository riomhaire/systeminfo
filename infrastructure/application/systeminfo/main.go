package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/riomhaire/systeminfo/infrastructure/application/systeminfo/bootstrap"
)

func main() {
	// tracefile, err := os.Create("app.trace")
	// check(err)

	// pprof.StartCPUProfile(tracefile)
	//	trace.Start(tracefile)
	// Shutdown
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting Down")
		// pprof.StopCPUProfile()
		// //trace.Stop()
		// tracefile.Close()
		os.Exit(0)
	}()

	application := bootstrap.Application{}
	application.Initialize()
	application.Run()
}
