// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

type Kind uint8

const (
	_      Kind = iota // 0x0
	Uint8              // 0x1
	Uint16             // 0x2
	Uint32             // 0x3
)

type BitValue struct {
	k Kind  // avoid unnecessary type assertion during routine operations
	s uint8 // size (bits, 8 to 32), set manually so we don't need reflect for 'Size'
	v any   // ptr to uint8 or whatever is being used, per Kind (t)
}

func (r BitValue) Value() any {
	return r.v
}

func (r BitValue) Int() (i int) {
	switch r.k {
	case Uint8:
		X := r.v.(*uint8)
		i = int(*X)
	case Uint16:
		X := r.v.(*uint16)
		i = int(*X)
	case Uint32:
		X := r.v.(*uint32)
		i = int(*X)
	}

	return
}

func (r BitValue) Kind() Kind {
	return r.k
}

func (r BitValue) Shift(x ...any) {
	for i := 0; i < len(x); i++ {
		if X, ok := r.verifyShiftValue(x[i]); ok {
			r.shift(X)
		}
	}
}

func (r BitValue) Unshift(x ...any) {
	for i := 0; i < len(x); i++ {
		if X, ok := r.verifyShiftValue(x[i]); ok {
			r.unshift(X)
		}
	}
}

func (r BitValue) Positive(x any) (posi bool) {
	if X, ok := r.verifyShiftValue(x); ok {
		posi = r.positive(X)
	}

	return
}

func (r BitValue) shift(x int) {
	if !r.positive(x) {
		switch r.k {
		case Uint8:
			*(r.v.(*uint8)) |= uint8(x)
		case Uint16:
			*(r.v.(*uint16)) |= uint16(x)
		case Uint32:
			*(r.v.(*uint32)) |= uint32(x)
		}
	}

	return
}

func (r BitValue) unshift(x int) {
	if r.positive(x) {
		switch r.k {
		case Uint8:
			*(r.v.(*uint8)) = *(r.v.(*uint8)) &^ uint8(x)
		case Uint16:
			*(r.v.(*uint16)) = *(r.v.(*uint16)) &^ uint16(x)
		case Uint32:
			*(r.v.(*uint32)) = *(r.v.(*uint32)) &^ uint32(x)
		}
	}

	return
}

func (r BitValue) positive(x int) (posi bool) {
	switch r.k {
	case Uint8:
		posi = *(r.v.(*uint8))&uint8(x) > 0
	case Uint16:
		posi = *(r.v.(*uint16))&uint16(x) > 0
	case Uint32:
		posi = *(r.v.(*uint32))&uint32(x) > 0
	}

	return
}

func toInt(x any) (v int, ok bool) {
	switch tv := x.(type) {
	case int:
		v = tv
		ok = true
	case uint8:
		v = int(tv)
		ok = true
	case uint16:
		v = int(tv)
		ok = true
	case uint32:
		v = int(tv)
		ok = true
	default:
		if X, err := strconv.Atoi(fmt.Sprintf("%d", tv)); err == nil {
			v = X
			ok = true
		}
	}

	return
}

func (r BitValue) Max() (max int) {
	switch r.k {
	case Uint8:
		max = int(^uint8(0))
	case Uint16:
		max = int(^uint16(0))
	case Uint32:
		max = int(^uint32(0))
	}

	return
}

func (r BitValue) Min() (min int) {
	return 0
}

func (r BitValue) verifyShiftValue(x any) (X int, ok bool) {
	if X, ok = toInt(x); !ok {
		return
	}

	if r.Min() <= X && X <= r.Max() {
		ok = true
	}

	return
}

func New(k Kind) (bv BitValue) {
	bv.k = k

	switch k {
	case Uint8:
		bv.s = 8
		bv.v = new(uint8)

	case Uint16:
		bv.s = 16
		bv.v = new(uint16)

	case Uint32:
		bv.s = 32
		bv.v = new(uint32)
	}

	return
}

/*
// define your specific constant values
// in powers of two. Type should match
// whatever type you selected (uint16
// per main below).
type Option uint16

const (
	UserOption1  Option = 1 << iota //     1
	UserOption2                     //     2
	UserOption3                     //     4
	UserOption4                     //     8
	UserOption5                     //    16
	UserOption6                     //    32
	UserOption7                     //    64
	UserOption8                     //   128
	UserOption9                     //   256
	UserOption10                    //   512
	UserOption11                    //  1024
	UserOption12                    //  2048
	UserOption13                    //  4096
	UserOption14                    //  8192
	UserOption15                    // 16384
	UserOption16                    // 32768
)

func main() {
	bv := New(Uint16)
	fmt.Printf("%T (%d, %v)\n",
		bv, bv.Int(), bv.Kind())
	fmt.Printf("Min:%d,Max:%d\n", bv.Min(), bv.Max())
	bv.Shift(UserOption4) // 8
	fmt.Printf("%T (%d, %v)\n",
		bv, bv.Int(), bv.Kind())
	bv.Shift(UserOption1) // 1
	fmt.Printf("%T (%d, %v)\n",
		bv, bv.Int(), bv.Kind())
	fmt.Printf("%t\n", bv.Positive(UserOption1))
	fmt.Printf("%t\n", bv.Positive(UserOption4))
	bv.Unshift(UserOption4, UserOption1)
	fmt.Printf("%t\n", bv.Positive(UserOption4))
	fmt.Printf("%t\n", bv.Positive(UserOption1))
}
*/
