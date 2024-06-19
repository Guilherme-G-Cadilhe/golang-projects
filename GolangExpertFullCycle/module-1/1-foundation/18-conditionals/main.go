package main

func main() {
	a := 1
	b := 2
	c := 3

	// a >= b
	// a <= b
	// a == b
	// a != b
	// a < b
	if a > b {
		println(a)

	} else {
		println(b)
	}

	if a > b && c > a {
		println(a)
	}

	if (a < b) || (a == b) {
		println(a)
	}

	if !(a == b) {
		println(a)
	}

	switch a {
	case 1:
		println("Switch case 1:", 'a')
	case 2:
		println("Switch case 2:", "b")
	default:
		println("Switch case 3:", "c")
	}

}
