package http

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
    api "github.com/teamlint/mons/sample2/services/user/api"
    "github.com/teamlint/mons/sample2/services/user/internal/endpoint"
)

func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]kithttp.ServerOption) http.Handler {
	m := http.NewServeMux()
	makeFindHandler(m, endpoints, options["Find"])
	makeUpdateHandler(m, endpoints, options["Update"])
	return m
}

func makeFindHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []kithttp.ServerOption) {
	m.Handle("/user/find", kithttp.NewServer(endpoints.FindEndpoint, decodeFindRequest, kithttp.EncodeJSONResponse, options...))
}

func makeUpdateHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []kithttp.ServerOption) {
	m.Handle("/user/update", kithttp.NewServer(endpoints.UpdateEndpoint, decodeUpdateRequest, kithttp.EncodeJSONResponse, options...))
}

func decodeFindRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request api.FindUserRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func decodeUpdateRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request api.UpdateUserRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

