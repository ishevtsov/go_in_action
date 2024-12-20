package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	// Write a string to the buffer.
	b.Write([]byte("Hello"))

	// Use Fprintf to concatenate a string to the Buffer.
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to stdout.
	io.Copy(os.Stdout, &b)
}
