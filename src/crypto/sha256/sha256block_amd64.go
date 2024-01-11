// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sha256

import "internal/cpu"

var useAVX2 = cpu.X86.HasAVX2 && cpu.X86.HasBMI2
var useSHA = useAVX2 && cpu.X86.HasSHA

func block(dig *digest, p []byte) {
	if useSHA || useAVX2 {
		block_std(dig, p)
	} else if cpu.X86.HasAVX {
		h := []uint32{dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7]}
		blockAVX(h[:], p[:], 0, 0, 0, 0)
		dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7] = h[0], h[1], h[2], h[3], h[4], h[5], h[6], h[7]
	} else if cpu.X86.HasSSSE3 {
		h := []uint32{dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7]}
		blockSSSE3(h[:], p[:], 0, 0, 0, 0)
		dig.h[0], dig.h[1], dig.h[2], dig.h[3], dig.h[4], dig.h[5], dig.h[6], dig.h[7] = h[0], h[1], h[2], h[3], h[4], h[5], h[6], h[7]	
	} else {
		block_std(dig, p)
	}
}
