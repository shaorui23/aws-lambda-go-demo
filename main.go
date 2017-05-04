package main

import (
	"log"
	"net/http"

	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
	"github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
)

// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

func init() {
	log.Println("Initialize aws-lambda-go demo site")

	ln := net.Listen()

	// Amazon API Gateway Binary support out of the box.
	Handle = apigatewayproxy.New(ln, []string{"image/png"}).Handle

	// Any Go framework complying with the Go http.Handler interface can be used.
	// This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, Goa, etc.
	go http.Serve(ln, http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World From AWS-lambda-go!"))
}
