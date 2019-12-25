package main

import (
	"fmt"
	"myProject/videoCollector/commom"
	"myProject/videoCollector/engine"
	"os"
	"os/signal"
	"syscall"
)

func main() {


	conf := commom.ReadConfig()
	fmt.Println(conf)
	exitChan := make(chan struct{})

	eng := engine.NewEngine(conf)

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		msg := <-sig

		fmt.Println("receive exit msg:",msg)
		eng.Stop()
		os.Exit(1)
	}()

	eng.Run()

	<-exitChan


}
