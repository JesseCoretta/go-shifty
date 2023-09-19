package main

import (
	"fmt"
	"github.com/JesseCoretta/go-shifty"
)

// define your specific constant values in powers
// of two (2) using an enumerator (shown below).
// Type should match whatever type the user chose
// (uint16, for example)
//
// Here we define an Option type. This isn't a
// strict requirement, but it is nice to be able
// to extend new methods, if needed. Otherwise,
// one can skip defining a type, and change the
// Option type in the enumerator to uint16 itself.
type Option uint16

const (
	// These constants may represent whatever the user needs.
	// Ideally, they should be named as something specific and
	// not so generic.  For example, if one were setting up options
	// for a log-level related value, one might choose names such as
	// Trace, Debug, Warning, etc.
	//
	// When defining constants, it is very important you limit their
	// enumeration to numbers less than (and NOT equal) to the return
	// value of BitValue.Max(). Never list that number in an iota
	// enumerator like as shown below. The reason for this is because
	// all of the other numbers shall sum to that value when all value
	// are selected together, therefore defining one would result in an
	// overflow.
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

	// sixteen (16) total = uint16 compatible (at max cap).
	//
	// In addition to all possible permutations of the
	// above, the underlying (whole) value can also be
	// read in a literal context for states zero (0) and
	// max (^uintN(0)). In other words, one need not use
	// the Positive method if the value is literally 0 or
	// the given maximum. A literal 0 means nothing has
	// shifted. A literal max value means everything has
	// shifted. Use of Positive method is only really
	// needed when the value is neither of those extremes.
)

func main() {
        bv := shifty.New(shifty.Uint16)
        fmt.Printf("%T (%d, %v)\n", bv, bv.Int(), bv.Kind())
        fmt.Printf("Min:%d,Max:%d\n", bv.Min(), bv.Max())

        bv.Shift(UserOption4) // 8
        fmt.Printf("%T (%d, %v)\n", bv, bv.Int(), bv.Kind())

        bv.Shift(UserOption1) // 1
        fmt.Printf("%T (%d, %v)\n", bv, bv.Int(), bv.Kind())

        fmt.Printf("%t\n", bv.Positive(UserOption1))
        fmt.Printf("%t\n", bv.Positive(UserOption4))

        bv.Unshift(UserOption4, UserOption1)

        fmt.Printf("%t\n", bv.Positive(UserOption4))
        fmt.Printf("%t\n", bv.Positive(UserOption1))
}
