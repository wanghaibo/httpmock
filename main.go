package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	flag "github.com/spf13/pflag"
)

var mocks []Mock
var kv KvStore
var bucketName []byte
var datapath string
var adminPort string
var serverPort string

func init() {
	var err error
	flag.StringVar(&datapath, "datapath", "", "path to store mockinfo")
	flag.StringVar(&serverPort, "serverport", "80", "server port")
	flag.StringVar(&adminPort, "adminport", "8089", "admin port")
	flag.Parse()

	if datapath == "" {
		log.Fatal("miss datapath")
		return
	}
	bucketName = []byte("httpmock")
	kv, err = NewKvStore(map[string]interface{}{"bucketName": bucketName, "datapath": datapath})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func mock(w http.ResponseWriter, r *http.Request) {
	var schema string
	var url []byte
	var mock Mock
	if r.TLS != nil {
		schema = "https://"
	} else {
		schema = "http://"
	}
	url = []byte(schema + r.Host + r.RequestURI)
	value, err := kv.Get(url)
	if err != nil {
		panic(err)
	}
	if value == nil {
		panic("url has not be set")
	}
	err = json.Unmarshal(value, &mock)
	if err != nil {
		panic(err)
	}
	for key, value := range mock.Headers {
		w.Header().Set(key, value)
	}
	w.Write([]byte(mock.Body))
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
		log.Fatal(http.ListenAndServe(":"+serverPort, mockServer))
	}()

	go func() {
		log.Fatal(http.ListenAndServe(":"+adminPort, adminServer))
	}()

	<-finish
}
