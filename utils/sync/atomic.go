// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package synci

import (
	"math"
	"sync/atomic"
	"unsafe"
)

// AddFloat64 add delta to given address atomically
func AddFloat64(addr *float64, delta float64) (new float64) {
	var old float64
	for {
		old = math.Float64frombits(atomic.LoadUint64((*uint64)(unsafe.Pointer(addr))))
		if atomic.CompareAndSwapUint64((*uint64)(unsafe.Pointer(addr)),
			math.Float64bits(old), math.Float64bits(old+delta)) {
			break
		}
	}
	return
}

func AddFloat32(addr *float32, delta float32) (new float32) {
	var old float32
	for {
		old = math.Float32frombits(atomic.LoadUint32((*uint32)(unsafe.Pointer(addr))))
		if atomic.CompareAndSwapUint32((*uint32)(unsafe.Pointer(addr)),
			math.Float32bits(old), math.Float32bits(old+delta)) {
			break
		}
	}
	return
}
