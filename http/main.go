package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Printf("bs: %v\n", string(bs))

	return len(bs), nil

}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("resp: %v\n", resp)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}
