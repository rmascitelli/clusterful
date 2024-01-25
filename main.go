package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var SERVICE_NAME string
var PORT int

// TODO: Cannot write to log file from Docker container?
func WriteMachineInfoToFile() {
	d1 := []byte(SERVICE_NAME + ":" + strconv.Itoa(PORT) + "\n")
	f, err := os.OpenFile("/tmp/test_file", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(d1)
}

func RequestMachineInfo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello my name is %s:%d\n", SERVICE_NAME, PORT)
}

func main() {
	service_name := flag.String("service_name", "", "Name of the service to spawn")
	port := flag.Int("port", 0, "Port to listen to")
	flag.Parse()

	fmt.Println("service_name=", *service_name)
	fmt.Println("port=", *port)
	SERVICE_NAME = *service_name
	PORT = *port

	//WriteMachineInfoToFile()

	http.HandleFunc("/query", RequestMachineInfo)
	log.Printf("Serving at port %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}
