package main

import (
	"crypto/rand"
	"fmt"
	"os"

	"./evil"
)

func main() {
	fmt.Println(evil.GetContent())
	buf := make([]byte, 32)
	n, err := rand.Read(buf)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Read %d bytes: 0x%x\n", n, buf)
}
