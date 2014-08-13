package js

// This file contains a thin js runtime layer so ClojureScript itself can run with minimal modifications.

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}
