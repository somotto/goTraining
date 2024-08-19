package main

import "fmt"

const Value = 50

func main() {
    i := Value
    f := float64(Value)
    
    fmt.Printf("i: %d\n", i)
    fmt.Printf("f: %f\n", f)
}