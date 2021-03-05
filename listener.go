package gowinkey

import "sync"

const keyDown = 0x8000

// listener listens for global key events.
type listener struct {
	keyStates map[VirtualKey]bool
	mu        *sync.Mutex
}

var defaultListener = &listener{
	keyStates: make(map[VirtualKey]bool),
	mu:        new(sync.Mutex),
}

// Listen listens for global key events, sending them on the
// events channel.
// Listen halts execution and closes the events channel as
// soon as stopFn is called.
// Listen does not block.
func Listen() (events <-chan KeyEvent, stopFn func()) {
	return defaultListener.listen()
}

// listen listens for global key events, sending them on the
// returned channel. listen halts execution and closes the
// returned channel as soon as the returned function is called.
// listen does not block.
func (l *listener) listen() (<-chan KeyEvent, func()) {
	events, stopChan := make(chan KeyEvent), make(chan bool)

	go func() {
		l.swallowQueuedStates()
		l.doListen(events, stopChan)
	}()

	return events, func() { stopChan <- true; close(events) }
}

// doListen listens for global key events,
// sending them on the events channel.
func (l *listener) doListen(events chan KeyEvent, stopChan <-chan bool) {
Outer:
	for {
		select {
		case <-stopChan:
			break Outer
		default:
			l.listenOnce(events)
		}
	}
}

// listenOnce listens for the state of each of the 254 known
// virtual keys and sends according key events on the events
// channel.
func (l *listener) listenOnce(events chan KeyEvent) {
	for i := 0; i < 255; i++ {
		state := getKeyState(i)

		vk := VirtualKey(i)

		if vk == VK_SHIFT || vk == VK_CONTROL || vk == VK_MENU {
			continue
		}

		event := KeyEvent{VirtualKey: vk}

		if isKeyDown(state) {
			if !l.isPressed(vk) {
				event.Type = KeyPressed
				l.applyModifiers(&event)
				l.setIsPressed(vk, true)
				events <- event
			}
		} else {
			if l.isPressed(vk) {
				event.Type = KeyReleased
				l.setIsPressed(vk, false)
				events <- event
			}
		}
	}
}

// swallowQueuedStates drains the message queue so that the
// listener does not catch any events that were issue before
// listener.listen was called.
func (l listener) swallowQueuedStates() {
	for i := 0; i < 256; i++ {
		getKeyState(i)
	}
}

// applyModifiers applies modifiers for
// the currently pressed keys to the event.
func (l listener) applyModifiers(event *KeyEvent) {
	switch {
	case l.keyStates[VK_LSHIFT], l.keyStates[VK_RSHIFT]:
		event.Modifiers |= ModifierShift
	case l.keyStates[VK_LCONTROL], l.keyStates[VK_RCONTROL]:
		event.Modifiers |= ModifierControl
	case l.keyStates[VK_LMENU], l.keyStates[VK_RMENU]:
		event.Modifiers |= ModifierMenu
	}
}

// isPressed reports whether the given virtual key is
// registered as pressed in the listener.keyStates.
func (l listener) isPressed(key VirtualKey) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.keyStates[key]
}

// setIsPressed sets the status of the given virtual
// key in listener.keyStates to the given value.
func (l *listener) setIsPressed(key VirtualKey, value bool) {
	l.mu.Lock()
	l.keyStates[key] = value
	l.mu.Unlock()
}

func isKeyDown(keyState int32) bool {
	return keyState&keyDown == keyDown
}
