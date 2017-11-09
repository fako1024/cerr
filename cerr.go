////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 by Fabian Kohn
//
// This source code is licensed under the Apache License, Version 2.0, found in
// the LICENSE file in the root directory of this source tree.
////////////////////////////////////////////////////////////////////////////////

package cerr

// #include "cerr.h"
import "C"
import (
	"strings"
	"unsafe"
)

// Error represents an error encounteres in the C domain. It consists of a simple
// byte buffer that can be populated with characters inside C functions
type Error struct {
	buffer [C.CERR_BUFFER_SIZE]byte
}

// CPtr returns a C char pointer to be handed over to functions in the C domain
func (e *Error) CPtr() *C.char {
	return (*C.char)(unsafe.Pointer(&e.buffer[0]))
}

// IsError returns if an error was encountered or not (determined by the length
// of the NULL-terminated string inside the buffer)
func (e *Error) IsError() bool {
	return strings.TrimRight(string(e.buffer[:]), "\x00") != ""
}

// Error returns a Go string version of the NULL-terminated string inside the buffer
func (e *Error) Error() string {
	return strings.TrimRight(string(e.buffer[:]), "\x00")
}
