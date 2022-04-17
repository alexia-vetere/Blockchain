package main

import (
	"blockchain_proyect/blockchain/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.IsFoundHost("127.0.0.1", 5000))
}
