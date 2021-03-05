package gowinkey

// These flags define which modifiers are being pressed alongside
// some virtual key. Bits are or'ed to build a modifier field. See
// KeyEvent.Modifiers for more info.
//
// Note that these flags do not differentiate between the
// 'left' and 'right' versions of keys. So, for instance,
// pressing either 'left shift' or 'right shift' will simply
// result in ModifierShift being used to build the modifier
// field.
const (
	// ModifierShift identifies any 'shift' modifier.
	ModifierShift = 1 << iota
	// ModifierMenu identifies any 'alt' modifier.
	ModifierMenu
	// ModifierControl identifies any 'ctrl' modifier.
	ModifierControl
)
