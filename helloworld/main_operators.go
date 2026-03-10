package main

import "fmt"

func main() {
	// Base values
	a, b, c, d := 10, 3, 15, 7
	x, y := true, false
	p := 12 // For bitwise

	// Arithmetic
	fmt.Printf("Arithmetic: %d + %d = %d, %d - %d = %d, %d * %d = %d\n",
		a, b, a+b, a, b, a-b, a, b, a*b)

	// Division & Modulus
	fmt.Printf("Division: %d / %d = %d, %d %% %d = %d\n", a, b, a/b, a, b, a%b)

	// Unary +/-
	fmt.Printf("Unary: +%d = %d, -%d = %d\n", a, +a, a, -a)

	// Increment/Decrement (statements only)
	a++ // Now 11
	b-- // Now 2
	fmt.Printf("Inc/Dec: a++ → %d, b-- → %d\n", a, b)

	// Assignment (=)
	e := 5
	fmt.Printf("Assignment: e = 5 → %d\n", e)

	// Compound Assignment (+=, etc.)
	e += 2 // Now 7
	fmt.Printf("Compound: e += 2 → %d\n", e)

	// Comparison
	fmt.Printf("Comparison: %d == %d → %t, %d > %d → %t\n",
		c, d, c == d, c, d, c > d)

	// Logical
	fmt.Printf("Logical: %t && %t → %t, %t || %t → %t, !%t → %t\n",
		x, y, x && y, x, y, x || y, x, !x)

	// Bitwise
	fmt.Printf("Bitwise: %b & %b = %b (%d), %b | %b = %b (%d)\n",
		p, b, p&b, p&b, p, b, p|b, p|b)
	fmt.Printf("Shift: %b << 1 = %b (%d), %b >> 1 = %b (%d)\n",
		p, p<<1, p<<1, p, p>>1, p>>1)
	fmt.Printf("AND NOT: %b &^ %b = %b (%d)\n", p, 1, p&^1, p&^1)
}
