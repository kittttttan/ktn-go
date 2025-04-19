package main

import (
	"os"
	"strconv"

	"ktn/math/collatz"
)

func main() {
	var n uint64 = 10
	if len(os.Args) > 1 {
		t, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			println("parse error: ", os.Args[1])
			println("required uint64")
			return
		}

		n = t
	}

	println("Collatz series starts with:", n)

	cnt, max, err := collatz.Count(n)
	if err != nil {
		println("Error:", err)
		return
	}

	println("cnt =", cnt, " max =", max)

	iter := collatz.Iter(n)
	for v := range iter {
		println(v)
	}
}
