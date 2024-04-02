package foo

import (
	// unused dependency, just to populate the cache with _something_
	_ "golang.org/x/mod/modfile"
)

func Add(x, y int) int {
	return x + y
}
