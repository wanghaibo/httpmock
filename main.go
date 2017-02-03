package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	flag "github.com/spf13/pflag"
)

var mocks []Mock
var kv KvStore
var bucketName []byte
var datapath string

func init() {
	flag.StringVar(&datapath, "datapath", "", "path to store mockinfo")
	flag.Parse()

	if datapath == "" {
		log.Fatal("miss datapath")
		return
	}
	bucketName = []byte("httpmock")
	kv, err := NewKvStore(map[string]interface{}{"bucketName": bucketName, "datapath": datapath})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func mock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mock" + r.Host + r.URL.String()))
}

func adminMocks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	} else if r.Method == "POST" {

	}
}

func adminMock(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
	case "UPDATE":
	case "DELETE":
	}
}

func main() {
	finish := make(chan bool)
	mockServer := http.NewServeMux()
	mockServer.HandleFunc("/", mock)

	adminServer := mux.NewRouter()
	adminServer.HandleFunc("/mocks/", adminMocks)
	adminServer.HandleFunc("/mocks/{sum}", adminMock)

	go func() {
		log.Fatal(http.ListenAndServe(":8090", mockServer))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8091", adminServer))
	}()

	<-finish
}
