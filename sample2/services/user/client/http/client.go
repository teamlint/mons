package http

import (
	"bytes"
	"context"
	"encoding/json"
    "errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

    api "github.com/teamlint/mons/sample2/services/user/api"
    "github.com/teamlint/mons/sample2/services/user/internal/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
)

func New(instance string, options map[string][]kithttp.ClientOption) (api.UserServer, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	findEndpoint := kithttp.NewClient(
		"POST",
		copyURL(u, "/user/find"),
		encodeHTTPGenericRequest, decodeFindResponse,
		options["Find"]...,
	).Endpoint()
	updateEndpoint := kithttp.NewClient(
		"POST",
		copyURL(u, "/user/update"),
		encodeHTTPGenericRequest, decodeUpdateResponse,
		options["Update"]...,
	).Endpoint()
	return endpoint.Endpoints{
		FindEndpoint:   findEndpoint,
		UpdateEndpoint:   updateEndpoint,
	}, nil
}

func encodeHTTPGenericRequest(ctx context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}
func decodeFindResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp api.FindUserReply
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}
func decodeUpdateResponse(ctx context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp api.UpdateUserReply
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}
