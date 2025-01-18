package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	library_http_client_go "github.com/harryosmar/http-client-go"
	"net/http"
	url2 "net/url"
	"os"
)

type (
	Response[T any] struct {
		Content    T
		StatusCode int
		Duration   int64 // in millisecond
		Header     http.Header
	}

	ResponseErr struct {
		Content    map[string]any
		StatusCode int
		Duration   int64 // in millisecond
		Header     http.Header
	}
)

func (r ResponseErr) MarshalJSON() ([]byte, error) {
	if r.Content == nil || len(r.Content) == 0 {
		return json.Marshal(map[string]any{
			"message": func() string {
				statusText := http.StatusText(r.StatusCode)
				if statusText != "" {
					return statusText
				}

				return fmt.Sprintf("err http status code %d", r.StatusCode)
			}(),
		})
	}
	return json.Marshal(r.Content)
}

func (r ResponseErr) Error() string {
	marshal, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("%+v", r.Content)
	}

	return string(marshal)
}

func ErrToResponseError(err error, response *library_http_client_go.Response) ResponseErr {
	responseErr := ResponseErr{
		Content: map[string]any{
			"message": err.Error(),
		},
	}
	if response == nil {
		return responseErr
	}

	responseErr.Duration = response.Duration
	responseErr.StatusCode = response.Status
	responseErr.Header = response.Header
	return responseErr
}

func UnmarshalResponseToError(response *library_http_client_go.Response) error {
	var (
		contentByte = response.Content
	)
	if response == nil {
		return ResponseErr{
			Content: map[string]any{
				"message": errors.New("nil response"),
			},
		}
	}

	responseErr := ResponseErr{
		Content:    nil,
		StatusCode: response.Status,
		Duration:   response.Duration,
		Header:     response.Header,
	}

	if contentByte == nil || len(contentByte) == 0 {
		return responseErr
	}

	errData := map[string]any{}
	err := json.Unmarshal(contentByte, &errData)
	if err != nil {
		responseErr.Content = map[string]any{
			"message": string(contentByte),
		}

		return responseErr
	}

	responseErr.Content = errData
	return responseErr
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

func PostFormUrlEncoded[ResT any](ctx context.Context, client library_http_client_go.HttpClientRepository, url string, payload url2.Values, headers map[string]string) (Response[ResT], error) {
	return Send[url2.Values, ResT](ctx, url, payload, headers, client.PostFormUrlEncoded)
}

func PostMultipart[ResT any](ctx context.Context, client library_http_client_go.HttpClientRepository, url string, file *os.File, headers map[string]string) (Response[ResT], error) {
	return Send[*os.File, ResT](ctx, url, file, headers, client.PostMultipart)
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
		return defaultReps, ErrToResponseError(err, resp)
	}
	var content ResT
	defaultReps = Response[ResT]{
		Content:    content,
		StatusCode: resp.Status,
		Duration:   resp.Duration,
		Header:     resp.Header,
	}

	if !(resp.Status >= 200 && resp.Status < 300) {
		return defaultReps, UnmarshalResponseToError(resp)
	}

	err = json.Unmarshal(resp.Content, &content)
	if err != nil {
		return defaultReps, ErrToResponseError(err, resp)
	}

	defaultReps.Content = content
	return defaultReps, nil
}

func SendAndReturnBytes[ReqT any](
	ctx context.Context,
	url string,
	payload ReqT,
	headers map[string]string,
	fn func(ctx context.Context, url string, data ReqT, headers map[string]string) (*library_http_client_go.Response, error),
) ([]byte, error) {
	resp, err := fn(ctx, url, payload, headers)
	if err != nil {
		return nil, ErrToResponseError(err, resp)
	}

	if !(resp.Status >= 200 && resp.Status < 300) {
		return nil, UnmarshalResponseToError(resp)
	}

	return resp.Content, nil
}
