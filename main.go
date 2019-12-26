package main

import (
	"fmt"
	"myProject/videoCollector/common"
	"myProject/videoCollector/engine"
	"os"
	"os/signal"
	"syscall"
)

func main() {


	conf := common.ReadConfig()
	fmt.Println(conf)

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


}
