package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	r := strings.NewReader("Henrique")

	io.Copy(f, r)

	_, e := os.Open("nofile.txt")
	if e != nil {
		fmt.Println(e)
		log.Println(e)
		log.Fatalln(e)
		panic(e)
	}

}
