// Copyright 2021 utils. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"testing"
)

func TestBitmap(t *testing.T) {
	bmp := DefaultBitMap
	for i := 0; i < 500; i++ {
		bmp.Set(1000000 + i)
	}
	fmt.Println(bmp.Include(63))
	fmt.Println(bmp.Include(67))
	fmt.Println(bmp.All())
	fmt.Println(len(bmp.bits))
}
