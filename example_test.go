package token_test

import (
	"fmt"
)

func a() string {
	return "d9c544"
}
func ExampleGetToken() {
	//fmt.Println(token.GetToken(6))
	fmt.Println(a())
	//Output:d9c544
}
