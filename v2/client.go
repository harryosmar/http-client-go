package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	library_http_client_go "github.com/harryosmar/http-client-go"
	"net/http"
)

type (
	Response[T any] struct {
		Content    T
		StatusCode int
		Duration   int64 // in millisecond
		Header     http.Header
	}

	ResponseError[T any] struct {
		Success bool `json:"success,omitempty"`
		Details T    `json:"details,omitempty"`
	}

	ResponseSuccess[T any] struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    T      `json:"data"`
	}
)

func (r ResponseError[T]) Error() string {
	marshal, err := json.Marshal(r.Details)
	if err != nil {
		return fmt.Sprintf("%+v", r.Details)
	}

	return string(marshal)
}

func ErrToResponseError(err error) ResponseError[map[string]any] {
	return ResponseError[map[string]any]{
		Details: map[string]any{
			"message": err.Error(),
		},
	}
}

func UnmarshalToResponseError(contentByte []byte, statusCode int) ResponseError[map[string]any] {
	details := map[string]any{
		"message": http.StatusText(statusCode),
	}
	resp := ResponseError[map[string]any]{
		Details: details,
	}

	if contentByte == nil || len(contentByte) == 0 {
		return resp
	}

	err := json.Unmarshal(contentByte, &details)
	if err != nil {
		resp.Details["message"] = fmt.Sprintf("%v", string(contentByte))
		return resp
	}

	resp.Details = details
	return resp
}

func MarshalToBuffer[T any](content T) (*bytes.Buffer, error) {
	marshal, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(marshal), nil
}

func Post[ReqT any, ResT any](ctx context.Context, client library_http_client_go.HttpClientRepository, url string, payload ReqT, headers map[string]string) (Response[ResT], error) {
	defaultReps := Response[ResT]{}
	buffer, err := MarshalToBuffer[ReqT](payload)
	if err != nil {
		return defaultReps, err
	}

	return Send[*bytes.Buffer, ResT](ctx, url, buffer, headers, client.Post)
}

func PostRaw[ResT any](ctx context.Context, client library_http_client_go.HttpClientRepository, url string, payload []byte, headers map[string]string) (Response[ResT], error) {
	return Send[*bytes.Buffer, ResT](ctx, url, bytes.NewBuffer(payload), headers, client.Post)
}

func Put[ReqT any, ResT any](ctx context.Context, client library_http_client_go.HttpClientRepository, url string, payload ReqT, headers map[string]string) (Response[ResT], error) {
	defaultReps := Response[ResT]{}
	buffer, err := MarshalToBuffer[ReqT](payload)
	if err != nil {
		return defaultReps, err
	}

	return Send[*bytes.Buffer, ResT](ctx, url, buffer, headers, client.Put)
}

func Get[ResT any](ctx context.Context, client library_http_client_go.HttpClientRepository, url string, queries map[string][]string, headers map[string]string) (Response[ResT], error) {
	return Send[map[string][]string, ResT](ctx, url, queries, headers, client.Get)
}

func Delete[ResT any](ctx context.Context, client library_http_client_go.HttpClientRepository, url string, headers map[string]string) (Response[ResT], error) {
	return Send[any, ResT](ctx, url, nil, headers, client.DeleteX)
}

func Send[ReqT any, ResT any](
	ctx context.Context,
	url string,
	payload ReqT,
	headers map[string]string,
	fn func(ctx context.Context, url string, data ReqT, headers map[string]string) (*library_http_client_go.Response, error),
) (Response[ResT], error) {
	defaultReps := Response[ResT]{}
	resp, err := fn(ctx, url, payload, headers)
	if err != nil {
		return defaultReps, ErrToResponseError(err)
	}
	var content ResT
	defaultReps = Response[ResT]{
		Content:    content,
		StatusCode: resp.Status,
		Duration:   resp.Duration,
		Header:     resp.Header,
	}

	if !(resp.Status >= 200 && resp.Status < 300) {
		return defaultReps, UnmarshalToResponseError(resp.Content, resp.Status)
	}

	err = json.Unmarshal(resp.Content, &content)
	if err != nil {
		return defaultReps, ErrToResponseError(err)
	}

	defaultReps.Content = content
	return defaultReps, nil
}
