package main

import (
	"fmt"
	"math"
)

func main() {
    var b byte = math.MaxUint8
    var smallI int32 = math.MaxInt32
    var bigI uint64 = math.MaxUint64

    fmt.Printf("Initial values:\n")
    fmt.Printf("b (byte):     %d\n", b)
    fmt.Printf("smallI (int32): %d\n", smallI)
    fmt.Printf("bigI (uint64):  %d\n", bigI)

    b++
    smallI++
    bigI++

    fmt.Printf("\nAfter adding 1:\n")
    fmt.Printf("b (byte):     %d\n", b)
    fmt.Printf("smallI (int32): %d\n", smallI)
    fmt.Printf("bigI (uint64):  %d\n", bigI)
}