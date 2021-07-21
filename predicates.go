package gowinkey

// Predicate represents a filter for key events.
type Predicate func(event KeyEvent) bool
