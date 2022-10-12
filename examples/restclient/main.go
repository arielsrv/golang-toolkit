package main

import (
	"github.com/arielsrv/golang-toolkit/restclient"
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
	requestBuilder := restclient.RequestBuilder{
		Timeout:        time.Millisecond * 3000,
		ConnectTimeout: time.Millisecond * 5000,
		BaseURL:        "https://gorest.co.in/public/v2",
	}

	response := requestBuilder.Get("/users")
	if response.StatusCode != http.StatusOK {
		log.Fatal(response.Err.Error())
	}

	var usersDto []UserDto
	response.Unmarshal(&usersDto)

	var futures []*restclient.FutureResponse

	restclient.ForkJoin(func(c *restclient.Concurrent) {
		for i := 0; i < len(usersDto); i++ {
			futures = append(futures, c.Get("https://gorest.co.in/public/v2/users/"+strconv.Itoa(usersDto[i].ID)))
		}
	})

	log.Println("Wait all ...")
	startTime := time.Now()
	for i := range futures {
		if futures[i].Response().StatusCode == http.StatusOK {
			var userDto UserDto
			futures[i].Response().Unmarshal(&userDto)
			log.Println("\t" + userDto.Name)
		}
	}
	elapsedTime := time.Since(startTime)
	log.Printf("Elapsed time: %d", elapsedTime)
}
