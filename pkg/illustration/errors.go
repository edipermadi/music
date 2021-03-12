package illustration

import "errors"

var (
	// ErrNoSuchKeyboardKey when note does not have corresponding keyboard key
	ErrNoSuchKeyboardKey = errors.New("no such keyboard key")
)
