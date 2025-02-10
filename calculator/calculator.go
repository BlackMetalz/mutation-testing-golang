// calculator/calculator.go
package calculator

// Calculator performs basic arithmetic operations
type Calculator struct{}

// Add returns the sum of two numbers
func (c *Calculator) Add(a, b int) int {
    return a + b
}

// Multiply returns the product of two numbers
func (c *Calculator) Multiply(a, b int) int {
    return a * b
}

// Divide returns the quotient of two numbers
// Returns 0 if dividing by zero
func (c *Calculator) Divide(a, b int) int {
    if b == 0 {
        return 0
    }
    return a / b
}