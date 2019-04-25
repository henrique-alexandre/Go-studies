package main

import "fmt"

func main() {
	name := "Todd McLeod"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Printf("%T \n", tpl)
	fmt.Println(tpl)
}
