package main

import (
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
)

func main() {

	http.HandleFunc("/multic", handleMultik)
	http.HandleFunc("/wanna_get_multic", giveMultipart)
	http.ListenAndServe(":9999", nil)
}

func giveMultipart(w http.ResponseWriter, r *http.Request) {
	//	::: check if client is EXPECTING multipart data
	//	:::
	mediatype, _, err := mime.ParseMediaType(r.Header.Get("Accept"))
	if err != nil {
		log.Printf("Error parsing media type: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if mediatype != "multipart/form-data" {

		log.Printf("media type not 'multipart/form-data': %v", mediatype)
		http.Error(w, "media type not 'multipart/form-data'", http.StatusBadRequest)
	}

	//	:::
	//	::: check if client is EXPECTING multipart data
	mw := multipart.NewWriter(w)
	w.Header().Set("Content-Type", mw.FormDataContentType())
	for _, value := range getValues() {

		fv, err := mw.CreateFormField("value")
		if err != nil {
			log.Printf("Error creating form field: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		_, err = fv.Write([]byte(value))
		if err != nil {
			log.Printf("Error writing form field: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		mw.Close()

	}
}

func getValues() [][]byte {
	return [][]byte{[]byte{132, 21}}
}

func handleMultik(writer http.ResponseWriter, request *http.Request) {

	all, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(all))
	writer.Write([]byte("Gotta"))
}
