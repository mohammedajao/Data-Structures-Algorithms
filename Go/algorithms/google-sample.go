package main

func fibonacci_sequence(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		println(a)
		a, b = a+b, a
	}
	dict := make(map[string]int)
	dict["Age"] = 10
}

func main() {
	fibonacci_sequence(10)
}
