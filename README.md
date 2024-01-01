# Static File Server
### Serving files online can be challenging, so why not make it difficult as an executable?

## How to use
- Move your files to `files`
- You may want some secure traffic while serving, look at our usage below:
```
Usage of ./fileserver:
  -address string
        The address to listen to (default ":80")
  -cert string
        Path to the tls certificate
  -key string
        Path to the tls key
  -quic tlsCert
        Must have tlsCert and `tlsKey` to enable QUIC (default true)
  ```
- Those can be also set with environment variables as `ADDRESS`, `CERT`, `KEY` and `QUIC`.
- A simple start for HTTP-only is easy as `go run .` or build the binary and run as `./static-file-server`
- You can pass the arguments in CLI to start with QUIC `./static-file-server -cert <path/to/cert> -key <path/to/key>`
- You can also set as an environment variable `ADDRESS=":8080"" ./static-file-server`, and your file server will be listening on port 8080.
###### QUIC can be disabled with `-quic=false` or `QUIC="<any value>"`

---

#### Made by @Dviih
