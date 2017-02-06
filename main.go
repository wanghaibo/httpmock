package main

import (
	"encoding/json"
	"fmt"
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
	_ = kv
}

func mock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mock" + r.Host + r.URL.String()))
}

func adminMocks(w http.ResponseWriter, r *http.Request) {
	var mock Mock
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&mock)
		if err != nil {
			panic(err)
		}
		//@todo  profile
		content, err := json.Marshal(mock)
		if err != nil {
			panic(err)
		}
		fmt.Println(content)
		err = kv.Put([]byte(mock.Url), content)

		if err != nil {
			panic(err)
		}
	}
}

func main() {
	finish := make(chan bool)
	mockServer := http.NewServeMux()
	mockServer.HandleFunc("/", mock)

	adminServer := mux.NewRouter()
	adminServer.HandleFunc("/mocks/", adminMocks)

	go func() {
		log.Fatal(http.ListenAndServe(":8090", mockServer))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":8091", adminServer))
	}()

	<-finish
}
