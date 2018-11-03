package main

// NOTA:
// devo specificare la directory per ottenere la giusta dipendenza,
// ma poi il nome importato e' preso dal file `package` dentro a quella directory, se non specifico altrimenti
import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SimoSca/godocker-server-info/godock"
	"github.com/SimoSca/godocker-server-info/server"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "docker" {
		godock.GetDockerHosts()
	} else {
		port := ":8080"

		fmt.Println("Server starts listening on port ", port)
		http.HandleFunc("/", server.Handler)
		log.Fatal(http.ListenAndServe(port, nil))
	}
}
