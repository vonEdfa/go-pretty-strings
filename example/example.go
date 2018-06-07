package main

import (
	"fmt"

	pretty "github.com/vonEdfa/go-pretty-strings"
)

func main() {
	fmt.Println(pretty.BorderSquare("My example bordered box", "Free Willy!", "Super Silly!\nAlmost like Billy!"))
}
