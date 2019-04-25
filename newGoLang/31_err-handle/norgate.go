package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("an error occured: %v %v %v ", n.lat, n.long, n.err)
}

func main() {

	_, err := sqrt(-20)
	if err != nil {
		log.Println(err)
	}

}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("cannot do n.n.: %v", f)
		return 0, norgateMathError{"50.50 N", "99.99 W", e}
	}
	return 40, nil
}
