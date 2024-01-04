package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"io/fs"
	"net/http"
	"os"
)

//go:embed files
var files embed.FS

var (
	addr    = envDefault("ADDRESS", ":80")
	tlsCert = os.Getenv("CERT")
	tlsKey  = os.Getenv("KEY")
	quic    = os.Getenv("QUIC") != ""
)

func envDefault(key, _default string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return _default
}

func init() {
	flag.StringVar(&addr, "address", addr, "The address to listen to")
	flag.StringVar(&tlsCert, "cert", tlsCert, "Path to the tls certificate")
	flag.StringVar(&tlsKey, "key", tlsKey, "Path to the tls key")
	flag.BoolVar(&quic, "quic", quic, "Must have tlsCert and tlsKey to enable QUIC")

	flag.Parse()

	if quic && (tlsCert == "" || tlsKey == "") {
		quic = false
	}

	if (tlsCert != "" || tlsKey != "") && addr == ":80" {
		addr = ":443"
	}
}

func main() {
	sub, err := fs.Sub(files, "files")
	if err != nil {
		panic(err)
	}

	e := make(chan error, 1)

	handler := http.FileServer(http.FS(sub))

	fmt.Printf("Listening and serving at %v\n", addr)

	if tlsCert != "" || tlsKey != "" {
		if quic {
			e <- http3.ListenAndServe(addr, tlsCert, tlsKey, handler)
		} else {
			e <- http.ListenAndServeTLS(addr, tlsCert, tlsKey, handler)
		}
	} else {
		e <- http.ListenAndServe(addr, http.FileServer(http.FS(sub)))
	}

	for {
		select {
		case err = <-e:
			panic(err)
		}
	}
}
