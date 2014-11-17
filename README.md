A golang library for solving quadratic equations

Usage:

```go
import "github.com/tonestuff/quadratic"

xpos, xneg := quadratic.Solve(a,b,c)
```

a,b,c correspond to the coefficients of the quadratic to be solved.

Solve() returns 2 complex128 values that represent the positive and negative roots of the equation.

Use real(ComplexType) and imag(ComplexType) to access the real and imaginary components of the roots
