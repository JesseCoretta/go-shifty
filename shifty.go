package shifty

import (
	"fmt"
	"strconv"
)

/*
Kind represents the specific unsigned integer type
selected for use within instances of BitValue.
*/
type Kind uint8

const (
	_      Kind = iota // 0x0
	Uint8              // 0x1
	Uint16             // 0x2
	Uint32             // 0x3
)

/*
String returns the string name of the receiver instance.
*/
func (r Kind) String() (k string) {
	k = `unknown`

	switch r {
	case Uint8:
		k = `uint8`
	case Uint16:
		k = `uint16`
	case Uint32:
		k = `uint32`
	}

	return
}

/*
Size returns the bit size of the receiver.
*/
func (r Kind) Size() (size int) {
	size = 0

	switch r {
	case Uint8:
		size = 8
	case Uint16:
		size = 16
	case Uint32:
		size = 32
	}

	return
}

/*
BitValue contains the allocated value type, a Kind and a bitsize per value type. Shift and Unshift
operations may be conducted against instances of this type. New instances of this type are created
using the New package-level function.
*/
type BitValue struct {
	k Kind  // avoid unnecessary type assertion during routine operations
	s uint8 // size (bits, 8 to 32), set manually so we don't need reflect for 'Size'
	v any   // ptr to uint8 or whatever is being used, per Kind (t)
}

/*
Value returns the unasserted instance of any from
within the receiver instance.
*/
func (r BitValue) Value() any {
	return r.v
}

/*
Int returns the integer form of the receiver value.
*/
func (r BitValue) Int() (i int) {
	switch r.k {
	case Uint8:
		i = int(*(r.v.(*uint8)))
	case Uint16:
		i = int(*(r.v.(*uint16)))
	case Uint32:
		i = int(*(r.v.(*uint32)))
	}

	return
}

/*
Kind returns the instance of Kind assigned to the receiver
instance.
*/
func (r BitValue) Kind() Kind {
	return r.k
}

/*
Shift shall left-shift the bits within the receiver to include
input value(s) x.
*/
func (r BitValue) Shift(x ...any) {
	for i := 0; i < len(x); i++ {
		if X, ok := r.verifyShiftValue(x[i]); ok {
			r.shift(X)
		}
	}
}

/*
Unshift shall right-shift the bits within the receiver to remove
input value(s) x.
*/
func (r BitValue) Unshift(x ...any) {
	for i := 0; i < len(x); i++ {
		if X, ok := r.verifyShiftValue(x[i]); ok {
			r.unshift(X)
		}
	}
}

/*
Positive returns a Boolean value indicative of whether input value
x's bits are set within the receiver.
*/
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

/*
Min returns eight (8), sixteen (16) or thirty-two (32), each
of which indicate the maximum bitsize of the receiver.
*/
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

/*
Size returns the underlying bitsize of the receiver.
*/
func (r BitValue) Size() (size int) {
	return r.k.Size()
}

/*
Min returns zero (0), which indicates the lowest permitted
value within the receiver.
*/
func (r BitValue) Min() (min int) {
	return 0
}

func (r BitValue) verifyShiftValue(x any) (X int, ok bool) {
	if X, ok = toInt(x); ok {
		ok = r.Min() <= X && X <= r.Max()
	}

	return
}

func New(k Kind) (bv BitValue) {
	bv.k = k

	switch k {
	case Uint8:
		bv.s = uint8(k.Size())
		bv.v = new(uint8)

	case Uint16:
		bv.s = uint8(k.Size())
		bv.v = new(uint16)

	case Uint32:
		bv.s = uint8(k.Size())
		bv.v = new(uint32)
	}

	return
}
