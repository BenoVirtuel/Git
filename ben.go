package main

import "fmt"

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func main() {

	x := 3
	y := 4
	z := 5
	h := 11
	i := 11

	max_xy := max(x, y) //appeler la fonction max(x,y)
	max_xz := max(x, z) // appeler la fonction max(x,y)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d)=%d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z))
	fmt.Printf("max(%d, %d) = %d\n", h, i, max(h, i))

}
