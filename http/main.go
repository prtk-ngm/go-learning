package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"crypto/x509"
)

type logWriter struct{}

func main() {

	//certificate
	certPool := x509.NewCertPool()
	conf := &tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	transport := &http.Transport{TLSClientConfig: conf}
	client := &http.Client{Transport: transport}
	headers := http.Header{}
	resp, err := http.Get("http://localhost:50052/health")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp.StatusCode)
	lw := logWriter{}
	io.Copy(lw, resp.Body)

	resp.Body.Close()

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("len of bytes:", len(bs))

	return len(bs), nil

}
