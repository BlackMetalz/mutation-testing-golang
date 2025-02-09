// calculator/calculator_test.go
package calculator

import "testing"

// func TestAdd(t *testing.T) {
//     tests := []struct {
//         name     string
//         a, b     int
//         expected int
//     }{
//         {"positive numbers", 2, 3, 5},
//         {"negative numbers", -2, -3, -5},
//         {"zero", 0, 5, 5},
//     }

//     c := &Calculator{}
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             result := c.Add(tt.a, tt.b)
//             if result != tt.expected {
//                 t.Errorf("Add(%d, %d) = %d; want %d", 
//                     tt.a, tt.b, result, tt.expected)
//             }
//         })
//     }
// }

// func TestMultiply(t *testing.T) {
//     tests := []struct {
//         name     string
//         a, b     int
//         expected int
//     }{
//         {"positive numbers", 2, 3, 6},
//         {"negative numbers", -2, 3, -6},
//         {"zero", 0, 5, 0},
//     }

//     c := &Calculator{}
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             result := c.Multiply(tt.a, tt.b)
//             if result != tt.expected {
//                 t.Errorf("Multiply(%d, %d) = %d; want %d", 
//                     tt.a, tt.b, result, tt.expected)
//             }
//         })
//     }
// }

// func TestDivide(t *testing.T) {
//     tests := []struct {
//         name     string
//         a, b     int
//         expected int
//     }{
//         {"positive division", 6, 2, 3},
//         {"negative division", -6, 2, -3},
//         {"division by zero", 5, 0, 0},
//     }

//     c := &Calculator{}
//     for _, tt := range tests {
//         t.Run(tt.name, func(t *testing.T) {
//             result := c.Divide(tt.a, tt.b)
//             if result != tt.expected {
//                 t.Errorf("Divide(%d, %d) = %d; want %d", 
//                     tt.a, tt.b, result, tt.expected)
//             }
//         })
//     }
// }

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