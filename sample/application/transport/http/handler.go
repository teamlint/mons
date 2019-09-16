package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/teamlint/mons/sample/application/command"
	"github.com/teamlint/mons/sample/application/endpoint"
	"github.com/teamlint/mons/sample/application/query"
)

func NewHTTPHandler(endpoints endpoint.Endpoints, options map[string][]kithttp.ServerOption) http.Handler {
	m := http.NewServeMux()
	makeFindUserHandler(m, endpoints, options["FindUser"])
	makeUpdateUserHandler(m, endpoints, options["UpdateUser"])
	return m
}

func makeFindUserHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []kithttp.ServerOption) {
	m.Handle("/user/find", kithttp.NewServer(endpoints.FindUserEndpoint, decodeFindUserRequest, kithttp.EncodeJSONResponse, options...))
}
func makeUpdateUserHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []kithttp.ServerOption) {
	m.Handle("/user/update", kithttp.NewServer(endpoints.UpdateUserEndpoint, decodeUpdateUserRequest, kithttp.EncodeJSONResponse, options...))
}

func decodeFindUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request query.UserQuery
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request command.UpdateUserCommand

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return &request, nil
}

func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
