package gowinkey

import (
	"fmt"
	"time"
)

func Example() {
	events, stopFn := Listen()

	time.AfterFunc(time.Minute, func() {
		stopFn()
	})

	for e := range events {
		switch e.State {
		case KeyDown:
			fmt.Println("pressed", e)
		case KeyUp:
			fmt.Println("released", e)
		}
	}
}
