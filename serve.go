package main

import (
	"encoding/json"
	"github.com/bmizerany/pat"
	"io"
	"log"
	"net/http"
)

type Message struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func PutJSON(w http.ResponseWriter, req *http.Request) {
	message := Message{
		X: req.URL.Query().Get(":x"),
		Y: req.URL.Query().Get(":y"),
	}

	j, err := json.Marshal(message)

	if err != nil {
		w.Header().Add("Content-Type", "text/html")
		io.WriteString(w, "ERROR: "+err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, string(j))
	}
}

func GetJSON(w http.ResponseWriter, req *http.Request) {
	p := req.URL.Query().Get("p")
	var message Message
	err := json.Unmarshal([]byte(p), &message)
	w.Header().Add("Content-Type", "text/plain")
	if err != nil {
		io.WriteString(w, "ERROR: "+err.Error())
	} else {
		io.WriteString(w, "OK")
	}
}

func main() {
	m := pat.New()
	m.Get("/put/:x/:y", http.HandlerFunc(PutJSON))
	m.Get("/get", http.HandlerFunc(GetJSON))
	http.Handle("/", m)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
