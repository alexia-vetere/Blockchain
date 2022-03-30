package main

import (
	"blockchain_proyect/blockchain/block"
	"blockchain_proyect/blockchain/wallet"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

// Creación del puerto:
type BlockchainServer struct {
	port uint16
}

// Función de retorno del puerto:
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{
		port: port,
	}
}

// Método para obtener la blockchain utilizada:
func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		// Creación de wallet:
		minerWallet := wallet.NewWallet()
		// Nodo Blockchain:
		bc = block.NewBlockchain(minerWallet.BlockchainAddress(), bcs.Port())
		cache["blockchain"] = bc
		log.Printf("private_key %v", minerWallet.PrivateKeyStr())
		log.Printf("public_key %v", minerWallet.PublicKeyStr())
		log.Printf("blockchain_address %v", minerWallet.BlockchainAddress())

	}
	return bc
}

// Método de retorno del puerto:
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

//
func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Println("ERROR: Invalid HTTP Method")
	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	fmt.Println("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil)
	log.Fatal(
		http.ListenAndServe("localhost:"+strconv.Itoa(int(bcs.Port())), nil))
}
