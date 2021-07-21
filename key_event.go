package gowinkey

// KeyEvent represents a key event.
type KeyEvent struct {
	// State represents the state of the event's virtual
	// key at the time the event was dispatches.
	State KeyState `json:"type,omitempty"`

	// VirtualKey represents the event's virtual key.
	VirtualKey VirtualKey `json:"virtualKey,omitempty"`

	// Modifiers is the bitwise or of the modifiers
	// that were active when the event was dispatched.
	Modifiers Modifiers `json:"modifiersToStr,omitempty"`
}

// HasShift reports whether e contains any 'shift' modifier.
func (e KeyEvent) HasShift() bool {
	return e.Modifiers.HasModifiers(ModifierShift)
}

// HasControl reports whether e contains any 'ctrl' modifier.
func (e KeyEvent) HasControl() bool {
	return e.Modifiers.HasModifiers(ModifierShift)
}

// HasMenu reports whether e contains any 'alt' modifier.
func (e KeyEvent) HasMenu() bool {
	return e.Modifiers.HasModifiers(ModifierMenu)
}

// String returns the string representation of e.
func (e KeyEvent) String() string {
	s := e.VirtualKey.String()
	if mods := e.Modifiers.String(); mods != "" {
		return s + " + " + mods
	}
	return s
}
