package main

import (
	"log"
	"net/http"
)

const (
	MACHINE_NAME = "Frederick"
)

func RequestMachineInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello my name is", MACHINE_NAME)
}

func main() {
	http.HandleFunc("/query", RequestMachineInfo)
	log.Println("Serving at port 1234")
	http.ListenAndServe(":1234", nil)
}
