# gowinkey

<div align="center">
  <a href="https://golang.org/">
    <img
      src="https://img.shields.io/badge/MADE%20WITH-GO-%23EF4041?style=for-the-badge"
      height="30"
    />
  </a>
  <a href="https://pkg.go.dev/github.com/daspoet/gowinkey">
    <img
      src="https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge"
      height="30"
    />
  </a>
  <a href="https://goreportcard.com/report/github.com/daspoet/gowinkey">
    <img
      src="https://goreportcard.com/badge/github.com/daspoet/gowinkey?style=for-the-badge"
      height="30"
    />
  </a>
</div>

## Contents

- [gowinkey](#gowinkey)
	- [Contents](#contents)
	- [Installation](#installation)
	- [If you prefer using Java...](#if-you-prefer-using-java)
	- [Getting started](#getting-started)
	- [Predicates](#predicates)
		- [Example](#example)

## Installation

To use `gowinkey`, you need to have [Go](https://golang.org) installed and set
up.

Now, you can get `gowinkey` by running

```shell
$ go get -u github.com/daspoet/gowinkey
```

and import it in your code:

```go
import "github.com/daspoet/gowinkey"
```

## If you prefer using Java...

... be sure to check out
the [Java version](https://github.com/lukasl-dev/jwinkey) of this package.

## Getting started

To start listening to key events, simply run `gowinkey.Listen()`. It returns a
channel on which all key events will be sent, and a function that you can call
whenever you want `gowinkey` to stop listening to events.

Note that calling this function will also close the channel.

Consider the following example:

```go
package main

import (
	"fmt"
	"github.com/daspoet/gowinkey"
	"time"
)

func main() {
	events, stopFn := gowinkey.Listen()

	time.AfterFunc(time.Minute, func() {
		stopFn()
	})

	for e := range events {
		switch e.State {
		case gowinkey.KeyDown:
			fmt.Println("pressed", e)
		case gowinkey.KeyUp:
			fmt.Println("released", e)
		}
	}
}
```

## Predicates

To help customise the way that `Listen` works, gowinkey uses [Predicates](https://github.com/DasPoet/gowinkey/blob/master/predicates.go). You can either supply your own, or use those that have already been created for you. You can find a collection of predefined predicates [here](https://github.com/DasPoet/gowinkey/blob/master/key_event.go).

### Example

Suppose we don't want to get bombarded with key events. For instance, we could only be interested in events for the keys *W*, *A*, *S* and *D*, because we want to write some basic movement for a game, or something.

We can now alter the code from above just slightly by passing the predefined predicate `Selective` to `Listen`. Notice that `Selective` takes the keys we want to listen for and returns an appropriate predicate that will handle all the work for us.

Consider the following code snippet:

```go
package main

import (
	"fmt"
	"github.com/daspoet/gowinkey"
	"time"
)

func main() {
	keys := []gowinkey.VirtualKey{
		gowinkey.VK_W,
		gowinkey.VK_A,
		gowinkey.VK_S,
		gowinkey.VK_D,
	}
	events, stopFn := gowinkey.Listen(gowinkey.Selective(keys...))

	time.AfterFunc(time.Minute, func() {
		stopFn()
	})

	for e := range events {
		switch e.State {
		case gowinkey.KeyDown:
			fmt.Println("pressed", e)
		case gowinkey.KeyUp:
			fmt.Println("released", e)
		}
	}
}
```
