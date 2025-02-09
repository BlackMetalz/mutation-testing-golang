package calculator

type Calculator struct{}

func (c *Calculator) Add(a, b int) int {
    // Adding complexity to create mutation opportunities
    // if a < 0 && b < 0 {
    if a < 0 || b < 0 {
        return -(abs(a) + abs(b))
    }
    return a + b
}

func (c *Calculator) Multiply(a, b int) int {
    // Adding edge cases that tests might miss
    if a == 0 || b == 0 {
        return 0
    }
    if a < 0 && b < 0 {
        return abs(a) * abs(b)
    }
    return a * b
}

func (c *Calculator) Divide(a, b int) int {
    // More complex zero handling
    if b == 0 {
        if a < 0 {
            return -1
        }
        return 0
    }
    return a / b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}