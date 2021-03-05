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
		switch e.Type {
		case KeyPressed:
			fmt.Println("pressed", e)
		case KeyReleased:
			fmt.Println("released", e)
		}
	}
}
