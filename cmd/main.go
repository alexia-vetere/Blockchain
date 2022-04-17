package main

import (
	"blockchain_proyect/blockchain/utils"
	"fmt"
)

func main() {
	fmt.Println(utils.FindNeighbords("127.0.0.1", 5000, 0, 3, 5000, 5003))
}
