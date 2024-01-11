// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha256

//go:noescape
func block_std(dig *digest, p []byte)

//go:noescape
func blockAVX(h []uint32, message []uint8, reserved0, reserved1, reserved2, reserved3 uint64)

//go:noescape
func blockSSSE3(h []uint32, message []uint8, reserved0, reserved1, reserved2, reserved3 uint64)
