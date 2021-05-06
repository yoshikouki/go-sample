package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/yoshikouki/go-sample/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
