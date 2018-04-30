package main;

import (
	"github.com/pkg/errors"
	"fmt"
)

func main() {
	err := errors.New("hello my error")
	fmt.Println(err)
}
