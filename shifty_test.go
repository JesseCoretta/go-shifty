package shifty

import (
	"fmt"
	"testing"
)

var testMap map[int]Kind

func ExampleNew() {
	bits := New(Uint16)
	fmt.Printf("%T size %d, max %d", bits, bits.Size(), bits.Max())
	// Output: shifty.BitValue size 16, max 65535
}

func ExampleBitValue_Shift() {
	bits := New(Uint8)
	bits.Shift(2, 4, 32)
	fmt.Printf("Value: %d", bits.Int())
	// Output: Value: 38
}

func ExampleBitValue_Unshift() {
	bits := New(Uint8)
	bits.Shift(2, 4, 32)
	bits.Unshift(32)
	fmt.Printf("Value: %d", bits.Int())
	// Output: Value: 6
}

func ExampleBitValue_Min() {
	var bits BitValue
	fmt.Printf("%d", bits.Min())
	// Output: 0
}

func ExampleBitValue_Max_for8Bit() {
	bits := New(Uint8)
	fmt.Printf("%d", bits.Max())
	// Output: 255
}

func ExampleBitValue_Max_for16Bit() {
	bits := New(Uint16)
	fmt.Printf("%d", bits.Max())
	// Output: 65535
}

func ExampleBitValue_Max_for32bit() {
	bits := New(Uint32)
	fmt.Printf("%d", bits.Max())
	// Output: 4294967295
}

func ExampleBitValue_Int() {
	bits := New(Uint8)
	bits.Shift(2, 4, 32)
	fmt.Printf("%d", bits.Int())
	// Output: 38
}

func ExampleBitValue_Int_mixed() {
	var ints []int
	for i := 0; i < 3; i++ {
		bits := New(Kind(i + 1))
		bits.Shift(bits.Size() << i)
		ints = append(ints, bits.Int())
	}
	fmt.Printf("%v", ints)
	// Output: [8 32 128]
}

func ExampleBitValue_Value() {
	bits := New(Uint32)
	bits.Shift(bits.Max())
	fmt.Printf("%T", bits.Value())
	// Output: *uint32
}

func ExampleBitValue_Kind() {
	bits := New(Uint32)
	fmt.Printf("%s", bits.Kind())
	// Output: uint32
}

func ExampleBitValue_Size() {
	bits := New(Uint32)
	fmt.Printf("%d", bits.Size())
	// Output: 32
}

func ExampleKind_Size() {
	k := Uint32
	fmt.Printf("%d", k.Size())
	// Output: 32
}

func ExampleKind_String() {
	fmt.Printf("%s", Uint32)
	// Output: uint32
}

func ExampleBitValue_Positive() {
	// user-defined shift values
	type B uint8
	const (
		Bopt1 B = 1 << iota //   1
		Bopt2               //   2
		Bopt3               //   4
		Bopt4               //   8
		Bopt5               //  16
		Bopt6               //  32
		Bopt7               //  64
		Bopt8               // 128	// go no higher (else, overflow uint8)
	)

	bits := New(Uint8)
	bits.Shift(Bopt1, Bopt3, Bopt6)
	fmt.Printf("Value contains B-options #6 (32): %t", bits.Positive(Bopt6))
	// Output: Value contains B-options #6 (32): true
}

func TestBitValue_codecov(t *testing.T) {
	var bits BitValue
	bits.Kind()
	bits.Int()
	bits.Value()
	bits.Shift(-1)
	bits.Shift(8 << 8)
	bits.Unshift(-1)
	bits.Positive(-1)
	bits.Unshift(40000000000)
	if i := bits.Int(); i != 0 {
		t.Errorf("%s failed: bogus value set (%d) where none should be",
			t.Name(), i)
	}

	bits = New(Uint8)
	bits.Shift(bits.Max())
	bits.Shift(8 << 8)
	bits.Shift(8 << 1)
	bits.Positive(8 << 2)
	bits.Unshift(8 << 8)
	bits.Kind()
	bits.Int()
	bits.Value()

	for _, kind := range testMap {
		instance := New(kind)
		size := instance.Size()
		_ = kind.String()
		_ = bits.Int()
		for i := 0; i < size; i++ {
			instance.Shift(size << i)
			instance.Positive(size << i)
			instance.Unshift(size << i)
			instance.Shift(instance.Max())
			instance.Unshift(instance.Max())
			switch instance.Value().(type) {
			case *uint8:
				_, _ = toInt(uint8(size))
			case *uint16:
				_, _ = toInt(uint16(size))
			case *uint32:
				_, _ = toInt(uint32(size))
			}
		}
	}
}

func init() {
	testMap = map[int]Kind{
		8:  Uint8,
		16: Uint16,
		32: Uint32,
	}
}
