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
  - [Filtering events](#filtering-events)

## Installation

To use `gowinkey`, you need to have [Go](https://golang.org) installed and set up.

Now, you can get `gowinkey` by running

```shell
$ go get -u github.com/daspoet/gowinkey
```

and import it in your code:

```go
import "github.com/daspoet/gowinkey"
```

## If you prefer using Java...


... be sure to check out the [Java version](https://github.com/lukasl-dev/jwinkey) of this package.


## Getting started

To start listening to key events, simply run `gowinkey.Listen()`. It returns a channel on which all key events will be sent, and a function that you can call whenever you want `gowinkey` to stop listening to events. 

Note that calling this function will also close the channel.

Consider the following example:

```go
package main

import (
    "github.com/daspoet/gowinkey"
    "time"
)

func main() {
    events, stopFn := gowinkey.Listen()
    
    time.AfterFunc(time.Minute, func() {
        stopFn()
    })
    
    for e := range events {
        switch e.Type {
        case gowinkey.KeyPressed:
            fmt.Println("pressed", e)
        case gowinkey.KeyReleased:
            fmt.Println("released", e)
        }
    }
}
```

This will listen for and print out key events for one minute, providing additional information on whether an event was raised because of a key press or release.

## Filtering events

Suppose that we don't want to get bombarded with key events. For instance, we could only be interested in events for the keys *W*, *A*, *S* and *D*, because we want to write some basic movement for a game, or something.

Well, that is what `gowinkey.ListenSelective()` was made for.

We can change our previous example only slightly to discard all events we don't want:

```go
package main

import (
    "github.com/daspoet/gowinkey"
    "time"
)

func main() {
    events, stopFn := gowinkey.ListenSelective(gowinkey.VK_W, gowinkey.VK_A, gowinkey.VK_S, gowinkey.VK_D)
    
    time.AfterFunc(time.Minute, func() {
        stopFn()
    })
    
    for e := range events {
        switch e.Type {
        case gowinkey.KeyPressed:
            fmt.Println("pressed", e)
        case gowinkey.KeyReleased:
            fmt.Println("released", e)
        }
    }
}
```

Notice that we pass the virtual keycodes we want to listen for explicitly and `gowinkey` handles all the work for us, so that we can concentrate on the important parts of our application. How wonderful!
