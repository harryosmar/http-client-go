package v2_test

import (
	"bufio"
	"context"
	"github.com/google/uuid"
	http_client_go "github.com/harryosmar/http-client-go"
	v2 "github.com/harryosmar/http-client-go/v2"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"testing"
	"time"
)

func readImageFileAsByte(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func TestMethods(t *testing.T) {
	type args struct {
		fn func(client http_client_go.HttpClientRepository) (any, error)
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name: "Test Get method",
			args: args{
				fn: func(client http_client_go.HttpClientRepository) (any, error) {
					type FactsResponse struct {
						Status struct {
							Verified  bool `json:"verified"`
							SentCount int  `json:"sent_count"`
						}
						Id        string    `json:"_id"`
						User      string    `json:"user"`
						Text      string    `json:"text"`
						UpdatedAt time.Time `json:"updated_at"`
						CreatedAt time.Time `json:"created_at"`
						Deleted   bool      `json:"deleted"`
						Used      bool      `json:"used"`
					}
					ctx := context.WithValue(context.TODO(), http_client_go.XRequestIdContext, uuid.New().String())
					resp, err := v2.Get[[]FactsResponse](
						ctx,
						client,
						"https://40ccde1b-9e7d-4930-9fa8-f24f460649e9.mock.pstmn.io/api/v1/course-schedule/active/46982328",
						nil,
						map[string]string{
							"Content-Type": "application/json",
						},
					)
					if err != nil {
						return nil, err
					}
					log.Infof("%+v", resp)
					return resp, nil
				},
			},
			expectedResult: "",
		},
		{
			name: "Test Post method",
			args: args{
				fn: func(client http_client_go.HttpClientRepository) (any, error) {
					type (
						FaceDetectHttpClientResponse struct {
							Faces []struct {
								Coordinates struct {
									Height int `json:"height"`
									Width  int `json:"width"`
									X      int `json:"x"`
									Y      int `json:"y"`
								} `json:"coordinates"`
								EyesDetected bool `json:"eyes_detected"`
							} `json:"faces"`
						}
					)

					content, err := readImageFileAsByte("./example.jpg")
					if err != nil {
						return nil, err
					}
					ctx := context.WithValue(context.TODO(), http_client_go.XRequestIdContext, uuid.New().String())
					resp, err := v2.PostRaw[FaceDetectHttpClientResponse](
						ctx,
						client.DisableDebug(),
						"http://192.168.11.168:8000/detect_faces",
						content,
						map[string]string{
							"Content-Type": "image/jpeg",
						},
					)
					if err != nil {
						return nil, err
					}
					log.Infof("%+v", resp)
					return resp, nil
				},
			},
			expectedResult: "",
		},
		{
			name: "Test Post Multipart method",
			args: args{
				fn: func(client http_client_go.HttpClientRepository) (any, error) {
					type (
						FaceDetectHttpClientResponse struct {
							Faces []struct {
								Coordinates struct {
									Height int `json:"height"`
									Width  int `json:"width"`
									X      int `json:"x"`
									Y      int `json:"y"`
								} `json:"coordinates"`
								EyesDetected bool `json:"eyes_detected"`
							} `json:"faces"`
						}
					)

					content, err := os.Open("./example.jpg")
					if err != nil {
						return nil, err
					}

					ctx := context.WithValue(context.TODO(), http_client_go.XRequestIdContext, uuid.New().String())
					resp, err := v2.PostMultipart[FaceDetectHttpClientResponse](
						ctx,
						client.EnableDebug(),
						"http://192.168.11.168:5003/detect_faces",
						content,
						nil,
					)
					if err != nil {
						return nil, err
					}
					log.Infof("%+v", resp)
					return resp, nil
				},
			},
			expectedResult: "",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			client := http_client_go.NewHttpClientRepository(&http.Client{
				Timeout: 5 * time.Second,
			}).EnableDebug()
			resp, err := tt.args.fn(client)
			if err != nil {
				s := err.Error()
				log.Error(s)
				return
			}
			log.Print(resp)
		})
	}
}
