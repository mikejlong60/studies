package main

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/tree"
)

//Same determines whether the trees
//t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c := make(chan []int, 2)
	go FoldAsync(t1, c)
	go FoldAsync(t2, c)

	a1, a2 := <-c, <-c
	fmt.Println(a1)
	fmt.Println(a2)
	return reflect.DeepEqual(a1, a2)
}

func main() {

	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t2, t3))
}

func FoldAsync(T *tree.Tree, c chan []int) {
	c <- Fold(T)
}

func Fold(T *tree.Tree) []int {
	switch {
	case T.Left != nil && T.Right != nil:
		return append(append(Fold(T.Left), T.Value), Fold(T.Right)...)
	case T.Left != nil:
		return append(Fold(T.Left), T.Value)
	case T.Right != nil:
		return append([]int{T.Value}, Fold(T.Right)...)
	default:
		return []int{T.Value}
	}
}
