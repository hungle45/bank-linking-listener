package main

import (
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/internal"
	"demo/bank-linking-listener/pkg/thread"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	config.LoadEnv("./.env")

	httpServer := internal.InitializeHTTPServer("./config/app.yml")
	consumerJob := internal.InitializeConsumerJob("./config/app.yml")

	rg := thread.NewRoutineGroup()

	rg.Run(httpServer.Run)
	rg.Run(consumerJob.Run)

	<-sigchan

	rg.Run(httpServer.Shutdown)
	rg.Run(consumerJob.Shutdown)

	rg.Wait()
}
