package main

import (
	"blockchain_proyect/blockchain/utils"
	"blockchain_proyect/blockchain/wallet"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

// Ruta de templates:
const tempDir = "templates"

// Clase para el servidor de la wallet:
type WalletServer struct {
	port    uint16
	gateway string
}

// Función para obtener el servidor de la wallet:
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

// Método para obtener el puerto del servidor:
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

// Método para obtener el gateway del servidor creado:
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

// Método para la redircción al template:
func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, err := template.ParseFiles(path.Join(tempDir, "index.html"))
		if err != nil {
			log.Println("ERROR: Invalid Index Directory:", path.Join(tempDir, "index.html"))
		} else {
			t.Execute(w, "")
		}
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

// Método que retorna la información del index:
func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := wallet.NewWallet()
		m, err := myWallet.MarshalJSON()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(w, string(m[:]))

	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

// Método para retornar datos al enviar dinero:
func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Panicf("ERROR: %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}
		if !t.Validate() {
			log.Panicln("ERROR: missing field(s)")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*t.SenderPrivateKey, publicKey)
		value, err := strconv.ParseFloat(*t.Value, 32)
		if err != nil {
			log.Panicln("ERROR: parse error")
			io.WriteString(w, string(utils.JsonStatus("fail")))
		}
		value32 := float32(value)

		w.Header().Add("Content-Type", "application/json")
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Invalid HTTP Method")
	}
}

// Método para correr el servidor:
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(int(ws.Port())), nil))
}
