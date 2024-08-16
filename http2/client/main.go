package main

import (
	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/http2"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	client := makeClient()
	ti := time.Now()
	for i := 0; i < 100; i++ {
		reqest := makeRequest()
		resp := makeRequestWithClient(client, reqest)
		if resp == nil || resp.StatusCode != 200 {
			log.Printf("Request failed")
		}
		log.Println("DONE!!!")
		log.Printf("StatusCode: %d\n", resp.StatusCode)
		//makePrintResponse(resp)

	}
	log.Println("time::: ", time.Since(ti))

}

func makePrintResponse(resp *http.Response) {

	log.Println("-=-=-=-=-=-=-=   HEaDErs   -=-=-=-==-=-=-=--=")
	log.Println("proto ", resp.Proto)
	log.Println("Gotta response with headers", resp.Header)

	println()

	log.Println("-=-=-=-=-=-=-=   BODY   -=-=-=-==-=-=-=--=")
	respBo, _ := io.ReadAll(resp.Body)
	log.Println(respBo)

}

func makeRequestWithClient(client *http.Client, reqest *http.Request) *http.Response {

	resultDo, err := client.Do(reqest)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return nil
	}
	return resultDo

}

func makeRequest() *http.Request {

	req, _ := http.NewRequest("GET", "https://localhost:9000/gorik", nil)
	return req
}

func makeClient() *http.Client {

	client := &http.Client{}
	tlsConfig := makeTlsconfig()
	client.Transport = &http2.Transport{
		TLSClientConfig: tlsConfig,
	}
	//client.Transport = &http.Transport{
	//	TLSClientConfig: tlsConfig,
	//}
	return client

}

func makeTlsconfig() *tls.Config {

	rootCAS := x509.NewCertPool()
	cert, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}
	rootCAS.AppendCertsFromPEM(cert)
	tlsConfig := tls.Config{RootCAs: rootCAS}
	return &tlsConfig
}
