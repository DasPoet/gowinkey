package gowinkey

// An EventType represents the type of a KeyEvent.
type EventType int

const (
	KeyPressed EventType = iota + 1 // Start counting from 1 so that KeyPressed does not become the default value of EventType.
	KeyReleased
)

// KeyEvent represents a key event.
type KeyEvent struct {
	// Type represents the event's type.
	Type EventType `json:"type,omitempty"`

	// VirtualKey represents the event's virtual key.
	VirtualKey VirtualKey `json:"virtualKey,omitempty"`

	// Modifiers is the bitwise or of the modifiers
	// that were active when the event was dispatched.
	Modifiers uint `json:"modifiers,omitempty"`
}

// Pressed reports whether e represents a key press.
func (e KeyEvent) Pressed() bool {
	return e.Type == KeyPressed
}

// Released reports whether e represents a key release.
func (e KeyEvent) Released() bool {
	return e.Type == KeyReleased
}

// HasShift reports whether e contains any 'shift' modifier.
func (e KeyEvent) HasShift() bool {
	return e.Modifiers&ModifierShift > 0
}

// HasControl reports whether e contains any 'ctrl' modifier.
func (e KeyEvent) HasControl() bool {
	return e.Modifiers&ModifierControl > 0
}

// HasMenu reports whether e contains any 'alt' modifier.
func (e KeyEvent) HasMenu() bool {
	return e.Modifiers&ModifierMenu > 0
}

// String returns the string representation of e.
func (e KeyEvent) String() string {
	s := e.VirtualKey.String()
	if e.HasControl() {
		s += "+Ctrl"
	}
	if e.HasShift() {
		s += "+Shift"
	}
	if e.HasMenu() {
		s += "+Alt"
	}
	return s
}
