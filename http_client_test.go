package library_http_client_go_test

import (
	"context"
	"github.com/google/uuid"
	http_client_go "github.com/harryosmar/http-client-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
	"time"
)

func TestGet(t *testing.T) {

	type args struct {
		t string
	}
	testData := []struct {
		name           string
		args           args
		expectedResult string
	}{
		{
			name:           "GET https://cat-fact.herokuapp.com/facts",
			args:           args{},
			expectedResult: "",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			client := http_client_go.NewHttpClientRepository(
				&http.Client{
					Timeout: 3 * time.Second,
				},
			).EnableDebug()
			ctx := context.WithValue(context.TODO(), http_client_go.XRequestIdContext, uuid.New().String())

			resp, err := client.Get(
				ctx,
				"https://cat-fact.herokuapp.com/facts",
				nil,
				map[string]string{
					"Content-Type": "application/json",
				},
			)
			if err != nil {
				t.Error(err)
				return
			}
			log.Infof("resp %+v", resp)
		})
	}
}
