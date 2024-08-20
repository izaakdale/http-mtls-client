package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	srvAddr := os.Getenv("SERVER_ADDR")
	crt, err := tls.LoadX509KeyPair(os.Getenv("CLIENT_CRT"), os.Getenv("CLIENT_KEY"))
	if err != nil {
		panic(err)
	}

	// fCa, err := os.OpenFile(os.Getenv("ROOT_CA"), os.O_RDONLY, os.ModeTemporary)
	// if err != nil {
	// 	panic(err)
	// }
	// pemBytes, err := io.ReadAll(fCa)
	// if err != nil {
	// 	panic(err)
	// }

	// caCertPool := x509.NewCertPool()
	// ok := caCertPool.AppendCertsFromPEM(pemBytes)
	// if !ok {
	// 	panic("error appending certs from pem")
	// }

	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		caCertPool = x509.NewCertPool()
	}

	cli := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{crt},
				RootCAs:      caCertPool,
			},
		},
	}
	for {
		req, err := http.NewRequest(http.MethodGet, srvAddr+"/ping", nil)
		if err != nil {
			panic(err)
		}
		resp, err := cli.Do(req)
		if err != nil {
			panic(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		log.Printf("%+v\n", string(bodyBytes))
		time.Sleep(3 * time.Second)
	}
}
