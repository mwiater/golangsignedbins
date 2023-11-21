package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("%s Heartbeat...\n", t.Format(time.RFC3339))
		}
	}
}
