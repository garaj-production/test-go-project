package main

import "github.com/garaj-production/private-go-dependency/pkg"

func main() {
	collection := map[string]int{
		"a": 1,
		"b": 2,
	}
	sl := make([]int, 0)
	for _, i2 := range collection {
		sl = append(sl, i2)
	}

	pkg.CustomPrint(collection)
}
