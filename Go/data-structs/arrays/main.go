package main

import "log"

func main() {
	log.Println("Hello, World!")
	var array1 [6]int
	array1[0] = 1
	array1[1] = 2

	array2 := [6]int{1, 2, 3, 4, 5, 6}
	slice1 := array2[1:3]
	a := [...]int{3, 4, 5, 6, 7, 8, 9, 10, 122}
	log.Println(array1, array2, slice1)
	for _, v := range a {
		log.Println(v)
	}
}
