package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//Always check the errors

func main() {

	f, err := os.Open("name.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))

}
