package common

import (
	"errors"
)

// ErrNotFound mean Failed to find object
var ErrNotFound = errors.New("ErrNotFound: failed to find the object")

// ErrConflict mean Object already exists
var ErrConflict = errors.New("Object already exists")

// ErrMultipleResults mean Multiple results where only one expected
var ErrMultipleResults = errors.New("Multiple results where only one expected")
