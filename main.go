package main

import (
	"fmt"

	"github.com/yeonhodev/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "nicolas", Balance: 1000}
	fmt.Println(account)
}
