package main

func main() {

	switch {
	case false:
		println("This shouldn't print")
	case (2 == 2):
		println("This should print")
		fallthrough
	case (3 == 4):
		println("This should print")

	default:
		println(10)
	}

}
