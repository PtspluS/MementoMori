package main

import (
	"fmt"
	coldfire "github.com/redcode-labs/Coldfire"
	analyse "inferno.com/main/annalyse"
)

func main() {
	fmt.Println(analyse.Analyser())
	coldfire.PrintGood("Ok")
	return
}
