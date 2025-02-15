package helpers

import (
	"fmt"
	"strings"
	"time"
)

func ShowLoadingWheel(done chan bool) {
	spinner := []string{
		"▪▫▫▫",
		"▫▪▫▫",
		"▫▫▪▫",
		"▫▫▫▪",
		"▫▫▫▫",
	}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r\r\n")
			return
		default:
			i = (i + 1) % len(spinner)
			fmt.Printf("\rGenerating response... %s", spinner[i])
			time.Sleep(170 * time.Millisecond)
		}
	}
}

func StreamResponse(input string) {
	for _, w := range strings.Fields(input) {
		fmt.Print(w + " ")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Print("\n")
}
