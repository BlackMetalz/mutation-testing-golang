package calculator

import "testing"

func TestCalculator(t *testing.T) {
    c := &Calculator{}

    // Minimal test that won't catch negative number mutations
    t.Run("Add", func(t *testing.T) {
        if got := c.Add(2, 3); got != 5 {
            t.Errorf("Add(2, 3) = %d; want 5", got)
        }
    })

    // Won't catch the negative number multiplication case
    t.Run("Multiply", func(t *testing.T) {
        if got := c.Multiply(4, 3); got != 12 {
            t.Errorf("Multiply(4, 3) = %d; want 12", got)
        }
    })

    // Won't catch the negative number division case
    t.Run("Divide", func(t *testing.T) {
        if got := c.Divide(6, 2); got != 3 {
            t.Errorf("Divide(6, 2) = %d; want 3", got)
        }
    })
}