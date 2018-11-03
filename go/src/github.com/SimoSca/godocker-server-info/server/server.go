package server

import (
	"fmt"
	"html/template"
	"net/http"

	logger "github.com/SimoSca/godocker-server-info/enlogger"
	"github.com/SimoSca/godocker-server-info/godock"
	"github.com/gobuffalo/packr"
)

/// PAGE
type Page struct {
	Title string
	// The Body element is a []byte rather than string because that is the type expected by the io libraries
	Body []byte
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

func Handler(w http.ResponseWriter, r *http.Request) {
	// var dump della richiesta
	// fmt.Printf("%#v", r)
	// fmt.Print("\n")

	fmt.Println(r.URL.Path)
	p := &Page{Title: "oops", Body: []byte("ooops, something went wrong!!!")}
	if r.URL.Path == "/" {
		p = &Page{Title: "Home", Body: []byte("Welcome to HOME!")}
	} else {
		p = &Page{Title: r.URL.Path[1:], Body: []byte("another page")}
	}

	// Bundling static assets inside of Go binaries.
	// 		Note that pack takes path relative to this script file, not cwd!
	boxT := packr.NewBox("../assets/templates")
	s, _ := boxT.FindString("home.html")

	// t, err := template.ParseFiles( "./go/src/github.com/SimoSca/godocker-server-info/assets/templates/home.html" )
	t, err := template.New("home").Parse(s)
	if err != nil {
		logger.Print("Error while processing template.Parse()")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Page  *Page
		Hosts []godock.PortMap
	}{p, godock.GetDockerHosts()}
	err = t.Execute(w, data)
	if err != nil {
		logger.Print("Error while processing t.Execute()")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
