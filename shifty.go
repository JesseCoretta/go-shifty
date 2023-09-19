package shifty

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
		i = int(*(r.v.(*uint8)))
	case Uint16:
		i = int(*(r.v.(*uint16)))
	case Uint32:
		i = int(*(r.v.(*uint32)))
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
