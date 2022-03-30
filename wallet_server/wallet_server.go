package main

import (
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

// Método para correr el servidor:
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(int(ws.Port())), nil))
}
