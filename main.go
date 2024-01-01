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
	addr    = os.Getenv("ADDRESS")
	tlsCert = os.Getenv("CERT")
	tlsKey  = os.Getenv("KEY")
	quic    = os.Getenv("QUIC") != ""
)

func init() {
	flag.StringVar(&addr, "address", ":80", "The address to listen to")
	flag.StringVar(&tlsCert, "cert", "", "Path to the tls certificate")
	flag.StringVar(&tlsKey, "key", "", "Path to the tls key")
	flag.BoolVar(&quic, "quic", true, "Must have `tlsCert` and `tlsKey` to enable QUIC")

	flag.Parse()

	if quic && (tlsCert == "" || tlsKey == "") {
		quic = false
	}

