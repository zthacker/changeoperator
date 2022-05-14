package main

import (
	coAPI "coAPI/internal"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	go coAPI.Run()

	qChannel := make(chan os.Signal, 1)
	signal.Notify(qChannel, syscall.SIGINT, syscall.SIGTERM)
	<-qChannel
	fmt.Println("Hasta la vista, baby")
}
