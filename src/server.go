package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var version string = "v2"

type OS struct {
	Time    string
	Host    string
	OSType  string
	Version string
}

type newAPIHandler struct{}

func (eh *newAPIHandler) getOperatingSystemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	host, _ := os.Hostname()
	ostype := runtime.GOOS

	msg := OS{
		time.Now().Format(time.RFC850),
		host,
		ostype,
		version}

	json.NewEncoder(w).Encode(msg)
}

func (eh *newAPIHandler) optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	handler := newAPIHandler{}

	r := mux.NewRouter()
	apirouter := r.PathPrefix("/api").Subrouter()
	apirouter.Methods("GET").Path("/os").HandlerFunc(handler.getOperatingSystemHandler)
	apirouter.Methods("OPTIONS").Path("/os").HandlerFunc(handler.optionsHandler)

	server := cors.Default().Handler(r)

	//Azure Functions sets FUNCTIONS_CUSTOMHANDLER_PORT to a integer value.  Must append :
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	fmt.Print("Listening on ", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, server))
}
