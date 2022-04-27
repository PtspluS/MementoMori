package main

import (
	"fmt"

	coldfire "github.com/redcode-labs/Coldfire"
	analyse "inferno.com/annalyse"
)

func main() {
	var infos map[string]string = analyse.Infos()
	fmt.Println(infos)
	fmt.Println(analyse.UserInfo(infos["user"]))
	coldfire.PrintGood("Ok")
}
