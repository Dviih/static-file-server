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

