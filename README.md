# golang-toolkit

[![CI](https://github.com/tj-actions/coverage-badge-go/workflows/CI/badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-83.1%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/coverage-badge-go/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/coverage-badge-go/actions?query=workflow%3A%22Update+release+version.%22)

## Developer tools

- [Golang Lint](https://golangci-lint.run/)
- [Golang Task](https://taskfile.dev/)
- [Golang Dependencies Update](https://github.com/oligot/go-mod-upgrade)
- [jq](https://stedolan.github.io/jq/)

### For macOs

```shell
$ brew install go-task/tap/go-task
$ brew install golangci-lint
$ go install github.com/oligot/go-mod-upgrade@latest
$ brew install jq
```

## Table of contents

* [RESTClient](#rest-client)
* [KeyValueStore](#key-value-store)

## Rest Client

# Installation

```sh
go get -u github.com/arielsrv/golang-toolkit/rest
```

# ⚡️ Quickstart

```go
package main

import (
	"github.com/arielsrv/golang-toolkit/rest"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UserDto struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

func main() {
	requestBuilder := rest.RequestBuilder{
		Timeout:        time.Millisecond * 3000,
		ConnectTimeout: time.Millisecond * 5000,
		BaseURL:        "https://gorest.co.in/public/v2",
	}

	// This won't be blocked.
	requestBuilder.AsyncGet("/users", func(response *rest.Response) {
		if response.StatusCode == http.StatusOK {
			log.Println(response)
		}
	})

	response := requestBuilder.Get("/users")
	if response.StatusCode != http.StatusOK {
		log.Fatal(response.Err.Error())
	}

	var usersDto []UserDto
	response.FillUp(&usersDto)

	var futures []*rest.FutureResponse

	requestBuilder.ForkJoin(func(c *rest.Concurrent) {
		for i := 0; i < len(usersDto); i++ {
			futures = append(futures, c.Get("/users/"+strconv.Itoa(usersDto[i].ID)))
		}
	})

	log.Println("Wait all ...")
	startTime := time.Now()
	for i := range futures {
		if futures[i].Response().StatusCode == http.StatusOK {
			var userDto UserDto
			futures[i].Response().FillUp(&userDto)
			log.Println("\t" + userDto.Name)
		}
	}
	elapsedTime := time.Since(startTime)
	log.Printf("Elapsed time: %d", elapsedTime)
}

```

## Key Value Store

# TODO