package main

import (
	"fmt"
	"runtime"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("")
	fmt.Printf("os   %s\n", runtime.GOOS)
	fmt.Printf("arch %s\n", runtime.GOARCH)
	fmt.Printf("uuid %s\n", uuid.New().String())
	fmt.Println("")
}
