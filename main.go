// main.go
package main

import (
    "fmt"
    "calculator-project/calculator"  // Use module path instead of relative path
)

func main() {
    calc := &calculator.Calculator{}

    // Example usage of calculator
    a, b := 10, 5

    sum := calc.Add(a, b)
    fmt.Printf("%d + %d = %d\n", a, b, sum)

    product := calc.Multiply(a, b)
    fmt.Printf("%d * %d = %d\n", a, b, product)

    quotient := calc.Divide(a, b)
    fmt.Printf("%d / %d = %d\n", a, b, quotient)

    // Example of division by zero handling
    result := calc.Divide(10, 0)
    fmt.Printf("10 / 0 = %d (division by zero returns 0)\n", result)
}