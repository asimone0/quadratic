A go library for solving quadratic equations

Usage:
xpos, xneg := quadratic.Solve(a,b,c)

a,b,c correspond to the coefficients of the quadratic to be solved.

It return 2 complex128 values that represent the roots of the equation.
Use real(xpos) and imag(xpos) to access the real and imaginary components of the roots
