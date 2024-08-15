package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	client := &http.Client{}
	var pathToFle = "texti.txt"
	fil, err := os.Open(pathToFle)
	if err != nil {
		log.Fatal(err)
	}
	var ma map[string]string
	ma = map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets",
	}
	req, _ := makeRequest("http://localhost:9999/multic", fil, "file", ma)
	//log.Printf("%+v", req)

	//boGet := make([]byte, 2)
	//bu := new(bytes.Buffer)
	//bu.Read(boGet)
	//reqGet, _ := http.NewRequest("GET", "http://localhost:9999/multic", nil)
	//respGet, err := client.Do(reqGet)
	//log.Println(respGet)
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//println(response.Body)
	defer response.Body.Close()
	//log.Printf("%+v", response)

	b, err := httputil.DumpResponse(response, true)

	p := make([]byte, 10)
	response.Body.Read(p)
	log.Println(string(p))
	log.Println(string(b))
	//
	//_, err = response.Body.Read(input)
	//if err != nil {
	//	panic(err)
	//
	//}
	//response.Body.Close()
	//log.Println(input)
	//out := bufio.NewWriter(os.Stdout)
	//defer out.Flush()
	//fmt.Fprintln(out, input)
	//fmt.Fprintln(out, "dfdf")

}

func makeRequest(url string, fil *os.File, fieldName string, params map[string]string) (*http.Request, error) {

	fileContent, err := io.ReadAll(fil)
	if err != nil {
		return nil, err
	}

	fileInfo, err := fil.Stat()
	if err != nil {
		return nil, err
	}

	fileName := fileInfo.Name()

	fil.Close()

	var body []byte
	bodyWriter := bytes.NewBuffer(body)
	multiPartWriter := multipart.NewWriter(bodyWriter)

	filePart, err := multiPartWriter.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil, err
	}
	filePart.Write(fileContent)

	for k, p := range params {
		multiPartWriter.WriteField(k, p)
	}
	return http.NewRequest("POST", url, bodyWriter)
}
