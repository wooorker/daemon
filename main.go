package main

import (
	"daemon/jobs/emailSend"
	"time"
)

// * * * * *
// minute 0-59
// hour 0-23
// day 1-31
// month 1-12
// week 1-7

func main() {
	for {
		select {
		case <-time.After(1 * time.Second):
			go consume()
		}
	}
}

func consume() {
	// register email queue consume
	emailSend.Register()
}
