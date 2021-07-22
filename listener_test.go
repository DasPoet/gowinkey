package gowinkey

import (
	"fmt"
	"testing"
	"time"
)

func TestPredicates(t *testing.T) {
	events, stopFn := Listen(Hotkey(VK_1, VK_0))

	time.AfterFunc(time.Minute, func() {
		stopFn()
	})

	for e := range events {
		fmt.Println(e)
	}
}

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
