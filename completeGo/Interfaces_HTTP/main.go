package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	r, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999)
	//r.Body.Read(bs)
	//fmt.Println(string(bs))

	io.Copy(os.Stdout, r.Body)

	lw := logWriter{}

	io.Copy(lw, r.Body)

}

func (l logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote out this number of bytes: ", len(bs))
	return len(bs), nil
}
