package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
)

func makeInfoEndpoint(svc RedisInfoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(infoRequest)
		v, err := svc.Info(req.Addr, req.Passwd)
		if err != nil {
			return infoResponse{v, err.Error()}, nil
		}
		return infoResponse{v, ""}, nil
	}
}

func decodeInfoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request infoRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type infoRequest struct {
	Addr   string `json:"addr"`
	Passwd string `json:"passwd"`
}

type infoResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
