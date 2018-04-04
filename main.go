package main

import (
	"net/http"
	"fmt"
	"log"
	"strings"
)

type App struct{
	title string
}
func (app App) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	route := fmt.Sprintf("%s %s", strings.ToUpper(req.Method), req.URL.Path)
	switch route {
	case "GET /file":
		fmt.Fprintf(w, "%s\n", route)
	case "GET /file/:id":
		fmt.Fprintf(w, "%s\n", route)
	case "POST /file":
		fmt.Fprintf(w, "%s\n", route)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error route %s\n",route)
	}
}
func main()  {
	const addr = "localhost:8080"
	app := App{"hello world"}
	log.Printf("listen and serve : %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, app))
}

