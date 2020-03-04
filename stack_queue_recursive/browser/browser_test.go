package browser

import (
	"fmt"
	"testing"
)

func Test_Browser(t *testing.T){
	b := NewBrowser()

	b.View("a")
	b.View("b")
	b.View("c")

	fmt.Println(b.Next())
	fmt.Println(b.Prev())
	fmt.Println(b.Prev())
	fmt.Println(b.Prev())
	fmt.Println(b.Prev())
	fmt.Println(b.Next())

}