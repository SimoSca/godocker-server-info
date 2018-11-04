package main

// NOTA:
// devo specificare la directory per ottenere la giusta dipendenza,
// ma poi il nome importato e' preso dal file `package` dentro a quella directory, se non specifico altrimenti
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"

	"github.com/SimoSca/godocker-server-info/godock"
	"github.com/SimoSca/godocker-server-info/server"
	enwssocket "github.com/SimoSca/godocker-server-info/websocket"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "docker" {
		godock.GetDockerHosts()
	} else {

		// Setup WebSocket
		mxWS := http.NewServeMux()
		mxWS.Handle("/", websocket.Handler(enwssocket.WsHandler))
		fmt.Println("Listen and serve websocket on :8082")
		go func() {
			http.ListenAndServe(":8082", mxWS)
		}()

		//  mxHTTP := http.NewServeMux()
		//  mxHTTP.HandleFunc("/", homeHandler)
		//  fmt.Println("Listen and serve HTTP on :8080")
		//  http.ListenAndServe(":8080", mxHTTP)

		// Setup Http
		port := ":8080"

		fmt.Println("Server starts listening on port ", port)
		http.HandleFunc("/", server.Handler)
		log.Fatal(http.ListenAndServe(port, nil))
	}
}
