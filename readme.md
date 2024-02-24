## Usage

[![Go Test](https://github.com/harryosmar/http-client-go/actions/workflows/go_test.yml/badge.svg)](https://github.com/harryosmar/http-client-go/actions/workflows/go_test.yml)

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
)

client := http_client_go.NewHttpClientRepository(&http.Client{
Timeout: 3 * time.Second,
}).EnableDebug()
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
	log.Error(err)
	return
}

log.Infof("resp %+v", resp)
```