package main

import (
	"fmt"
	"github.com/yoshikouki/go-sample/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main()  {
	fmt.Println(morestrings.ReverseRunes("Hello, world."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
