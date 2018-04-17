package osssrv

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
	"log"
)

type Server struct {
	Addr string
	Router *mux.Router
	Debug bool
}

func (srv Server) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	if srv.Debug {
		log.Printf("%s %s", req.Method, req.URL.Path)
	}
	srv.Router.ServeHTTP(w, req)
}




func (srv Server) run(){

	api := new(FileApi)

	srv.Router.HandleFunc("/files/{query_id}/info", api.GetFileInfo).Methods("GET")
	srv.Router.HandleFunc("/files/{query_id}", api.AccessFile).Methods("GET")
	srv.Router.HandleFunc("/files/{query_id}", api.UpdateFileInfo).Methods("PUT", "PATCH")
	srv.Router.HandleFunc("/files/{query_id}", api.DeleteFile).Methods("DELETE")
	srv.Router.HandleFunc("/files", api.UploadFile).Methods("POST")
	srv.Router.HandleFunc("/files", api.GetFileList).Methods("GET")

	http.Handle("/", srv)
	httpSrv := &http.Server{
		Handler: srv,
		Addr: srv.Addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	fmt.Printf("listen and server: %s\n", srv.Addr)
	log.Fatal(httpSrv.ListenAndServe())
}
