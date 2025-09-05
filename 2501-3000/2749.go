package main

import (
    "math/bits"
)

func makeTheIntegerZero(num1 int, num2 int) int {
    for t := 0; t <= 60; t++ {
        s := int64(num1) - int64(t)*int64(num2)
        if s < 0 { continue }
        if s < int64(t) { continue }
        ones := bits.OnesCount64(uint64(s))
        if ones <= t { return t }
    }
    return -1
}
