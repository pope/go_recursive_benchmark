package main

func BasicRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return BasicRecursion(n-1) + BasicRecursion(n-2)
}

func ExtraCaseRecursion(n int) int {
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	case 3:
		return 2
	default:
		return ExtraCaseRecursion(n-1) + ExtraCaseRecursion(n-2)
	}
}

func TCORecursion(n int) int {
	return tcoRecursion(n, 0, 1)
}

func tcoRecursion(n int, a int, b int) int {
	switch n {
	case 0:
		return a
	case 1:
		return b
	default:
		return tcoRecursion(n-1, b, a+b)
	}
}
