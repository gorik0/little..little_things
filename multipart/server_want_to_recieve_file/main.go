package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/multic", handleMultik)
	http.ListenAndServe(":9999", nil)
}

func handleMultik(writer http.ResponseWriter, request *http.Request) {

	all, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(all))
	writer.Write([]byte("Gotta"))
}
