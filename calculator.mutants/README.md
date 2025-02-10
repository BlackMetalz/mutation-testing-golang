## **Why did the mutants survive?**
### **1. LIVED ARITHMETIC_BASE at `calculator.go:9:14` (Add)**
   - **Possible mutation**: `return a + b` → `return a - b`
   - **Why it survived?** The test only checks `Add(2, 3)`, which assumes normal addition. However, if the mutant changed addition to subtraction, a test like `Add(3, 3) == 0` wouldn't be detected because there's no test for negative numbers or different values.
   
### **2. LIVED ARITHMETIC_BASE at `calculator.go:14:14` (Multiply)**
   - **Possible mutation**: `return a * b` → `return a / b`
   - **Why it survived?** The test only checks `Multiply(4, 3)`, assuming normal multiplication. If a mutant changed it to division, it might still return an integer that does not fail the test in some cases (e.g., `Multiply(2, 2) → 2 / 2 = 1`).
   
### **3. LIVED ARITHMETIC_BASE at `calculator.go:23:14` (Divide)**
   - **Possible mutation**: `return a / b` → `return a * b`
   - **Why it survived?** Your test suite only checks `Divide(6, 2)`, which might not cover mutations like replacing division with multiplication.
   
### **4. LIVED CONDITIONALS_NEGATION at `calculator.go:20:10` (Zero division case)**
   - **Possible mutation**: `if b == 0` → `if b != 0`
   - **Why it survived?** Your tests do **not check what happens when `b == 0`**. If the mutant inverted the condition, it would **always** return `0` for nonzero divisors, but your test would never detect this issue because it only tests division with nonzero values.


