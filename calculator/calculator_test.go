// calculator/calculator_test.go
package calculator

import "testing"

func TestCalculator(t *testing.T) {
    c := &Calculator{}

    // Weak test for Add - won't catch mutations like a - b
    t.Run("Add basic", func(t *testing.T) {
        if c.Add(2, 2) == 4 {
            return // Only testing one case
        }
        t.Error("Add failed")
    })

    // Weak test for Multiply - won't catch mutations like a / b
    t.Run("Multiply basic", func(t *testing.T) {
        if c.Multiply(3, 3) == 9 {
            return // Only testing one case
        }
        t.Error("Multiply failed")
    })

    // Weak test for Divide - won't catch mutations in zero check
    t.Run("Divide basic", func(t *testing.T) {
        if c.Divide(8, 2) == 4 {
            return // Not testing divide by zero
        }
        t.Error("Divide failed")
    })
}