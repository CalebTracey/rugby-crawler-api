// Package openapi3 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.0 DO NOT EDIT.
package openapi3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetLeaderboardData request with any body
	GetLeaderboardDataWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GetLeaderboardData(ctx context.Context, body GetLeaderboardDataJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetAllLeaderboardData request
	GetAllLeaderboardData(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetLeaderboardDataWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLeaderboardDataRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetLeaderboardData(ctx context.Context, body GetLeaderboardDataJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetLeaderboardDataRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetAllLeaderboardData(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAllLeaderboardDataRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetLeaderboardDataRequest calls the generic GetLeaderboardData builder with application/json body
func NewGetLeaderboardDataRequest(server string, body GetLeaderboardDataJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetLeaderboardDataRequestWithBody(server, "application/json", bodyReader)
}

// NewGetLeaderboardDataRequestWithBody generates requests for GetLeaderboardData with any type of body
func NewGetLeaderboardDataRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/leaderboard")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewGetAllLeaderboardDataRequest generates requests for GetAllLeaderboardData
func NewGetAllLeaderboardDataRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/leaderboards")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetLeaderboardData request with any body
	GetLeaderboardDataWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetLeaderboardDataResponse, error)

	GetLeaderboardDataWithResponse(ctx context.Context, body GetLeaderboardDataJSONRequestBody, reqEditors ...RequestEditorFn) (*GetLeaderboardDataResponse, error)

	// GetAllLeaderboardData request
	GetAllLeaderboardDataWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllLeaderboardDataResponse, error)
}

type GetLeaderboardDataResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		LeaderboardData *CompetitionLeaderboardData `json:"leaderboardData,omitempty"`
		Message         *Message                    `json:"message,omitempty"`
	}
	JSON400 *struct {
		LeaderboardData *CompetitionLeaderboardData `json:"leaderboardData,omitempty"`
		Message         *Message                    `json:"message,omitempty"`
	}
	JSON500 *struct {
		LeaderboardData *CompetitionLeaderboardData `json:"leaderboardData,omitempty"`
		Message         *Message                    `json:"message,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetLeaderboardDataResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetLeaderboardDataResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAllLeaderboardDataResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		LeaderboardDataList *CompetitionLeaderboardDataList `json:"leaderboardDataList,omitempty"`
		Message             *Message                        `json:"message,omitempty"`
	}
	JSON400 *struct {
		LeaderboardDataList *CompetitionLeaderboardDataList `json:"leaderboardDataList,omitempty"`
		Message             *Message                        `json:"message,omitempty"`
	}
	JSON500 *struct {
		LeaderboardDataList *CompetitionLeaderboardDataList `json:"leaderboardDataList,omitempty"`
		Message             *Message                        `json:"message,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r GetAllLeaderboardDataResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAllLeaderboardDataResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetLeaderboardDataWithBodyWithResponse request with arbitrary body returning *GetLeaderboardDataResponse
func (c *ClientWithResponses) GetLeaderboardDataWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetLeaderboardDataResponse, error) {
	rsp, err := c.GetLeaderboardDataWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLeaderboardDataResponse(rsp)
}

func (c *ClientWithResponses) GetLeaderboardDataWithResponse(ctx context.Context, body GetLeaderboardDataJSONRequestBody, reqEditors ...RequestEditorFn) (*GetLeaderboardDataResponse, error) {
	rsp, err := c.GetLeaderboardData(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetLeaderboardDataResponse(rsp)
}

// GetAllLeaderboardDataWithResponse request returning *GetAllLeaderboardDataResponse
func (c *ClientWithResponses) GetAllLeaderboardDataWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAllLeaderboardDataResponse, error) {
	rsp, err := c.GetAllLeaderboardData(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAllLeaderboardDataResponse(rsp)
}

// ParseGetLeaderboardDataResponse parses an HTTP response from a GetLeaderboardDataWithResponse call
func ParseGetLeaderboardDataResponse(rsp *http.Response) (*GetLeaderboardDataResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetLeaderboardDataResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			LeaderboardData *CompetitionLeaderboardData `json:"leaderboardData,omitempty"`
			Message         *Message                    `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			LeaderboardData *CompetitionLeaderboardData `json:"leaderboardData,omitempty"`
			Message         *Message                    `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			LeaderboardData *CompetitionLeaderboardData `json:"leaderboardData,omitempty"`
			Message         *Message                    `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseGetAllLeaderboardDataResponse parses an HTTP response from a GetAllLeaderboardDataWithResponse call
func ParseGetAllLeaderboardDataResponse(rsp *http.Response) (*GetAllLeaderboardDataResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAllLeaderboardDataResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			LeaderboardDataList *CompetitionLeaderboardDataList `json:"leaderboardDataList,omitempty"`
			Message             *Message                        `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			LeaderboardDataList *CompetitionLeaderboardDataList `json:"leaderboardDataList,omitempty"`
			Message             *Message                        `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			LeaderboardDataList *CompetitionLeaderboardDataList `json:"leaderboardDataList,omitempty"`
			Message             *Message                        `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
