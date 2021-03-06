package bootstrap

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/riomhaire/systeminfo/infrastructure/web/rest"
	"github.com/urfave/negroni"
)

const VERSION = "SystemInfo Version 1.2"

// Application - Structure containing app level info
type Application struct {
	Port    int
	Negroni *negroni.Negroni
}

// Initialize - Application  data strure and set up the routers etc
func (a *Application) Initialize() {
	// Config
	port := flag.Int("port", 3333, "Port to use")
	flag.Parse()

	mux := http.NewServeMux()
	negroni := negroni.Classic()
	// Add handlers
	mux.HandleFunc("/metrics", rest.Metrics)

	negroni.UseHandler(mux)
	// Remember
	a.Port = *port
	a.Negroni = negroni

}

// Run - start up Application listener on given port
func (a *Application) Run() {
	fmt.Printf("Starting %v on port %v\n", VERSION, a.Port)
	a.Negroni.Run(fmt.Sprintf(":%d", a.Port))
}
