package main

import (
	"errors"
	"fmt"

	"sentinelerror/helloworld"
)

func wrapNewPointerError() error {
	return fmt.Errorf("wrap err: %w", fmt.Errorf("i am error"))
}

func wrapConstantPointerError() error {
	return fmt.Errorf("wrap err: %w", sentinelErr)
}

var sentinelErr = fmt.Errorf("i am error")

func main() {
	fmt.Println(errors.Is(wrapNewPointerError(), fmt.Errorf("i am error"))) // false
	fmt.Println(errors.Is(wrapConstantPointerError(), sentinelErr))         //
	fmt.Println(errors.Is(fmt.Errorf("something err: %w", helloworld.ResourceErrUnknown()), helloworld.ResourceErrUnknown()))
	fmt.Println(fmt.Errorf("something err: %w", helloworld.ResourceErrUnknown()))
	// fmt.Println(errors.Is(fmt.Errorf("i am error"), fmt.Errorf("i am error")))   // false
	// fmt.Println(fmt.Errorf("i am error") == fmt.Errorf("i am error")) // false
}
