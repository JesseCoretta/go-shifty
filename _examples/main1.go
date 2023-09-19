package main

import (
	"fmt"
	"github.com/JesseCoretta/go-shifty"
)

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
