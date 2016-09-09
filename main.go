package main

import (
	"flag"
	"net/http"
	"os"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	ctx := context.Background()
	logger := log.NewLogfmtLogger(os.Stderr)

	port := flag.String("port", "8080", "http server port")
	flag.Parse()

	var svc RedisInfoService
	svc = redisInfoService{}
	svc = loggingMiddleware{logger, svc}

	infoHandler := httptransport.NewServer(
		ctx,
		makeInfoEndpoint(svc),
		decodeInfoRequest,
		encodeResponse,
	)

	http.Handle("/info", infoHandler)
	logger.Log("msg", "HTTP", "addr", *port)
	logger.Log("err", http.ListenAndServe(":"+*port, nil))

}
