package main

import (
	"github.com/arielsrv/golang-toolkit/examples/service"
	"github.com/arielsrv/golang-toolkit/restclient"
	"log"
	"time"
)

func main() {
	restPool, err := restclient.
		NewRESTPoolBuilder().
		WithName("users").
		WithTimeout(time.Millisecond * 1000).
		WithMaxConnectionsPerHost(20).
		WithMaxIdleConnectionsPerHost(20).
		Build()

	if err != nil {
		log.Fatalln(err)
	}

	restClient := restclient.NewRESTClient(*restPool)
	userClient := service.NewUserClient(*restClient)
	userService := service.NewUserService(userClient)

	usersDto, err := userService.GetUsers()

	if err != nil {
		log.Fatal(err)
	}

	for _, userResponse := range usersDto {
		log.Printf("User: ID: %d, FullName: %s", userResponse.ID, userResponse.FullName)
	}

	var userDto service.UserDto
	userDto.FullName = "John Doe"  //nolint:nolintlint,govet
	userDto.Email = "john@doe.com" //nolint:nolintlint,govet
	userDto.Gender = "male"
	userDto.Status = "active"

	err = userService.CreateUser(userDto)
	if err != nil {
		log.Fatal(err)
	}
}
