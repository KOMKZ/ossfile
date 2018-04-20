package ossfile

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"fmt"
	"log"
	"os"
	"encoding/json"
)

type Server struct {
	Addr string
	Router *mux.Router
	Req *Request
	Debug bool
	RuntimeDir string
}

type Request struct {
	*http.Request
	curContType string
}

func (r *Request)GetContType() (conType string){
	if r.curContType != "" {
		return r.curContType
	}
	conType = r.Header.Get("Content-Type")
	log.Println(conType)
	r.curContType = conType
	return conType
}



func (r *Request)IsJsonForm() bool{
	if t := r.GetContType(); t != "application/json" {
		return false
	}else{
		return true
	}
}

func (r *Request)IsMultipartForm() bool{
	if t := r.GetContType(); t != "multipart/form-data" {
		return false
	}else{
		return true
	}
}

func (r *Request)ParseJsonForm(jsonData interface{}) error{
	err := json.NewDecoder(r.Body).Decode(jsonData)
	if err != nil {
		return err
	}
	return nil

}

func (srv *Server) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	if srv.Debug {
		log.Printf("%s %s", req.Method, req.URL.Path)
	}
	srv.Req = &Request{req, ""}
	srv.Router.ServeHTTP(w, req)
}


func (srv *Server) init() error{
	if srv.RuntimeDir == "" {
		return fmt.Errorf("Server RumtimeDir can't not be empty")
	}
	err := os.Mkdir(srv.RuntimeDir, 0777)
	if !os.IsExist(err) {
		return err
	}

	return nil
}

func (srv *Server) run(){

	err := srv.init()
	if err != nil {
		log.Fatal(err)
	}

	api := FileApi{srv: srv}
	siteApi := SiteApi{srv: srv}

	srv.Router.HandleFunc("/", siteApi.Index).Methods("GET")
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
