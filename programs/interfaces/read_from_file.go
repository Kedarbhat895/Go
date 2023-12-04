package main

import (
	"fmt"
	"io"
	"os"
)

type terminalWriter struct{}

func main() {

	file, error := os.Open(os.Args[1])
	if error != nil {
		fmt.Println("Error occured: ", error)
		os.Exit(1)
	}

	tw := terminalWriter{}
	io.Copy(tw, file)

}

func (terminalWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
