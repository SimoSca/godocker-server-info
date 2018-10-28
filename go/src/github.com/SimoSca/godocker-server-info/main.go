package main

// NOTA:
// devo specificare la directory per ottenere la giusta dipendenza, 
// ma poi il nome importato e' preso dal file `package` dentro a quella directory, se non specifico altrimenti
import (
	"fmt"
	// "io/util"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/SimoSca/godocker-server-info/enlogger"
	"github.com/SimoSca/godocker-server-info/godock"
)

/// PAGE
type Page struct {
	Title string
	// The Body element is a []byte rather than string because that is the type expected by the io libraries
    Body  []byte
}

// func (p *Page) save() error {
//     filename := p.Title + ".txt"
//     return ioutil.WriteFile(filename, p.Body, 0600)
// }

// func loadPage(title string) (*Page, error) {
//     filename := title + ".txt"
//     body, err := ioutil.ReadFile(filename)
//     if err != nil {
//         return nil, err
//     }
//     return &Page{Title: title, Body: body}, nil
// }

func handler(w http.ResponseWriter, r *http.Request) {
	// var dump della richiesta
	// fmt.Printf("%#v", r)
	// fmt.Print("\n")

	fmt.Println(r.URL.Path)
	p := &Page{Title: "oops", Body: []byte("ooops, something went wrong!!!")}
	if r.URL.Path == "/" {
		p = &Page{Title: "Home", Body: []byte("Welcome to HOME!")}
	}else{
		p = &Page{Title: r.URL.Path[1:], Body: []byte("another page")}
	}
	
	// Bundling static assets inside of Go binaries. 
	// 		Note that pack takes path relative to this script file, not cwd!
	boxT := packr.NewBox("./assets/templates")
	s := boxT.String("home.html")

	// t, err := template.ParseFiles( "./go/src/github.com/SimoSca/godocker-server-info/assets/templates/home.html" )
	t, err := template.New("home").Parse( s )
	if err != nil {
		logger.Print("Error while processing template.Parse()")
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
		logger.Print("Error while processing t.Execute()")
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
	if os.Args[1] == "docker" {
		godock.PrintList()
		godock.PrintEvents()
	}else{
		port := ":8080"

		fmt.Println("Server starts listening on port ", port)
    	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(port, nil))
	}
}
