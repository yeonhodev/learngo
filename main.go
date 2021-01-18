package main

import (
	"fmt"

	"github.com/yeonhodev/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)
}
