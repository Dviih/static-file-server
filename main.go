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

