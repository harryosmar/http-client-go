## Usage

[![Latest Version](https://img.shields.io/github/release/harryosmar/http-client-go.svg?style=flat-square)](https://github.com/harryosmar/http-client-go/releases)

![Build Status](https://github.com/harryosmar/http-client-go/actions/workflows/go-test.yml/badge.svg)

```go
package main

import (
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	http_client_go "github.com/harryosmar/http-client-go"
	"net/http"
	"testing"
	"time"
	v2 "github.com/harryosmar/http-client-go/v2"
)

func main() {
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

	// curl --location 'https://cat-fact.herokuapp.com/facts?limit=5' --header 'Content-Type: application/json'
	resp, err := v2.Get[[]FactsResponse](
		context.WithValue(context.TODO(), http_client_go.XRequestIdContext, uuid.New().String()),
		http_client_go.NewHttpClientRepository(&http.Client{Timeout: 3 * time.Second}).EnableDebug(),
		"https://cat-fact.herokuapp.com/facts",
		map[string][]string{"limit": {"5"}},
		map[string]string{"Content-Type": "application/json"},
	)
	if err != nil {
		return nil, err
	}

	log.Infof("resp %+v", resp)
}
```