package token_test

import (
	"fmt"

	"github.com/geosoft1/token"
)

func ExampleGetToken() {
	fmt.Println(token.GetToken(6))
	//Output:d9c544
}
