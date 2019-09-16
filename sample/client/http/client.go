package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/teamlint/mons/sample/application/dto"
	"github.com/teamlint/mons/sample/application/endpoint"
	"github.com/teamlint/mons/sample/application/service"
	handler "github.com/teamlint/mons/sample/application/transport/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

func New(instance string, options map[string][]kithttp.ClientOption) (service.UserService, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	var findUserEndpoint kitendpoint.Endpoint
	var updateUserEndpoint kitendpoint.Endpoint
	{
		findUserEndpoint = kithttp.NewClient("POST", copyURL(u, "/user/find"), encodeHTTPGenericRequest, decodeFindUserResponse, options["FindUser"]...).Endpoint()
		updateUserEndpoint = kithttp.NewClient("POST", copyURL(u, "/user/update"), encodeHTTPGenericRequest, decodeUpdateUserResponse, options["UpdateUser"]...).Endpoint()
	}

	return endpoint.Endpoints{FindUserEndpoint: findUserEndpoint, UpdateUserEndpoint: updateUserEndpoint}, nil
}

// EncodeHTTPGenericRequest is a transport/http.EncodeRequestFunc that
// JSON-encodes any request to the request body. Primarily useful in a client.
func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

// decodeFooResponse is a transport/http.DecodeResponseFunc that decodes
// a JSON-encoded concat response from the HTTP response body. If the response
// as a non-200 status code, we will interpret that as an error and attempt to
//  decode the specific error message from the response body.
func decodeFindUserResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		log.Fatalf("[client.http] decodeFundUserResponse: %+v", *r)
		return nil, handler.ErrorDecoder(r)
	}
	var resp dto.User
	err := json.NewDecoder(r.Body).Decode(&resp)
	// log.Printf("[client.http] find user result: %v, err: %v", resp, err)
	return &resp, err
}
func decodeUpdateUserResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, handler.ErrorDecoder(r)
	}
	var resp error
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
