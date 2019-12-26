package main

import (
	"fmt"
	"goDemo/Project/ETHWallet/balance"
	"time"
)

func main() {

	Delta := 100 * time.Millisecond

	for i := 0; i < 3; i++ {
		ticker := time.NewTicker(Delta)
		fmt.Println("1")
		<-ticker.C
		fmt.Println("2")
		ticker.Stop()
	}

	ticker2 := time.NewTicker(1 * time.Second)
	for range ticker2.C {
		fmt.Printf("trying count %d \n", balance.Count)
	}

	return

}
