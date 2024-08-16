package main

import (
	"bytes"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
)

func main() {
	client := &http.Client{}
	//var pathToFle = "texti.txt"
	//fil, err := os.Open(pathToFle)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var ma map[string]string
	//ma = map[string]string{
	//	"title":       "My Document",
	//	"author":      "Matt Aimonetti",
	//	"description": "A document with all the Go programming language secrets",
	//}
	//req, _ := makeRequest("http://localhost:9999/multic", fil, "file", ma)
	//response, err := client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer response.Body.Close()
	//b, err := httputil.DumpResponse(response, true)
	//log.Println(string(b))

	buf := make([]byte, 0)

	body := bytes.NewBuffer(buf)
	reqToGetMultipart, err := http.NewRequest("GET", "http://localhost:9999/wanna_get_multic", body)
	reqToGetMultipart.Header.Set("Accept", "multipart/form-data")
	if err != nil {
		panic(err)
	}

	do, err := client.Do(reqToGetMultipart)
	if err != nil {
		log.Printf("Error while downloading multipart file: %v", err)

	}
	_, params, err := mime.ParseMediaType(do.Header.Get("Content-Type"))
	if err != nil {
		log.Printf("Error while parsing multipart file: %v", err)
	}
	mr := multipart.NewReader(do.Body, params["boundary"])
	for p, err := mr.NextPart(); err == nil; p, err = mr.NextPart() {
		all, _ := io.ReadAll(p)
		log.Println(string(all))
	}

}

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

//func makeRequest(url string, fil *os.File, fieldName string, params map[string]string) (*http.Request, error) {
//
//	//fileContent, err := io.ReadAll(fil)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//fileInfo, err := fil.Stat()
//	//if err != nil {
//	//	return nil, err
//	//}
//	//
//	//fileName := fileInfo.Name()
//	//
//	//fil.Close()
//
//	var body []byte
//	bodyWriter := bytes.NewBuffer(body)
//	multiPartWriter := multipart.NewWriter(bodyWriter)
//
//	filePart, err := multiPartWriter.CreateFormFile(fieldName, fileName)
//	if err != nil {
//		return nil, err
//	}
//	filePart.Write(fileContent)
//
//	for k, p := range params {
//		multiPartWriter.WriteField(k, p)
//	}
//	return http.NewRequest("POST", url, bodyWriter)
//}
