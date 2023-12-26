# go-plugin
go tools set

# begin

1. import the package first:

    import (
    	"github.com/gz4z2b/go-plugin/slice"
    )

2. run the command:

    go mod tidy

3. use the function just like:

    source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dst := []int{2, 4, 5, 8, 10}
	println(slice.Diff(source, dst))
    